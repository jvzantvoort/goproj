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
	if verbose {
		log.SetLevel(log.DebugLevel)
	}
	log.Debugf("%s: start", cmd.Use)
	defer log.Debugf("%s: end", cmd.Use)


	reg := registry.NewRegistry()
	reg.Index()
}

func init() {
	registryCmd.AddCommand(IndexRegCmd)
}
