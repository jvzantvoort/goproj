package main

import (
	"fmt"

	"github.com/jvzantvoort/goproj/messages"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// EditRegCmd represents the edit command
var EditRegCmd = &cobra.Command{
	Use:   "edit",
	Short: "A brief description of your command",
	Long:  messages.GetLong("registry/edit"),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("edit called")
	},
}

func handleEditRegCmd(cmd *cobra.Command, args []string) {
	log.Debug("%s: start", cmd.Use)
	defer log.Debug("%s: end", cmd.Use)

	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	// name, _ := cmd.Flags().GetString("name")
}

func init() {
	registryCmd.AddCommand(EditRegCmd)
	EditRegCmd.Flags().StringP("name", "n", "", "project name")
	EditRegCmd.MarkFlagRequired("name")

}
