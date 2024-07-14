package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func SetLogLevel() {
	if verbose {
		log.SetLevel(log.DebugLevel)
	}
}

func ReturnSingleArguments(cmd *cobra.Command, args []string) (string, error) {
	// We need at least the project name
	if len(args) != 1 {
		cmd.Help()
		return "", fmt.Errorf("argument missing")
	}
	return args[0], nil
}
