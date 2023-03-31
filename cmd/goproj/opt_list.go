package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	gop "github.com/jvzantvoort/goproj"
	"github.com/jvzantvoort/goproj/session"
	"github.com/jvzantvoort/goproj/config"
	log "github.com/sirupsen/logrus"
)

type ListSubCmd struct {
	sessiontype string
	sessionname string
	printfull   bool
	verbose     bool
}

func (*ListSubCmd) Name() string {
	return "list"
}

func (*ListSubCmd) Synopsis() string {
	return "List sessions"
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
	f.StringVar(&c.sessionname, "sessionname", "", "Name of session")
	f.StringVar(&c.sessionname, "n", "", "Name of session")
	f.BoolVar(&c.printfull, "f", false, "Print full")
	f.BoolVar(&c.verbose, "v", false, "Verbose logging")
}

func (c *ListSubCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	if c.verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugln("Start")
	cfg := config.NewMainConfig()
	cfg.Read()

	sessions := session.NewSessions()

	sessions.CacheDir = cfg.Main.CacheDir

	sessions.Load()

	sessions.Writer(os.Stdout)

	log.Debugln("End")

	return subcommands.ExitSuccess
}
