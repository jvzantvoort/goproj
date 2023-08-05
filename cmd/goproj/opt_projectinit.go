package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	log "github.com/sirupsen/logrus"
)

// InitProjSubCmd missing godoc.
type InitProjSubCmd struct {
	projecttype string
	force       bool
	verbose     bool
}

// Name missing godoc.
func (*InitProjSubCmd) Name() string {
	return "init"
}

// Synopsis missing godoc.
func (*InitProjSubCmd) Synopsis() string {
	return "Initialize a new project type"
}

// Usage missing godoc.
func (c *InitProjSubCmd) Usage() string {
	filename := fmt.Sprintf("messages/usage_%s", c.Name())
	msgstr, err := Content.ReadFile(filename)
	if err != nil {
		log.Error(err)
		msgstr = []byte("undefined")
	}
	return string(msgstr)
}

// SetFlags missing godoc.
func (c *InitProjSubCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.projecttype, "projecttype", "default", "Type of project")
	f.StringVar(&c.projecttype, "t", "default", "Type of project")
	f.BoolVar(&c.force, "f", false, "Force (re)creation")
	f.BoolVar(&c.verbose, "v", false, "Verbose logging")
}

// Execute missing godoc.
func (c *InitProjSubCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	if c.verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugln("Start")
	//
	if len(c.projecttype) == 0 {
		log.Fatalf("no type provided")
	} else if c.projecttype == "default" {
		if !c.force {
			log.Fatalf("Cannot overwrite default")
		}
	}
	log.Debugf("type: %s", c.projecttype)

	log.Debugln("End")

	return subcommands.ExitSuccess
}
