package template

// File missing godoc.
type File struct {
	Name        string `yaml:"name"`
	Destination string `yaml:"destination"`
	Mode        string `yaml:"mode"`
}

// Setup missing godoc.
type Setup struct {
	Clone    []CloneUrl
	Commands []string
}

// Template missing godoc.
type Template struct {
	Name    string `yaml:"name"`
	Pattern string `yaml:"pattern"`
	Version int    `yaml:"version"`
	Setup   Setup  `yaml:"setup"`
	Files   []File `yaml:"files"`
}
