package template

// CloneUrl missing godoc.
type CloneUrl struct {
	Name        string `yaml:"name"`
	Url         string `yaml:"url"`
	Destination string `yaml:"destination"`
	Branch      string `yaml:"branch"`
}

// Exists missing godoc.
func (cu CloneUrl) Exists() {

}

// Clone missing godoc.
func (cu CloneUrl) Clone() {

}

// Pull missing godoc.
func (cu CloneUrl) Pull() {

}
