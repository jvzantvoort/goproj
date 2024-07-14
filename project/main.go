package project

// Project defines a structure of a project type
type Project struct {
	ConfigDir          string // main config dir
	ProjectTypeDir     string // type config dir
	ProjectType        string `yaml:"projecttype"`
	Workdir            string `yaml:"workdir"`
	Pattern            string `yaml:"pattern"`
	TmuxColors         string `yaml:"tmuxcolors"`
	ProjectDescription string `yaml:"description"`
	ProjectDir         string
	ProjectName        string
	TmuxContent        string
}

func NewProject(path string) *Project {

	retv := &Project{}

	return retv

}
