package projecttype

import (
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func WriteContent(target, content string) error {
	EnsureDir(target)

	file, err := os.Create(target)
	defer file.Close()

	_, err = file.Write([]byte(content))

	return err
}

func ReadContent(target string) (string, error) {
	if read, err := ioutil.ReadFile(target); err == nil {
		return string(read), nil
	} else {
		return "", err
	}
}

func EnsureDir(path string) {
	MkdirAll(filepath.Dir(path), 0755)
}

func MkdirAll(path string, mode int) {

	log.Debugf("MkdirAll: start")
	defer log.Debugf("MkdirAll: end")

	stat, err := os.Stat(path)

	// we found something
	if err == nil {
		// already exists
		if stat.IsDir() {
			log.Debugf("found dir: %s", path)
			return
		} else {
			log.Errorf("found target: %s but it is not a directory", path)
		}
	}

	mode_oct := os.FileMode(mode)
	os.MkdirAll(path, mode_oct)

}
