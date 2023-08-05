package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
	"github.com/jvzantvoort/goproj/project"
	"github.com/jvzantvoort/goproj/registry"
	log "github.com/sirupsen/logrus"
)

// RegisterSubCmd is a subcommand to register a project.
type RegisterSubCmd struct {
	path    string
	name    string
	verbose bool
}

// Name a function to return the name of the command
func (*RegisterSubCmd) Name() string {
	return "register"
}

// Synopsis a function to return the synopsis of the command
func (*RegisterSubCmd) Synopsis() string {
	return "Register project"
}

// Usage a function to return the usage of the command
func (c *RegisterSubCmd) Usage() string {
	filename := fmt.Sprintf("messages/usage_%s", c.Name())
	msgstr, err := Content.ReadFile(filename)
	if err != nil {
		log.Error(err)
		msgstr = []byte("undefined")
	}
	return string(msgstr)
}

// SetFlags handles flags
func (c *RegisterSubCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&c.verbose, "v", false, "Verbose logging")
}

// Execute subcommand
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
		for name, proj := range reg.Projects {
			fmt.Printf("%s %s\n", name, proj.Description())
		}
	case "info":
		name := remainder[0]
		remainder = remainder[1:]
		proj, ok := reg.Get(name)
		if !ok {
			fmt.Printf("Project %s not found\n", name)
			return subcommands.ExitSuccess
		}
		proj.WriteTable(os.Stdout)
	default:
		fmt.Printf(c.Usage())
	}

	return subcommands.ExitSuccess
}
