package main

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
)

func main() {
	Execute()
}

// GetHomeDir return the homedir
func GetHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return usr.HomeDir, nil
}

// ExpandHome expand the tilde in a given path.
func ExpandHome(pathstr string) (string, error) {

	if len(pathstr) == 0 {
		return pathstr, nil
	}

	if pathstr[0] != '~' {
		return pathstr, nil
	}

	homedir, err := GetHomeDir()

	if err != nil {
		return pathstr, err
	}

	return filepath.Join(homedir, pathstr[1:]), nil

}

func GetConfigDir() (string, error) {
	homedir, err := GetHomeDir()
	return path.Join(homedir, ".config", "goproj"), err

}

func GetCacheDir() string {
	homedir, _ := GetHomeDir()
	return path.Join(homedir, ".cache", "goproj")
}

func GetArchivePath(name string) string {
	return path.Join(GetCacheDir(), name)
}

func ExitOnError(err error) {
	if err != nil {
		fmt.Printf("Error:\n\t%s\n\n", err)
		os.Exit(1)
	}
}
