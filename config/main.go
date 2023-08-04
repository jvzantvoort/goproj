// Package config provides configuration data globally used
package config

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

type MainConfig struct {
	HomeDir            string
	TmuxDir            string
	ProjTypeConfigDir  string
	ProjTypeConfigMode int
}

func (m MainConfig) Prefix() string {
	pc, _, _, _ := runtime.Caller(1)
	elements := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	return fmt.Sprintf("GitCmd.%s", elements[len(elements)-1])
}

// ExpandHome expand the tilde in a given path.
func (m MainConfig) ExpandHome(pathstr string) (string, error) {

	m.GetHomeDir()

	if len(pathstr) == 0 {
		return pathstr, nil
	}

	if pathstr[0] != '~' {
		return pathstr, nil
	}

	return filepath.Join(m.HomeDir, pathstr[1:]), nil

}

func (m MainConfig) MkdirAll(path string, mode int) {
	log_prefix := g.Prefix()
	log.Debugf("%s: start", log_prefix)
	defer log.Debugf("%s: end", log_prefix)

	mode_oct := os.FileMode(mode)
	os.MkdirAll(path, mode_oct)

}

func (m *MainConfig) GetHomeDir() string {
	if len(m.HomeDir) != 0 {
		return m.HomeDir
	}
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	m.HomeDir = usr.HomeDir

	return m.HomeDir
}

func (m *MainConfig) GetTmuxDir() string {
	if len(m.TmuxDir) != 0 {
		return m.TmuxDir
	}
	m.TmuxDir = path.Join(m.ProjTypeConfigDir, "tmux")
	return m.TmuxDir
}

func (m MainConfig) GetProjTypeConfigDir() (string, int) {

	// return cached value
	if len(m.ProjTypeConfigDir) != 0 {
		return m.ProjTypeConfigDir, m.ProjTypeConfigMode
	}

	m.GetHomeDir() // set homedir

	retv := ""
	mode := 0755

	if runtime.GOOS == "windows" {
		mode = 0777
		retv = path.Join(m.HomeDir, "GOProj")
	} else {
		retv = path.Join(m.HomeDir, ".config", "goproj")
	}

	// check environment variable
	goproj_path, goproj_path_set := os.LookupEnv("GOPROJ_PATH")

	if goproj_path_set {
		retv = goproj_path
	}

	m.ProjTypeConfigDir = retv
	m.ProjTypeConfigMode = mode

	return m.ProjTypeConfigDir, m.ProjTypeConfigMode
}

func NewMainConfig() *MainConfig {

	log_prefix := g.Prefix()
	log.Debugf("%s: start", log_prefix)
	defer log.Debugf("%s: end", log_prefix)

	v := &MainConfig{}

	_, mode := v.GetProjTypeConfigDir()

	v.MkdirAll(v.ProjTypeConfigDir, mode)

	return v

}
