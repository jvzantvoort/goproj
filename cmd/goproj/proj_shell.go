package main

import (
	"bytes"
	"fmt"
	"path"
	"text/template"

	msg "github.com/jvzantvoort/goproj/messages"
	templates "github.com/jvzantvoort/goproj/templates"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type ShellCmdStruct struct {
	HomeDir    string
	ConfigDir  string
	SessionDir string
	Shell      string
}

// ShellCmd represents the shell command
var ShellCmd = &cobra.Command{
	Use:   "shell",
	Short: msg.GetUsage("proj/shell"),
	Long:  msg.GetLong("proj/shell"),
	Run:   handleShellCmd,
}

func returnShell(t ShellCmdStruct) (string, error) {
	var retv string
	var err error
	err = nil

	retv, _ = templates.GetShell(t.Shell)

	tmpl, err := template.New("shell").Parse(retv)
	if err != nil {
		return retv, err
	}

	buf := new(bytes.Buffer)

	err = tmpl.Execute(buf, t)
	if err != nil {
		return retv, err
	}
	retv = buf.String()

	return retv, err

}

func handleShellCmd(cmd *cobra.Command, args []string) {
	// projectname, err := ReturnSingleArguments(cmd, args)
	// if err != nil {
	// 	return
	// }

	SetLogLevel()

	log.Debugf("%s: start", cmd.Use)
	defer log.Debugf("%s: end", cmd.Use)

	t := ShellCmdStruct{}
	t.HomeDir, _ = GetHomeDir()
	t.ConfigDir, _ = GetConfigDir()
	t.SessionDir = path.Join(t.ConfigDir, "tmux.d")
	t.Shell = "bash"
	content, err := returnShell(t)
	if err != nil {
		log.Errorf("%s", err)
	}
	fmt.Print(content)

}

func init() {
	ProjCmd.AddCommand(ShellCmd)
	ShellCmd.Flags().StringP("shell", "s", "bash", "Name of the shell profile to provide")
}
