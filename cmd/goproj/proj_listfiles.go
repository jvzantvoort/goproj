package main

import (
	msg "github.com/jvzantvoort/goproj/messages"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ListFilesCmd represents the listfiles command
var ListFilesCmd = &cobra.Command{
	Use:   "listfiles",
	Short: msg.GetUsage("proj/listfiles"),
	Long:  msg.GetLong("proj/listfiles"),
	Run:   handleListFilesCmd,
}

func handleListFilesCmd(cmd *cobra.Command, args []string) {
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
	ProjCmd.AddCommand(ListFilesCmd)

}
