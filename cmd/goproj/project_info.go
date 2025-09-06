package main

import (
	"fmt"

	"github.com/jvzantvoort/goproj/messages"
	"github.com/spf13/cobra"
)

// infoProjCmd represents the info command
var infoProjCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Long:  messages.GetLong("project/info"),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("info called")
	},
}

func init() {
	projectCmd.AddCommand(infoProjCmd)
}
