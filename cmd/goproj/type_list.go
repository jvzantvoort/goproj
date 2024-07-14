package main

import (
	"fmt"
	"path"
	"sort"

	msg "github.com/jvzantvoort/goproj/messages"
	"github.com/jvzantvoort/goproj/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// TypeListCmd represents the list command
var TypeListCmd = &cobra.Command{
	Use:   "list",
	Short: msg.GetUsage("type/list"),
	Long:  msg.GetLong("type/list"),
	Run:   handleTypeListCmd,
}

func handleTypeListCmd(cmd *cobra.Command, args []string) {
	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugf("%s: start", cmd.Use)
	defer log.Debugf("%s: end", cmd.Use)
	configdir, _ := GetConfigDir()
	typesdir := path.Join(configdir, "types")

	names := types.ListProjectTypes(typesdir)
	sort.Strings(names)

	for _, indx := range names {
		fmt.Printf(" - %s\n", indx)
	}
}

func init() {
	TypeCmd.AddCommand(TypeListCmd)

}
