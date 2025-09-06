package project

import (
	"fmt"
	"os"
	"path"

	log "github.com/sirupsen/logrus"
)

// Locations the locations used in the object
type Locations struct {
	RootDir   string `json:"root"`
	BackupDirectory string `json:"backupdir"`
}

// ConfigDir missing godoc.
func (L Locations) ConfigDir() string {
	return path.Join(L.RootDir, "."+ProjectName)
}

// ConfigFile missing godoc.
func (L Locations) ConfigFile() string {
	return path.Join(L.ConfigDir(), "settings.json")
}

// BinPath missing godoc.
func (L Locations) BinPath() string {
	return path.Join(L.RootDir, "bin")
}

// ToolsPath missing godoc.
func (L Locations) ToolsPath() string {
	return path.Join(L.ConfigDir(), "libexec")
}

// BackupDir missing godoc.
func (L Locations) BackupDir(args ...string) string {
	if option, err := OneOrLess(args...); err == nil {
		L.BackupDirectory = option
	}
	if L.BackupDirectory == "" {
		L.BackupDirectory = path.Join(L.RootDir, "backup")
	}
	return L.BackupDirectory
}

// BackupDirRotating missing godoc.
func (L Locations) BackupDirRotating(name string, max int) string {
	log.Debugf("BackupDirRotating: %s %d", name, max)
	defer log.Debugf("BackupDirRotating: %s %d", name, max)
	basedir := L.BackupDir()
	paths := []string{}

	max += 1

	for num := 1; num < max; num++ {
		namestr := fmt.Sprintf("%s-%04d", name, num)
		paths = append(paths, path.Join(basedir, namestr))
	}
	paths = Reverse(paths)

	listlen := len(paths) - 1

	lastpath := paths[0]

	log.Infof("remove %s", lastpath)

	for indx, path := range paths {
		if indx == listlen {
			break
		}
		curpath := path
		prepath := paths[indx+1]
		if _, err := os.Stat(prepath); err != nil {
			log.Debugf("no such target %s", prepath)
			continue
		}
		log.Infof("mv %s %s", prepath, curpath)
		// err := os.Rename(curpath, prepath)
		// if err != nil {
		// 	log.Fatal(err)
		// }
	}
	return paths[len(paths)-1]
}
