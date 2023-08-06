package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	log "github.com/sirupsen/logrus"
)

// ListSubCmd missing godoc.
type ListSubCmd struct {
	projecttype string
	projectname string
	printfull   bool
	verbose     bool
}

// Name missing godoc.
func (*ListSubCmd) Name() string {
	return "list"
}

// Synopsis missing godoc.
func (*ListSubCmd) Synopsis() string {
	return "List projects"
}

// Usage missing godoc.
func (c *ListSubCmd) Usage() string {
	filename := fmt.Sprintf("messages/usage_%s", c.Name())
	msgstr, err := Content.ReadFile(filename)
	if err != nil {
		log.Error(err)
		msgstr = []byte("undefined")
	}
	return string(msgstr)
}

// SetFlags missing godoc.
func (c *ListSubCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.projectname, "projectname", "", "Name of project")
	f.StringVar(&c.projectname, "n", "", "Name of project")
	f.BoolVar(&c.printfull, "f", false, "Print full")
	f.BoolVar(&c.verbose, "v", false, "Verbose logging")
}

// Execute missing godoc.
func (c *ListSubCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	if c.verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugln("Start")

	log.Debugln("End")

	return subcommands.ExitSuccess
}
