package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/jvzantvoort/goproj/project"
	"github.com/jvzantvoort/goproj/registry"
	log "github.com/sirupsen/logrus"
)

// RegisterSubCmd missing godoc.
type RegisterSubCmd struct {
	path    string
	name    string
	verbose bool
}

// Name missing godoc.
func (*RegisterSubCmd) Name() string {
	return "register"
}

// Synopsis missing godoc.
func (*RegisterSubCmd) Synopsis() string {
	return "Register project"
}

// Usage missing godoc.
func (c *RegisterSubCmd) Usage() string {
	filename := fmt.Sprintf("messages/usage_%s", c.Name())
	msgstr, err := Content.ReadFile(filename)
	if err != nil {
		log.Error(err)
		msgstr = []byte("undefined")
	}
	return string(msgstr)
}

// SetFlags missing godoc.
func (c *RegisterSubCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&c.verbose, "v", false, "Verbose logging")
}

// Execute missing godoc.
func (c *RegisterSubCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	arguments := []string{}

	// Handle verbosity
	if c.verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugln("Start")
	defer log.Debugln("End")

	if f.NArg() == 0 {
		fmt.Printf(c.Usage())

	return subcommands.ExitSuccess
	}

	arguments = f.Args()[:]

	reg := registry.NewRegistry()
	action := arguments[0]
	remainder := arguments[1:]
	switch action {
	case "add":
		pathtoproject := remainder[0]
		remainder = remainder[1:]
		project := project.NewProject(pathtoproject)
		reg.Register(*project)
		reg.Save()
	case "list":
		for _, name := range reg.List() {
			fmt.Println(name)
		}
	default:
		fmt.Printf(c.Usage())
	}

	return subcommands.ExitSuccess
}
