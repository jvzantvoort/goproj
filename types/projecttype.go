package types

import (
	"fmt"
	"os"
	"path"

	templates "github.com/jvzantvoort/goproj/templates"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

// ProjectType defines a structure of a project type
type ProjectType struct {
	ConfigDir      string            // main config dir
	ProjectTypeDir string            // type config dir
	ProjectType    string            `yaml:"projecttype"`
	Workdir        string            `yaml:"workdir"`
	Pattern        string            `yaml:"pattern"`
	TmuxColors     string            `yaml:"tmuxcolors"`
	TmuxContent    string
	SetupActions   []string          `yaml:"setupactions"`
	Files          []ProjectTypeFile `yaml:"files"`
}

// Install a config file
func (ptc ProjectType) Install(name string, perm os.FileMode) error {
	var content string
	var dest string
	var err error
	err = nil

	dest = path.Join(ptc.ProjectTypeDir, name)

	// Get the template content
	content, _ = templates.GetProjectTypeTemplateContent(name)
	buf, _ := ParseTemplate(content, ptc)

	err = WriteToFile(dest, buf.Bytes(), perm)
	return err

}

func (ptc *ProjectType) Read() {

	viper.SetConfigName("config")
	viper.AddConfigPath(ptc.ProjectTypeDir)

	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	err = viper.Unmarshal(&ptc)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}

// Create a new project type
func (ptc *ProjectType) Create() {
	err := Mkdir(ptc.ProjectTypeDir)
	if err != nil {
		log.Errorf("%s", err)
		return
	}
	ptc.Install("config.yml", 0644)
	ptc.Install("default.rc", 0644)
	ptc.Install("default.env", 0644)

}

// Initialize config
func (ptc *ProjectType) Init(configdir, projecttype string) error {
	ptc.ProjectType = projecttype
	ptc.ConfigDir = configdir
	ptc.ProjectTypeDir = path.Join(ptc.ConfigDir, ptc.ProjectType)
	return nil
}

// NewProjectType read the relevant configfile and return
// ProjectType object with relevant data.
func NewProjectType(configdir, projecttype string) *ProjectType {
	v := &ProjectType{}
	v.Init(configdir, projecttype)
	return v
}
