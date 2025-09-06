package project

import (
	"fmt"
	"io"
	"strings"
)

// EnvVar missing godoc.
type EnvVar struct {
	RootPath string   `json:"-"`
	Name     string   `json:"name"`
	Paths    []string `json:"paths"`
}

// Environment missing godoc.
type Environment struct {
	RootDir string            `json:"-"`
	Vars    map[string]EnvVar `json:"vars"`
}

// Write missing godoc.
// ExportString write the name and value to the buffer as a bash export string
func (evar EnvVar) Write(w io.Writer) {
	fmt.Fprintf(w, "export %s=\"%s\"\n", evar.Name, strings.Join(evar.Paths, ":"))
}

// Has missing godoc.
func (evar EnvVar) Has(inputdir string) bool {
	for _, element := range evar.Paths {
		if element == inputdir {
			return true
		}
	}
	return false
}

// Append missing godoc.
func (evar *EnvVar) Append(inputdir string) {

	if (len(inputdir) == 0) || (evar.Has(inputdir)) {
		return
	}

	evar.Paths = append(evar.Paths, inputdir)

}

// Prepend missing godoc.
func (evar *EnvVar) Prepend(inputdir string) {
	if (len(inputdir) == 0) || (evar.Has(inputdir)) {
		return
	}

	evar.Paths = append([]string{inputdir}, evar.Paths...)
}
