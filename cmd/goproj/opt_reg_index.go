package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/jvzantvoort/goproj/registry"
	log "github.com/sirupsen/logrus"
)

// RegIndexSubCmd is a subcommand to register a project.
type RegIndexSubCmd struct {
	path    string
	name    string
	verbose bool
}

// Name a function to return the name of the command
func (*RegIndexSubCmd) Name() string {
	return "index"
}

// Synopsis a function to return the synopsis of the command
func (*RegIndexSubCmd) Synopsis() string {
	return "Update registry by scanning homedir"
}

// Usage a function to return the usage of the command
func (c *RegIndexSubCmd) Usage() string {
	filename := fmt.Sprintf("messages/usage_%s", c.Name())
	msgstr, err := Content.ReadFile(filename)
	if err != nil {
		log.Error(err)
		msgstr = []byte("undefined")
	}
	return string(msgstr)
}

// SetFlags handles flags
func (c *RegIndexSubCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&c.verbose, "v", false, "Verbose logging")
}

// Execute subcommand
func (c *RegIndexSubCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	// arguments := []string{}

	// Handle verbosity
	if c.verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugln("Start")
	defer log.Debugln("End")

	// arguments = f.Args()[:]

	reg := registry.NewRegistry()
	log.Infof("Start")
	reg.Index()
	log.Infof("End")

	return subcommands.ExitSuccess
}
