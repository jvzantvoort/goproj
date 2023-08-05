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
	return r.Projects[name], true
}

func (r *Registry) List() []string {
	var names []string
	for name := range r.Projects {
		names = append(names, name)
	}
	return names
}

func (r *Registry) Remove(name string) {
	delete(r.Projects, name)
}

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

func (r Registry) Write(writer io.Writer) error {
	jsonString, err := json.MarshalIndent(r, "", "  ")
	if err == nil {
		fmt.Fprint(writer, string(jsonString))
		fmt.Fprintf(writer, "\n")
	}
	return err
}

func (r *Registry) SetLastUpdated(t int64) {
	r.LastUpdated = t
}

func (r Registry) GetConfigFile() string {

	mc := config.NewMainConfig()
	return mc.ConfigFile(config.RegistryFile)
}

func (r *Registry) Save() error {

	fileh, err := os.OpenFile(r.GetConfigFile(), os.O_WRONLY|os.O_CREATE, 0644)
	defer fileh.Close()
	if err != nil {
		return err
	}

	r.SetLastUpdated(time.Now().Unix())

	return r.Write(fileh)
}

func (r *Registry) Load() error {
	fileh, err := os.Open(r.GetConfigFile())
	defer fileh.Close()
	if err != nil {
		return err
	}
	return r.Read(fileh)
}

func NewRegistry() *Registry {
	return &Registry{
		Projects: make(map[string]project.Project),
	}
}
