package registry

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/jvzantvoort/goproj/config"
	"github.com/jvzantvoort/goproj/project"
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
	defer fileh.Close()
	if err != nil {
		return err
	}

	r.SetLastUpdated()

	return r.Write(fileh)
}

// Load	loads the registry data	from the registry file
func (r *Registry) Load() error {
	fileh, err := os.Open(r.GetConfigFile())
	defer fileh.Close()
	if err != nil {
		return err
	}
	return r.Read(fileh)
}

func NewRegistry() *Registry {
	projects := make(map[string]project.Project)
	retv := &Registry{}
	retv.Projects = projects
	retv.Load()
	return retv
}
