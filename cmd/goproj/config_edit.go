package main

import (
	"github.com/jvzantvoort/goproj/messages"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// EditConfigCmd represents the edit command
var EditConfigCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the application configuration",
	Long:  messages.GetLong("config/edit"),
	Run:   handleEditConfigCmd,
}

func handleEditConfigCmd(cmd *cobra.Command, args []string) {
	log.Debug("%s: start", cmd.Use)
	defer log.Debug("%s: end", cmd.Use)

	if verbose {
		log.SetLevel(log.DebugLevel)
	}
}

func init() {
	ConfigCmd.AddCommand(EditConfigCmd)
	EditConfigCmd.Flags().StringP("name", "n", "", "project name")

}
