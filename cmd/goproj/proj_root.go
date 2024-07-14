package main

import (
	msg "github.com/jvzantvoort/goproj/messages"

	"github.com/spf13/cobra"
)

// ProjCmd represents the proj command
var ProjCmd = &cobra.Command{
	Use:   "proj",
	Short: msg.GetUsage("proj/root"),
	Long:  msg.GetLong("proj/root"),
}

func init() {
	rootCmd.AddCommand(ProjCmd)

}
