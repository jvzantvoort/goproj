package main

import (
	"fmt"

	"github.com/jvzantvoort/goproj/messages"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// EditProjCmd represents the edit command
var EditProjCmd = &cobra.Command{
	Use:   "edit",
	Short: "A brief description of your command",
	Long:  messages.GetLong("project/edit"),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("edit called")
	},
}

func handleEditProjCmd(cmd *cobra.Command, args []string) {
	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugf("%s: start", cmd.Use)
	defer log.Debugf("%s: end", cmd.Use)

	// name, _ := cmd.Flags().GetString("name")
}

func init() {
	projectCmd.AddCommand(EditProjCmd)
	EditProjCmd.Flags().StringP("name", "n", "", "project name")
	EditProjCmd.MarkFlagRequired("name")

}
