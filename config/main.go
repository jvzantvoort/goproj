// Package config provides configuration data globally used
package config

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

type CmdPaths struct {
	// Editor points to the used used editor command
	Editor string `ini:"editor" comment:"Editor path"`
	// Vcs points to the used version control system (e.g. git)
	Vcs string `ini:"vcs" comment:"Version control path"`
}

type MainConfigMain struct {
	// TypesDir path points to the location where session types are stored
	TypesDir string `ini:"types_dir" comment:"directory to types"`
	// CacheDir path points to the location where sessions are cached
	CacheDir string `ini:"cache_dir" comment:"directory to cache"`
}

type MainConfig struct {
	HomeDir            string         `ini:"-"`
	ProjConfigDir      string         `ini:"-"`
	ProjConfigDirMode  int            `ini:"-"`
	ProjConfigFile     string         `ini:"-"`
	ProjConfigFileMode int            `ini:"-"`
	ProjTypeConfigDir  string         `ini:"-"`
	Main               MainConfigMain `ini:"main"`
	CmdPaths           `ini:"commands"`
}

// prefix returns a prefix for logging and messages based on function name.
func (m MainConfig) prefix() string {
	pc, _, _, _ := runtime.Caller(1)
	elements := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	return fmt.Sprintf("GitCmd.%s", elements[len(elements)-1])
}

func (m MainConfig) InitConfig() error {

	var err error

	// Setup logging
	log_prefix := m.prefix()
	log.Debugf("%s: start", log_prefix)
	defer log.Debugf("%s: end", log_prefix)

	m.Main.CacheDir = path.Join(m.ProjConfigDir, "cache")
	m.Main.TypesDir = path.Join(m.ProjConfigDir, "types")

	m.CmdPaths.Editor, err = Which("vim")
	WarningOnError(err)

	m.CmdPaths.Vcs, err = Which("git")
	WarningOnError(err)

	MkdirAll(m.ProjConfigDir, m.ProjConfigDirMode)

	m.Write()

	return nil
}

func (m *MainConfig) GetProjConfigDir() (string, int) {

	// Setup logging
	log_prefix := m.prefix()
	log.Debugf("%s: start", log_prefix)
	defer log.Debugf("%s: end", log_prefix)

	// return cached value
	if len(m.ProjConfigDir) != 0 {
		log.Debugf("returning cached result")
		return m.ProjConfigDir, m.ProjConfigDirMode
	}

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

	m.ProjConfigDir = retv
	m.ProjConfigDirMode = mode

	return m.ProjConfigDir, m.ProjConfigDirMode
}

func (m *MainConfig) GetProjConfigFile() (string, int) {

	// Setup logging
	log_prefix := m.prefix()
	log.Debugf("%s: start", log_prefix)
	defer log.Debugf("%s: end", log_prefix)

	// return cached value
	if len(m.ProjConfigFile) != 0 {
		log.Debugf("returning cached result")
		return m.ProjConfigFile, m.ProjConfigFileMode
	}

	m.ProjConfigFile = path.Join(m.ProjConfigDir, "main.ini")

	m.ProjConfigFileMode = 0644

	return m.ProjConfigFile, m.ProjConfigFileMode

}

func NewMainConfig() *MainConfig {

	v := &MainConfig{}

	// Setup logging
	log_prefix := v.prefix()
	log.Debugf("%s: start", log_prefix)
	defer log.Debugf("%s: end", log_prefix)

	v.HomeDir, _ = GetHomeDir()

	v.GetProjConfigDir()
	v.ProjTypeConfigDir = path.Join(v.ProjConfigDir, "types")
	log.Debugf("ProjConfigDir:  %s", v.ProjConfigDir)
	log.Debugf("ProjConfigDirMode: %v", v.ProjConfigDirMode)

	return v

}
