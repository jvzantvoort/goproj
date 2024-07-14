package types

import (
	log "github.com/sirupsen/logrus"
)

// ProjectTypeFile defines a structure of a file
type ProjectTypeFile struct {
	Name        string `yaml:"name"`
	Destination string `yaml:"destination"`
	Mode        string `yaml:"mode"`
}

func (ptf ProjectTypeFile) Describe() {
	log.Debugf("    - name: %s", ptf.Name)
	log.Debugf("      destination: %s", ptf.Destination)
	log.Debugf("      mode: %s", ptf.Mode)

}
