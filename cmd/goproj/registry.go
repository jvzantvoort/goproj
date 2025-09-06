package main

import (
	"github.com/jvzantvoort/goproj/messages"
	"github.com/spf13/cobra"
)

// registryCmd represents the reg command
var registryCmd = &cobra.Command{
	Use:     "registry",
	Aliases: []string{"reg"},
	Short:   "Handle registry",
	Long:    messages.GetLong("registry/root"),
}

func init() {
	rootCmd.AddCommand(registryCmd)
}
