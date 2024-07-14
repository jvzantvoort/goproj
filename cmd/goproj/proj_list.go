package main

import (
	msg "github.com/jvzantvoort/goproj/messages"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ListCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: msg.GetUsage("proj/list"),
	Long:  msg.GetLong("proj/list"),
	Run:   handleListCmd,
}

func handleListCmd(cmd *cobra.Command, args []string) {
	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugf("%s: start", cmd.Use)
	defer log.Debugf("%s: end", cmd.Use)
}

func init() {
	ProjCmd.AddCommand(ListCmd)
	ListCmd.Flags().BoolP("full", "f", false, "print full")

}
