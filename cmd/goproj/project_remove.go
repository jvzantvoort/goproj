package main

import (
	"github.com/jvzantvoort/goproj/messages"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// RemoveProjCmd represents the remove command
var RemoveProjCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a project",
	Long:  messages.GetLong("project/remove"),
	Run:   handleRemoveProjCmd,
}

func handleRemoveProjCmd(cmd *cobra.Command, args []string) {
	log.Debug("%s: start", cmd.Use)
	defer log.Debug("%s: end", cmd.Use)

	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	// name, _ := cmd.Flags().GetString("name")
}

func init() {
	projectCmd.AddCommand(RemoveProjCmd)
	RemoveProjCmd.Flags().StringP("name", "n", "", "project name")
	RemoveProjCmd.MarkFlagRequired("name")
	RemoveProjCmd.Flags().BoolP("prune", "p", false, "prune the project from the disk")
}
