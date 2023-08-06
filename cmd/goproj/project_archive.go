package main

import (
	"github.com/jvzantvoort/goproj/messages"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ArchiveProjCmd represents the archive command
var ArchiveProjCmd = &cobra.Command{
	Use:   "archive",
	Short: "Archive a project",
	Long:  messages.GetLong("project/archive"),
	Run:   handleArchiveProjCmd,
}

func handleArchiveProjCmd(cmd *cobra.Command, args []string) {
	log.Debug("%s: start", cmd.Use)
	defer log.Debug("%s: end", cmd.Use)

	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	// name, _ := cmd.Flags().GetString("name")
}

func init() {

	// archive
	projectCmd.AddCommand(ArchiveProjCmd)
	ArchiveProjCmd.Flags().StringP("name", "n", "", "project name")
	ArchiveProjCmd.Flags().StringP("archivename", "a", "", "Archive file")
	ArchiveProjCmd.MarkFlagRequired("name")

}
