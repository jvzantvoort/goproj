package main

import (
	msg "github.com/jvzantvoort/goproj/messages"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CreateCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: msg.GetUsage("proj/create"),
	Long:  msg.GetLong("proj/create"),
	Run:   handleCreateCmd,
}

func handleCreateCmd(cmd *cobra.Command, args []string) {
	projectname, err := ReturnSingleArguments(cmd, args)
	if err != nil {
		return
	}

	SetLogLevel()

	log.Debugf("%s: start", cmd.Use)
	defer log.Debugf("%s: end", cmd.Use)

	// set archive name
	projecttype, err := cmd.Flags().GetString("projecttype")
	if err != nil {
		log.Errorf("Error: %s", err)
	}

	log.Debugf("projectname %s", projectname)
	log.Debugf("projecttype %s", projecttype)
}

func init() {
	ProjCmd.AddCommand(CreateCmd)
	CreateCmd.Flags().StringP("projecttype", "t", "default", "Type of project")
}
