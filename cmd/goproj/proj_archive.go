package main

import (
	msg "github.com/jvzantvoort/goproj/messages"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ArchiveCmd represents the archive command
var ArchiveCmd = &cobra.Command{
	Use:   "archive",
	Short: msg.GetUsage("proj/archive"),
	Long:  msg.GetLong("proj/archive"),
	Run:   handleArchiveCmd,
}

func handleArchiveCmd(cmd *cobra.Command, args []string) {
	projectname, err := ReturnSingleArguments(cmd, args)
	if err != nil {
		return
	}

	SetLogLevel()

	log.Debugf("%s: start", cmd.Use)
	defer log.Debugf("%s: end", cmd.Use)

	// set archive name
	archivename, err := cmd.Flags().GetString("archivename")
	if err != nil {
		log.Errorf("Error: %s", err)
	}

	if len(archivename) == 0 {
		archivename = GetArchivePath(projectname)
	}

	log.Debugf("projectname %s", projectname)
	log.Debugf("archivename %s", archivename)
}

func init() {
	ProjCmd.AddCommand(ArchiveCmd)
	ArchiveCmd.Flags().StringP("archivename", "a", "", "Archive file")
}
