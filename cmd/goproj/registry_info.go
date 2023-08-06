package main

import (
	"fmt"

	"github.com/jvzantvoort/goproj/messages"
	"github.com/spf13/cobra"
)

// infoRegCmd represents the info command
var infoRegCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Long:  messages.GetLong("registry/info"),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("info called")
	},
}

func init() {
	registryCmd.AddCommand(infoRegCmd)
}
