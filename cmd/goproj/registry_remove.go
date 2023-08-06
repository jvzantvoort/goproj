package main

import (
	"github.com/jvzantvoort/goproj/messages"
	"github.com/jvzantvoort/goproj/registry"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// RemoveRegCmd represents the remove command
var RemoveRegCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a project",
	Long:  messages.GetLong("registry/remove"),
	Run:   handleRemoveRegCmd,
}

func handleRemoveRegCmd(cmd *cobra.Command, args []string) {
	log.Debug("%s: start", cmd.Use)
	defer log.Debug("%s: end", cmd.Use)

	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	name, _ := cmd.Flags().GetString("name")
	reg := registry.NewRegistry()
	reg.Remove(name)
	reg.Save()

}

func init() {
	registryCmd.AddCommand(RemoveRegCmd)
	RemoveRegCmd.Flags().StringP("name", "n", "", "project name")
	RemoveRegCmd.MarkFlagRequired("name")
	RemoveRegCmd.Flags().BoolP("prune", "p", false, "prune the project from the disk")
}
