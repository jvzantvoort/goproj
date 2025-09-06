package registry

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/jvzantvoort/goproj/config"
	"github.com/jvzantvoort/goproj/project"
	"github.com/jvzantvoort/goproj/utils"

	"github.com/olekukonko/tablewriter"
)

type Registry struct {
	Projects    map[string]project.Project `json:"projects"`
	LastUpdated int64                      `json:"lastUpdated"`
}

func (r *Registry) Register(p project.Project) {
	r.Projects[p.Name()] = p
}

func (r *Registry) Get(name string) (project.Project, bool) {
	retv := project.Project{}
	retv, err := r.Projects[name]
	return retv, err
}

func (r *Registry) List() []string {
	var retv []string
	for name := range r.Projects {
		retv = append(retv, name)
	}
	return retv
}

func (r Registry) WriteTable(writer io.Writer) {
	table := tablewriter.NewWriter(writer)
	table.Header([]string{"Name", "Value"})

	for _, project := range r.Projects {
		table.Append([]string{project.Name(), project.Description()})
	}
	table.Render()
}

// Remove removes a project from the registry data
//
//	reg := registry.NewRegistry()
//	reg.Remove("foo")
func (r *Registry) Remove(name string) {
	delete(r.Projects, name)
}

// Read reads the registry data from the registry file
func (r *Registry) Read(reader io.Reader) error {
	data, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &r)
	if err != nil {
		return err
	}
	return nil
}

// Write writes the registry data to the registry file
func (r Registry) Write(writer io.Writer) error {
	jsonString, err := json.MarshalIndent(r, "", "  ")
	if err == nil {
		fmt.Fprint(writer, string(jsonString))
		fmt.Fprintf(writer, "\n")
	}
	return err
}

// SetLastUpdated sets the last updated time in the registry data
func (r *Registry) SetLastUpdated() {
	r.LastUpdated = time.Now().Unix()
}

// GetConfigFile returns the config file path from the registry
func (r Registry) GetConfigFile() string {

	mc := config.NewMainConfig()
	return mc.ConfigFile(config.RegistryFile)
}

// Save saves the registry data	to the registry file
func (r *Registry) Save() error {

	fileh, err := os.OpenFile(r.GetConfigFile(), os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer fileh.Close()

	r.SetLastUpdated()

	return r.Write(fileh)
}

// Prune removes all non-existing projects from the registry
func (r *Registry) Prune() error {
	for name, project := range r.Projects {
		if !project.IsGoProj() {
			delete(r.Projects, name)
		}
	}
	return nil
}

// Index scan the homedir and add goproj compatible projects
func (r *Registry) Index() error {

	basedir := utils.GetHomeDir()

	filewalk_error := filepath.Walk(basedir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && info.Name() == ".goproj" {
			projdir := filepath.Dir(path)
			proj := project.NewProject(projdir)
			r.Register(*proj)
		}
		return nil
	})
	if filewalk_error != nil {
		log.Fatal(filewalk_error)
		return filewalk_error
	}
	r.Prune() // remove non-existing projects
	r.Save()
	return nil
}

// Load	loads the registry data	from the registry file
func (r *Registry) Load() error {
	fileh, err := os.Open(r.GetConfigFile())
	if err != nil {
		return err
	}
	defer fileh.Close()
	return r.Read(fileh)
}

func NewRegistry() *Registry {
	projects := make(map[string]project.Project)
	retv := &Registry{}
	retv.Projects = projects
	retv.Load()
	return retv
}
