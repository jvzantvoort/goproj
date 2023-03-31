package config

import (
	"os"

	"gopkg.in/ini.v1"
	log "github.com/sirupsen/logrus"
)

func (m MainConfig) Write() error {

	// Setup logging
	log_prefix := m.prefix()
	log.Debugf("%s: start", log_prefix)
	defer log.Debugf("%s: end", log_prefix)

	configfile, mode := m.GetProjConfigFile()

	cfg := ini.Empty()
	err := ini.ReflectFrom(cfg, &m)
	if err != nil {
		return err
	}

	err = cfg.SaveTo(configfile)
	if err != nil {
		return err
	}
	mode_oct := os.FileMode(mode)
	os.Chmod(configfile, mode_oct)
	return nil
}

func (m *MainConfig) Read() error {

	// Setup logging
	log_prefix := m.prefix()
	log.Debugf("%s: start", log_prefix)
	defer log.Debugf("%s: end", log_prefix)

	configfile, _ := m.GetProjConfigFile()

	_, err := os.Stat(configfile)

	var cfg *ini.File

	if err != nil {
		err = m.Write()
		if err != nil {
			return err
		}
	}
	cfg, err = ini.Load(configfile)
	if err != nil {
		return err
	}

	err = cfg.MapTo(m)
	if err != nil {
		return err
	}
	return nil
}
