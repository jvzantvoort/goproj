package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	log "github.com/sirupsen/logrus"
)

type ShellProfileCmd struct {
	shellname string
	verbose   bool
}

func (*ShellProfileCmd) Name() string {
	return "shell"
}

func (*ShellProfileCmd) Synopsis() string {
	return "Edit a projects tmux configuration"
}

func (c *ShellProfileCmd) Usage() string {
	filename := fmt.Sprintf("messages/usage_%s", c.Name())
	msgstr, err := Content.ReadFile(filename)
	if err != nil {
		log.Error(err)
		msgstr = []byte("undefined")
	}
	return string(msgstr)
}

func (c *ShellProfileCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.shellname, "shellname", "bash", "Name of the shell profile to provide")
	f.StringVar(&c.shellname, "s", "bash", "Name of the shell profile to provide")
	f.BoolVar(&c.verbose, "v", false, "Verbose logging")
}

func (c *ShellProfileCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	if c.verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugln("Start")

	msgstr, err := Content.ReadFile("messages/" + c.shellname)
	if err != nil {
		msgstr = []byte("# undefined")
		if c.verbose {
			log.Errorf("Error: %s", err)

		}
	}
	fmt.Print(string(msgstr))

	log.Debugln("End")

	return subcommands.ExitSuccess
}
