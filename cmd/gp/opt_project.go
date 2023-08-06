package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/jvzantvoort/goproj/project"
	"github.com/jvzantvoort/goproj/utils"
	"github.com/jvzantvoort/goproj/project_ui"
	log "github.com/sirupsen/logrus"
)

// ProjectSubCmd missing godoc.
type ProjectSubCmd struct {
	projecttype string
	force       bool
	verbose     bool
}

// Name missing godoc.
func (*ProjectSubCmd) Name() string {
	return "project"
}

// Synopsis missing godoc.
func (*ProjectSubCmd) Synopsis() string {
	return "Project based actions"
}

// Usage missing godoc.
func (c *ProjectSubCmd) Usage() string {
	filename := fmt.Sprintf("messages/usage_%s", c.Name())
	msgstr, err := Content.ReadFile(filename)
	if err != nil {
		log.Error(err)
		msgstr = []byte("undefined")
	}
	return string(msgstr)
}

// SetFlags missing godoc.
func (c *ProjectSubCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.projecttype, "projecttype", "default", "Type of project")
	f.StringVar(&c.projecttype, "t", "default", "Type of project")
	f.BoolVar(&c.force, "f", false, "Force (re)creation")
	f.BoolVar(&c.verbose, "v", false, "Verbose logging")
}

func (c ProjectSubCmd) ProjectType() (string, error) {
	if len(c.projecttype) == 0 {
		return "", ErrTypeNotProvided
	}
	return c.projecttype, nil
}

// Execute missing godoc.
func (c *ProjectSubCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

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

	utils.PrintSuccess("This is %d %% oke\n", 80)

	arguments = f.Args()[:]

	action := arguments[0]
	remainder := arguments[1:]

	switch action {
	case "create":
		inputpath := remainder[0]
		proj := project.NewProject(inputpath)

		ui := project_ui.NewProjectUI(proj)
		ui.MainWindows()
	default:
		fmt.Printf(c.Usage())
	}

	return subcommands.ExitSuccess
}
