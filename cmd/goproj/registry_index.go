package main

import (
	"github.com/jvzantvoort/goproj/messages"
	"github.com/jvzantvoort/goproj/registry"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// IndexRegCmd represents the index command
var IndexRegCmd = &cobra.Command{
	Use:   "index",
	Short: "Update the registry by scanning the homedir",
	Long:  messages.GetLong("registry/index"),
	Run:   handleIndexRegCmd,
}

func handleIndexRegCmd(cmd *cobra.Command, args []string) {
	log.Debug("%s: start", cmd.Use)
	defer log.Debug("%s: end", cmd.Use)

	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	reg := registry.NewRegistry()
	reg.Index()
}

func init() {
	registryCmd.AddCommand(IndexRegCmd)
}
