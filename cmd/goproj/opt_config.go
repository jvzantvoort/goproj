package main

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"github.com/jvzantvoort/goproj/config"
	log "github.com/sirupsen/logrus"
)

// ConfigSubCmd missing godoc.
type ConfigSubCmd struct {
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
	f.BoolVar(&c.force, "f", false, "Force (re)creation")
	f.BoolVar(&c.verbose, "v", false, "Verbose logging")
}

// Execute missing godoc.
func (c *ConfigSubCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	if c.verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugln("Start")
	cfg := config.NewMainConfig()
	if c.force {
		cfg.ResetConfig()
	}
	cfg.Save()
	log.Debugln("End")
	return subcommands.ExitSuccess
}
