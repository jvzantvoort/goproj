package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	log "github.com/sirupsen/logrus"
)

// CreateSubCmd missing godoc.
type CreateSubCmd struct {
	projecttype string
	projectname string
	verbose     bool
}

// Name missing godoc.
func (*CreateSubCmd) Name() string {
	return "create"
}

// Synopsis missing godoc.
func (*CreateSubCmd) Synopsis() string {
	return "Create a new project"
}

// Usage missing godoc.
func (c *CreateSubCmd) Usage() string {
	filename := fmt.Sprintf("messages/usage_%s", c.Name())
	msgstr, err := Content.ReadFile(filename)
	if err != nil {
		log.Error(err)
		msgstr = []byte("undefined")
	}
	return string(msgstr)
}

// SetFlags missing godoc.
func (c *CreateSubCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.projecttype, "projecttype", "default", "Type of project")
	f.StringVar(&c.projecttype, "t", "default", "Type of project")
	f.StringVar(&c.projectname, "projectname", "", "Name of project")
	f.StringVar(&c.projectname, "n", "", "Name of project")
	f.BoolVar(&c.verbose, "v", false, "Verbose logging")
}

// Execute missing godoc.
func (c *CreateSubCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	if c.verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugln("Start")
	//
	if len(c.projectname) == 0 {
		log.Fatalf("no name provided")
	}

	log.Debugln("End")

	return subcommands.ExitSuccess
}
