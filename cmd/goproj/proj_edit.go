package main

import (
	msg "github.com/jvzantvoort/goproj/messages"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// EditCmd represents the edit command
var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: msg.GetUsage("proj/edit"),
	Long:  msg.GetLong("proj/edit"),
	Run:   handleEditCmd,
}

func handleEditCmd(cmd *cobra.Command, args []string) {
	projectname, err := ReturnSingleArguments(cmd, args)
	if err != nil {
		return
	}

	SetLogLevel()

	log.Debugf("%s: start", cmd.Use)
	defer log.Debugf("%s: end", cmd.Use)

	log.Debugf("projectname %s", projectname)
}

func init() {
	ProjCmd.AddCommand(EditCmd)

}
