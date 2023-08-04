package template

type CloneUrl struct {
	Name        string `yaml:"name"`
	Url         string `yaml:"url"`
	Destination string `yaml:"destination"`
	Branch      string `yaml:"branch"`
}

func (cu CloneUrl) Exists() {

}

func (cu CloneUrl) Clone() {

}

func (cu CloneUrl) Pull() {

}
