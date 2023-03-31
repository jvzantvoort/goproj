package projecttype

import log "github.com/sirupsen/logrus"

// Describe describe
func (ptc ProjectTypeConfig) Describe() {
	log.Debugf("Describe: %s", ptc.ProjectType)
	log.Debugf("  Pattern: %s", ptc.Pattern)

	fileno := len(ptc.Files)
	actionsno := len(ptc.SetupActions)

	if fileno > 0 {
		log.Debugf("  Files:")
		for _, act := range ptc.Files {
			log.Debugf("    - name: %s", act.Name)
			log.Debugf("      destination: %s", act.Destination)
			log.Debugf("      mode: %s", act.Mode)
		}
	}

	if actionsno > 0 {
		log.Debugf("  Actions:")
		for _, act := range ptc.SetupActions {
			log.Debugf("    - %s", act)
		}

	}

	log.Debugf("Describe: %s, end", ptc.ProjectType)

}
