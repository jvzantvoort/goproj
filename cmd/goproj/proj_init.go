package main

import (
	msg "github.com/jvzantvoort/goproj/messages"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// InitCmd represents the init command
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: msg.GetUsage("proj/init"),
	Long:  msg.GetLong("proj/init"),
	Run:   handleInitCmd,
}

func handleInitCmd(cmd *cobra.Command, args []string) {
	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugf("%s: start", cmd.Use)
	defer log.Debugf("%s: end", cmd.Use)
}

func init() {
	ProjCmd.AddCommand(InitCmd)

}
