package main

import (
	"github.com/jvzantvoort/goproj/messages"
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"proj"},
	Short:   "project related commands",
	Long:    messages.GetLong("project/root"),
}

func init() {

	// Setup root
	rootCmd.AddCommand(projectCmd)

}
