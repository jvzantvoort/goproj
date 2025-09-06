package project

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
)

// Project the project object
//
//	proj := NewProject("/home/foo/project")
type Project struct {
	MetaData  MetaData  `json:"metadata"`
	Locations Locations `json:"locations"`
	Targets   Targets   `json:"targets"`
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

func (p Project) WriteTable(writer io.Writer) {
	table := tablewriter.NewWriter(writer)
	table.Header([]string{"Name", "Value"})
	table.Append([]string{"Name", p.MetaData.Project.Name})
	table.Append([]string{"Description", p.MetaData.Project.Description})
	table.Append([]string{"Root", p.Locations.RootDir})
	table.Render()
}

func (p Project) IsGoProj() bool {
	info, err := os.Stat(p.Locations.ConfigDir())
	if err != nil {
		return false
	}
	return info.IsDir()
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

// Name missing godoc.
func (p *Project) Name(args ...string) string {
	if option, err := OneOrLess(args...); err == nil {
		p.SetName(option)
	}
	return p.MetaData.Project.Name
}

func (p *Project) SetName(name string) {
	log.Debugf("set name to %s", name)
	p.MetaData.Project.Name = name
}

// Description missing godoc.
func (p *Project) Description(args ...string) string {
	if option, err := OneOrLess(args...); err == nil {
		p.SetDescription(option)
	}
	return p.MetaData.Project.Description
}

func (p *Project) SetDescription(description string) {
	log.Debugf("set description to %s", description)
	p.MetaData.Project.Description = description
}

func (p Project) Save() {
}

// NewProject missing godoc.
func NewProject(path string) *Project {
	retv := &Project{}

	retv.Locations.RootDir = path

	retv.ReadFromFile()

	fn := NewFunctions(retv.Locations)
	retv.Functions = *fn
	return retv
}
