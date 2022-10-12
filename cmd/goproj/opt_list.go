package main

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	gop "github.com/jvzantvoort/goproj"
	log "github.com/sirupsen/logrus"
)

type ListSubCmd struct {
	projecttype string
	projectname string
	printfull   bool
	verbose     bool
}

func (*ListSubCmd) Name() string {
	return "list"
}

func (*ListSubCmd) Synopsis() string {
	return "List projects"
}

func (*ListSubCmd) Usage() string {
	msgstr, err := gop.Asset("messages/usage_list")
	if err != nil {
		log.Error(err)
		msgstr = []byte("undefined")
	}
	return string(msgstr)
}

func (c *ListSubCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.projectname, "projectname", "", "Name of project")
	f.StringVar(&c.projectname, "n", "", "Name of project")
	f.BoolVar(&c.printfull, "f", false, "Print full")
	f.BoolVar(&c.verbose, "v", false, "Verbose logging")
}

func (c *ListSubCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	if c.verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugln("Start")

	log.Debugln("End")

	return subcommands.ExitSuccess
}
