package main

import (
	"github.com/jvzantvoort/goproj/messages"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CreateProjCmd represents the create command
var CreateProjCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project",
	Long:  messages.GetLong("project/create"),
	Run:   handleCreateProjCmd,
}

func handleCreateProjCmd(cmd *cobra.Command, args []string) {

	if verbose {
		log.SetLevel(log.DebugLevel)
	}
	log.Debugf("%s: start", cmd.Use)
	defer log.Debugf("%s: end", cmd.Use)

	// name, _ := cmd.Flags().GetString("name")
}

func init() {
	// Create
	projectCmd.AddCommand(CreateProjCmd)
	CreateProjCmd.Flags().StringP("name", "n", "", "project name")
	CreateProjCmd.MarkFlagRequired("name")
	CreateProjCmd.Flags().StringP("type", "t", "default", "type of project")
}
