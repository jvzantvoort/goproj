package types

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"text/template"

	log "github.com/sirupsen/logrus"
)

func TargetExists(targetpath string) bool {
	_, err := os.Stat(targetpath)
	if err != nil {
		return false
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func Mkdir(path string) error {
	perm := os.FileMode(int(0755))

	target_stat, err := os.Stat(path)
	// target exists and is a directory
	if err == nil {
		if target_stat.IsDir() {
			return nil
		}
		return fmt.Errorf("target %s exists but is not a directory", path)
	}
	if err := os.MkdirAll(path, perm); err != nil {
		return fmt.Errorf("directory cannot be created: %s", path)
	}
	return nil
}

func ParseTemplate(content string, data any) (*bytes.Buffer, error) {
	retv := new(bytes.Buffer)

	// parse them as a template
	tmpl, err := template.New("tmpl").Parse(content)
	if err == nil {
		err = tmpl.Execute(retv, data)
	}
	return retv, err
}

func WriteToFile(name string, content []byte, perm os.FileMode) error {
	filehandle, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, perm)

	if err != nil {
		return err
	}

	if _, err := filehandle.Write(content); err != nil {
		filehandle.Close()
		log.Fatal(err)
		return err
	}
	if err := filehandle.Close(); err != nil {
		log.Fatal(err)
	}
	return err
}

func ListProjectTypes(configdir string) []string {
	var retv []string
	err := Mkdir(configdir)
	if err != nil {
		log.Errorf("Error: %s", err)
		return retv
	}

	targets, err := os.ReadDir(configdir)
	if err != nil {
		log.Errorf("Error: %s", err)
		return retv
	}

	for _, target := range targets {
		if !target.IsDir() {
			continue
		}
		name := target.Name()
		fpath := path.Join(configdir, name, "config.yml")
		if TargetExists(fpath) {
			retv = append(retv, name)
		}
	}
	return retv
}
