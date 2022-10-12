package main

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	gop "github.com/jvzantvoort/goproj"
	log "github.com/sirupsen/logrus"
)

type ResumeSubCmd struct {
	projectname string
	verbose     bool
}

func (*ResumeSubCmd) Name() string {
	return "resume"
}

func (*ResumeSubCmd) Synopsis() string {
	return "Resume a project"
}

func (*ResumeSubCmd) Usage() string {
	msgstr, err := gop.Asset("messages/usage_resume")
	if err != nil {
		log.Error(err)
		msgstr = []byte("undefined")
	}
	return string(msgstr)
}

func (c *ResumeSubCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.projectname, "projectname", "", "Name of project")
	f.StringVar(&c.projectname, "n", "", "Name of project")
	f.BoolVar(&c.verbose, "v", false, "Verbose logging")
}

func (c *ResumeSubCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

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
