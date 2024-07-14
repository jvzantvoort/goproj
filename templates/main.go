package messages

import (
	"embed"
	"fmt"

	log "github.com/sirupsen/logrus"
)

//go:embed shell/* projecttypes/* tmux/*
var Content embed.FS

func GetShell(name string) (string, error) {
	var msgstr []byte
	var err error
	msgstr, err = Content.ReadFile(fmt.Sprintf("shell/%s", name))
	if err != nil {
		log.Error(err)
	}
	return string(msgstr), err
}

func GetProjectTypeTemplateContent(name string) (string, error) {
	var msgstr []byte
	var err error
	msgstr, err = Content.ReadFile(fmt.Sprintf("projecttypes/%s", name))
	if err != nil {
		log.Error(err)
	}
	return string(msgstr), err
}

// templates/tmux/powerline/default/blue.tmuxtheme
func GetTmuxContent(name string) (string, error) {
	var msgstr []byte
	var err error
	msgstr, err = Content.ReadFile(fmt.Sprintf("tmux/%s.tmuxtheme", name))
	if err != nil {
		log.Error(err)
	}
	return string(msgstr), err
}
