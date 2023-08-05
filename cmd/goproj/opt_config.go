package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/jvzantvoort/goproj/config"
	log "github.com/sirupsen/logrus"
)

// ConfigSubCmd missing godoc.
type ConfigSubCmd struct {
	write   bool
	force   bool
	verbose bool
}

// Name missing godoc.
func (*ConfigSubCmd) Name() string {
	return "config"
}

// Synopsis missing godoc.
func (*ConfigSubCmd) Synopsis() string {
	return "Configure the goproj project"
}

// Usage missing godoc.
func (c *ConfigSubCmd) Usage() string {
	return GetUsage(c.Name())
}

// SetFlags missing godoc.
func (c *ConfigSubCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&c.write, "w", false, "Write settings")
	f.BoolVar(&c.force, "f", false, "Force (re)creation")
	f.BoolVar(&c.verbose, "v", false, "Verbose logging")
}

// Execute missing godoc.
func (c *ConfigSubCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	if c.verbose {
		log.SetLevel(log.DebugLevel)
	}

	arguments := []string{}

	log.Debugln("Start")
	cfg := config.NewMainConfig()
	if c.force {
		cfg.ResetConfig()
		cfg.Save()
	}

	if len(f.Args()) >= 1 {
		arguments = f.Args()[:]
	} else {
		fmt.Printf(c.Usage())
		fmt.Printf("\tfor more information use -h\n\n")
		return subcommands.ExitSuccess
	}

	if len(arguments) == 1 {
		val, err := cfg.Get(arguments[0])
		if err == nil {
			fmt.Println(val)
			return subcommands.ExitSuccess
		}

		if arguments[0] == "fields" {
			for _, indx := range cfg.Fields() {
				fmt.Printf(" - %s\n", indx)
			}
			return subcommands.ExitSuccess
		}

		fmt.Printf("Error: %s\n", err)
		return subcommands.ExitFailure
	}

	log.Debugln("End")
	return subcommands.ExitSuccess
}
