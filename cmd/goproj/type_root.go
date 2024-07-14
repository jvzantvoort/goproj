package main

import (
	msg "github.com/jvzantvoort/goproj/messages"

	"github.com/spf13/cobra"
)

// TypeCmd represents the type command
var TypeCmd = &cobra.Command{
	Use:   "type",
	Short: msg.GetUsage("type/root"),
	Long:  msg.GetLong("type/root"),
}

func init() {
	rootCmd.AddCommand(TypeCmd)

}
