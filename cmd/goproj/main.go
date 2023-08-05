package main

import (
	"embed"
	"fmt"

	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	log "github.com/sirupsen/logrus"
)

// Content missing godoc.
//
//go:embed messages/*
var Content embed.FS

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		TimestampFormat:        "2006-01-02 15:04:05",
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func GetUsage(name string) string {
	filename := fmt.Sprintf("messages/usage_%s", name)
	msgstr, err := Content.ReadFile(filename)
	if err != nil {
		log.Error(err)
		msgstr = []byte("undefined")
	}
	return string(msgstr)
}

func main() {

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")

	subcommands.Register(&ConfigSubCmd{}, "application")

	subcommands.Register(&SetupSubCmd{}, "update")

	subcommands.Register(&CreateSubCmd{}, "proj")
	subcommands.Register(&RegisterSubCmd{}, "proj")
	subcommands.Register(&EditSubCmd{}, "proj")
	subcommands.Register(&ListSubCmd{}, "proj")
	subcommands.Register(&ArchiveSubCmd{}, "proj")
	subcommands.Register(&ShellProfileCmd{}, "proj")
	subcommands.Register(&ListFilesSubCmd{}, "proj")
	subcommands.Register(&ResumeSubCmd{}, "proj")

	subcommands.Register(&InitProjSubCmd{}, "type")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))

}
