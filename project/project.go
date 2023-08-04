package project

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	log "github.com/sirupsen/logrus"
)

// Project the project object
//
//	proj := NewProject("/home/foo/project")
type Project struct {
	MetaData  MetaData  `json:"metadata"`
	Locations Locations `json:"locations"`
	Functions Functions `json:"-"`
}

// File handling
//
// Read
func (p *Project) Read(reader io.Reader) error {
	data, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &p)
	if err != nil {
		return err
	}

	return nil
}

// Write
func (p Project) Write(writer io.Writer) error {
	jsonString, err := json.MarshalIndent(p, "", "  ")
	if err == nil {
		fmt.Fprint(writer, string(jsonString))
		fmt.Fprintf(writer, "\n")
	}
	return err

}

// ReadFromFile
func (p *Project) ReadFromFile() error {
	settings := p.Locations.ConfigFile()

	fileh, err := os.Open(settings)
	defer fileh.Close()

	if err != nil {
		return err
	}

	return p.Read(fileh)
}

// WriteToFile
func (p Project) WriteToFile() error {
	settings := p.Locations.ConfigFile()
	log.Debugf("save to: %s, start", settings)
	defer log.Debugf("save to: %s, end", settings)
	fileh, err := os.OpenFile(settings, os.O_WRONLY|os.O_CREATE, 0644)
	defer fileh.Close()

	if err != nil {
		return err
	}
	return p.Write(fileh)
}

func (p *Project) Name(args ...string) string {
	if option, err := OneOrLess(args...); err == nil {
		log.Debugf("set name to %s", option)
		p.MetaData.Project.Name = option
	}
	return p.MetaData.Project.Name
}

func (p *Project) Description(args ...string) string {
	if option, err := OneOrLess(args...); err == nil {
		log.Debugf("set description to %s", option)
		p.MetaData.Project.Description = option
	}
	return p.MetaData.Project.Description
}

func NewProject(path string) *Project {
	retv := &Project{}

	retv.Locations.RootDir = path

	retv.ReadFromFile()

	fn := NewFunctions(retv.Locations)
	retv.Functions = *fn
	return retv
}