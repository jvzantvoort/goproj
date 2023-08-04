package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	log "github.com/sirupsen/logrus"
)

type ArchiveSubCmd struct {
	projecttype string
	projectname string
	archivename string
	verbose     bool
}

func (*ArchiveSubCmd) Name() string {
	return "archive"
}

func (*ArchiveSubCmd) Synopsis() string {
	return "Archive a project"
}

func (c *ArchiveSubCmd) Usage() string {
	filename := fmt.Sprintf("messages/usage_%s", c.Name())
	msgstr, err := Content.ReadFile(filename)
	if err != nil {
		log.Error(err)
		msgstr = []byte("undefined")
	}
	return string(msgstr)
}

func (c *ArchiveSubCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.archivename, "archivename", "", "Archive file")
	f.StringVar(&c.archivename, "a", "", "Archive file")
	f.StringVar(&c.projectname, "projectname", "", "Name of project")
	f.StringVar(&c.projectname, "n", "", "Name of project")
	f.BoolVar(&c.verbose, "v", false, "Verbose logging")
}

func (c *ArchiveSubCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

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
