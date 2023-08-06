package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/jvzantvoort/goproj/project"
	log "github.com/sirupsen/logrus"
)

// SetupSubCmd missing godoc.
type SetupSubCmd struct {
	path    string
	name    string
	verbose bool
}

// Name missing godoc.
func (*SetupSubCmd) Name() string {
	return "setup"
}

// Synopsis missing godoc.
func (*SetupSubCmd) Synopsis() string {
	return "Setup project"
}

// Usage missing godoc.
func (c *SetupSubCmd) Usage() string {
	filename := fmt.Sprintf("messages/usage_%s", c.Name())
	msgstr, err := Content.ReadFile(filename)
	if err != nil {
		log.Error(err)
		msgstr = []byte("undefined")
	}
	return string(msgstr)
}

// SetFlags missing godoc.
func (c *SetupSubCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.name, "name", "", "Name of project")
	f.StringVar(&c.name, "n", "", "Name of project")
	f.StringVar(&c.path, "path", "", "Path of project")
	f.StringVar(&c.path, "p", "", "Path of project")
	f.BoolVar(&c.verbose, "v", false, "Verbose logging")
}

// Execute missing godoc.
func (c *SetupSubCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	if c.verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugln("Start")
	defer log.Debugln("End")

	if len(c.path) == 0 {
		log.Errorf("option -path/-p not provided")

		return subcommands.ExitFailure
	}

	np := project.NewProject(c.path)
	np.Functions.SetupProject()

	if len(c.name) != 0 {
		np.Name(c.name)
		err := np.WriteToFile()
		TestResult("Write config", err)
	}
	fmt.Printf("backupdir %s\n", np.Locations.BackupDirRotating("foo", 20))

	return subcommands.ExitSuccess
}
