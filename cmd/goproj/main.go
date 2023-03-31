package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		DisableColors:          false,
		TimestampFormat:        "15:04:05",
	})

	//	TimestampFormat:        "2006-01-02 15:04:05",
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {

	subcommands.Register(subcommands.HelpCommand(), "common")
	subcommands.Register(subcommands.FlagsCommand(), "common")
	subcommands.Register(subcommands.CommandsCommand(), "common")

	subcommands.Register(&ConfigSubCmd{}, "main")
	subcommands.Register(&SetupgcSubCmd{}, "main")

	subcommands.Register(&SessionSubCmd{}, "session")
	subcommands.Register(&ListSubCmd{}, "session")
	subcommands.Register(&ArchiveSubCmd{}, "session")
	subcommands.Register(&ShellProfileCmd{}, "session")

	subcommands.Register(&CreateSubCmd{}, "")
	subcommands.Register(&EditSubCmd{}, "")
	subcommands.Register(&InitProjSubCmd{}, "")
	subcommands.Register(&ListFilesSubCmd{}, "")
	// subcommands.Register(&ResumeSubCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))

}
