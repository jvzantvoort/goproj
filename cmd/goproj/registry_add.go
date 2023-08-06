package main

import (
	"path/filepath"

	"github.com/jvzantvoort/goproj/messages"
	"github.com/jvzantvoort/goproj/project"
	"github.com/jvzantvoort/goproj/registry"
	"github.com/jvzantvoort/goproj/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// AddRegCmd represents the create command
var AddRegCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a project to the registry",
	Long:  messages.GetLong("registry/add"),
	Run:   handleAddRegCmd,
}

func handleAddRegCmd(cmd *cobra.Command, args []string) {
	log.Debug("%s: start", cmd.Use)
	defer log.Debug("%s: end", cmd.Use)

	var path string
	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	reg := registry.NewRegistry()

	if len(args) > 0 {
		path = args[0]
	} else {
		log.Error("No path provided")
	}
	path, _ = filepath.Abs(path)

	project := project.NewProject(path)

	if !project.IsGoProj() {
		utils.ErrorBox("This is not a valid Project")
		return
	}

	reg.Register(*project)
	reg.Save()
}

func init() {
	// Add
	registryCmd.AddCommand(AddRegCmd)
}
