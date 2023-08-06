package main

import (
	"github.com/jvzantvoort/goproj/messages"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CreateRegCmd represents the create command
var CreateRegCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project",
	Long:  messages.GetLong("registry/create"),
	Run:   handleCreateRegCmd,
}

func handleCreateRegCmd(cmd *cobra.Command, args []string) {

	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugf("%s: start", cmd.Use)
	defer log.Debugf("%s: end", cmd.Use)

	// name, _ := cmd.Flags().GetString("name")
}

func init() {
	// Create
	registryCmd.AddCommand(CreateRegCmd)
	CreateRegCmd.Flags().StringP("name", "n", "", "project name")
	CreateRegCmd.MarkFlagRequired("name")
}