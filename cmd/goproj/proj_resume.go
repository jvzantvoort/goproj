package main

import (
	msg "github.com/jvzantvoort/goproj/messages"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ResumeCmd represents the resume command
var ResumeCmd = &cobra.Command{
	Use:   "resume",
	Short: msg.GetUsage("proj/resume"),
	Long:  msg.GetLong("proj/resume"),
	Run:   handleResumeCmd,
}

func handleResumeCmd(cmd *cobra.Command, args []string) {
	projectname, err := ReturnSingleArguments(cmd, args)
	if err != nil {
		return
	}

	SetLogLevel()

	log.Debugf("%s: start", cmd.Use)
	defer log.Debugf("%s: end", cmd.Use)

	log.Debugf("projectname %s", projectname)
}

func init() {
	ProjCmd.AddCommand(ResumeCmd)

}
