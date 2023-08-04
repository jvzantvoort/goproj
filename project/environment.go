package project

import (
	"fmt"
	"io"
	"strings"
)

type EnvVar struct {
	RootPath string `json:"-"`
	Name  string   `json:"name"`
	Paths []string `json:"paths"`
}

type Environment struct {
	RootDir string            `json:"-"`
	Vars    map[string]EnvVar `json:"vars"`
}

// ExportString write the name and value to the buffer as a bash export string
func (evar EnvVar) Write(w io.Writer) {
	fmt.Fprintf(w, "export %s=\"%s\"\n", evar.Name(), strings.Join(evar.Paths, ":"))
}

func (evar EnvVar) Has(inputdir string) bool {
	for _, element := range evar.Paths {
		if element == inputdir {
			return true
		}
	}
	return false
}

func (evar *EnvVar) Append(inputdir string) {

	if (len(inputdir) == 0) || (evar.Has(inputdir)) {
		return
	}

	evar.Paths = append(evar.Paths, inputdir)

}

func (evar *Path) Prepend(inputdir string) {
	if (len(inputdir) == 0) || (evar.Has(inputdir)) {
		return
	}

	evar.Paths = append([]string{inputdir}, evar.Paths...)
}
