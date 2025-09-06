package main

import (
	"github.com/jvzantvoort/goproj/messages"
	"github.com/spf13/cobra"
)

// ConfigCmd represents the project command
var ConfigCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"cfg"},
	Short:   "Config commands",
	Long:    messages.GetLong("config/root"),
}

func init() {
	rootCmd.AddCommand(ConfigCmd)
}
