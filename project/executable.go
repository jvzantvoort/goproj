package project

import (
	"fmt"
	"runtime"

	log "github.com/sirupsen/logrus"
)

type Executable struct {
	PathToScript string
	OS           string
	Extensions   []string
}

func (e Executable) LookupExt() (string, error) {
	for _, ext := range e.Extensions {
		log.Debugf("check extension: %s for %s", ext, e.PathToScript)
		cmndpath := e.PathToScript

		// no extension
		if len(ext) != 0 {
			cmndpath = fmt.Sprintf("%s.%s", e.PathToScript, ext)
		}

		if FileIsExecutable(cmndpath) {
			return cmndpath, nil
		}
	}
	return e.PathToScript, ErrFileNotFound
}

func (e *Executable) PrependExt(extstr string) {
	e.Extensions = append([]string{extstr}, e.Extensions...)
}

func (e *Executable) AppendExt(extstr string) {
	e.Extensions = append(e.Extensions, extstr)
}

func NewExecutable(command string, extensions ...string) *Executable {
	retv := &Executable{}
	retv.PathToScript = command
	retv.OS = runtime.GOOS

	// build extension list
	for _, ext := range extensions {
		retv.AppendExt(ext)
	}

	if retv.OS == "windows" {
		retv.AppendExt("exe")
		retv.AppendExt("cmd")
		retv.AppendExt("bat")
	} else {
		retv.AppendExt("sh")
		retv.PrependExt("")
	}
	retv.AppendExt("py")
	retv.AppendExt("pl")

	return retv
}
