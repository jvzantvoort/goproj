package config

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

// MkdirAll create directory
func MkdirAll(path string, mode int) {
	log_prefix := "MkdirAll"
	log.Debugf("%s: start", log_prefix)
	defer log.Debugf("%s: end", log_prefix)

	finfo, err := os.Stat(path)
	// we found something
	if err == nil {
		// already exists
		if finfo.IsDir() {
			log.Debugf("found dir: %s", path)
			return
		} else {
			log.Errorf("found target: %s but it is not a directory", path)
		}
	}
	mode_oct := os.FileMode(mode)
	os.MkdirAll(path, mode_oct)

}

// Which returns a command's path
func Which(command string) (string, error) {
	for _, dirname := range strings.Split(os.Getenv("PATH"), ":") {
		fullpath := path.Join(dirname, command)

		stat, err := os.Stat(fullpath)
		if err != nil {
			continue
		}

		switch mode := stat.Mode(); {
		case mode.IsDir():
			continue
		case mode&0100 != 0:
			return fullpath, nil
		case mode&0010 != 0:
			return fullpath, nil
		case mode&0001 != 0:
			return fullpath, nil
		}
	}
	return command, fmt.Errorf("Command %s not found", command)
}

func GetHomeDir() (string, error) {
	usr, err := user.Current()
	if err == nil {
		return usr.HomeDir, nil
	}
	return "", err
}

// ExpandHome expand the tilde in a given path.
func ExpandHome(pathstr string) (string, error) {

	if len(pathstr) == 0 {
		return pathstr, nil
	}

	if pathstr[0] != '~' {
		return pathstr, nil
	}
	HomeDir, _ := GetHomeDir()

	return filepath.Join(HomeDir, pathstr[1:]), nil

}

func WarningOnError(err error) {
	if err != nil {
		log.Warningf("error %v\n", err)
	}
}

// ExitOnError check error and exit if not nil
func ExitOnError(err error) {
	if err != nil {
		log.Errorf("error %v\n", err)
		os.Exit(1)
	}
}
