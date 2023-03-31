package projecttype

import (
	"fmt"
	"os"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/jvzantvoort/goproj/config"
)

var (
	mainconfig = config.NewMainConfig()
)

// ProjectTypeFile defines a structure of a file
type ProjectTypeFile struct {
	Name        string `yaml:"name"`
	Destination string `yaml:"destination"`
	Mode        string `yaml:"mode"`
}

// ProjectTypeConfig defines a structure of a project type
type ProjectTypeConfig struct {
	ProjectType    string `yaml:"projecttype"`
	ProjectTypeDir string
	Pattern        string            `yaml:"pattern"`
	SetupActions   []string          `yaml:"setupactions"`
	Files          []ProjectTypeFile `yaml:"files"`
}

func (ptc ProjectTypeConfig) UpdateConfigFile(target string) error {

	content, err := ReadContent(target)
	if err != nil {
		return err
	}

	ncontent := strings.Replace(content, "PROJECTTYPE", ptc.ProjectType, -1)
	if content == ncontent {
		return nil
	}

	return WriteContent(target, ncontent)
}

func (ptc *ProjectTypeConfig) Init(projtypeconfigdir, projecttype string) error {

	log.Debugf("Init Start: %s", projecttype)
	projtypeconfigdir = path.Join(projtypeconfigdir, projecttype)

	ptc.ProjectType = projecttype
	ptc.ProjectTypeDir = projtypeconfigdir

	stat, err := os.Stat(ptc.ProjectTypeDir)
	if stat.IsDir() {
		return fmt.Errorf("Directory already exists: %s", ptc.ProjectTypeDir)
	} else {
		if ! os.IsNotExist(err) {
			return fmt.Errorf("Target already exists: %s but is not a directory", ptc.ProjectTypeDir)
		}
	}

	if err := os.MkdirAll(ptc.ProjectTypeDir, os.FileMode(int(0755))); err != nil {
		return fmt.Errorf("Directory cannot be created: %s", ptc.ProjectTypeDir)
	}

	// write basic files
	targets := ListTemplates()
	for _, target := range targets {
		fpath := path.Join(ptc.ProjectTypeDir, target)
		err := ptc.WriteTemplate(target, fpath)
		if err != nil {
			return fmt.Errorf("Error: %s", err)
		}
		err = ptc.UpdateConfigFile(fpath)
		if err != nil {
			return fmt.Errorf("Error: %s", err)
		}
	}

	return nil
}

// NewProjectTypeConfig read the relevant configfile and return
// ProjectTypeConfig object with relevant data.
func NewProjectTypeConfig(projecttype string) ProjectTypeConfig {

	// Load main configuration targets
	projtypeconfigdir := path.Join(mainconfig.ProjTypeConfigDir, projecttype)

	log.Debugf("project type config dir: %s", projtypeconfigdir)

	v := ProjectTypeConfig{}
	v.readConfig(projtypeconfigdir)

	return v
}

func CreateProjectType(projecttype string) error {
	var pt ProjectTypeConfig
	pt.Init(mainconfig.ProjTypeConfigDir, projecttype)
	return nil
}
