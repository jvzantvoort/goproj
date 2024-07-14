package main

import (
	msg "github.com/jvzantvoort/goproj/messages"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// TypeInitCmd represents the list command
var TypeInitCmd = &cobra.Command{
	Use:   "init",
	Short: msg.GetUsage("type/init"),
	Long:  msg.GetLong("type/init"),
	Run:   handleTypeInitCmd,
}

func handleTypeInitCmd(cmd *cobra.Command, args []string) {
	typename, err := ReturnSingleArguments(cmd, args)
	if err != nil {
		return
	}

	SetLogLevel()


	log.Debugf("%s: start", cmd.Use)
	defer log.Debugf("%s: end", cmd.Use)

	log.Debugf("typename %s", typename)
}

func init() {
	TypeCmd.AddCommand(TypeInitCmd)

}
