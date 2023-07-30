package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	gop "github.com/jvzantvoort/goproj"
	"github.com/jvzantvoort/goproj/project"
	log "github.com/sirupsen/logrus"
)

type SetupSubCmd struct {
	path string
	name string
	verbose     bool
}

func (*SetupSubCmd) Name() string {
	return "setup"
}

func (*SetupSubCmd) Synopsis() string {
	return "Setup project"
}

func (*SetupSubCmd) Usage() string {
	msgstr, err := gop.Asset("messages/usage_list")
	if err != nil {
		log.Error(err)
		msgstr = []byte("undefined")
	}
	return string(msgstr)
}

func (c *SetupSubCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.name, "name", "", "Name of project")
	f.StringVar(&c.name, "n", "", "Name of project")
	f.StringVar(&c.path, "path", "", "Path of project")
	f.StringVar(&c.path, "p", "", "Path of project")
	f.BoolVar(&c.verbose, "v", false, "Verbose logging")
}

func (c *SetupSubCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	if c.verbose {
		log.SetLevel(log.DebugLevel)
	}

	if len(c.path) == 0 {
		log.Errorf("option -path/-p not provided")

		return subcommands.ExitFailure

	}

	log.Debugln("Start")
	np := project.NewProject(c.path)
	fmt.Printf("err: %s\n", np.Locations.RootDir)

	log.Debugln("End")

	return subcommands.ExitSuccess
}
