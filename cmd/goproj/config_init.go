package main

import (
	"os/user"

	"github.com/jvzantvoort/goproj/config"
	"github.com/jvzantvoort/goproj/messages"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// InitConfigCmd represents the edit command
var InitConfigCmd = &cobra.Command{
	Use:   "init",
	Short: "Init the application configuration",
	Long:  messages.GetLong("config/init"),
	Run:   handleInitConfigCmd,
}

func GetString(cmd cobra.Command, name string) string {
	retv, _ := cmd.Flags().GetString(name)
	if len(retv) != 0 {
		log.Infof("Found %s as %s", name, retv)
	}
	return retv
}

func handleInitConfigCmd(cmd *cobra.Command, args []string) {

	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugf("%s: start", cmd.Use)
	defer log.Debugf("%s: end", cmd.Use)

	cfg := config.NewMainConfig()

	force, _ := cmd.Flags().GetBool("force")

	if force {
		cfg.ResetConfig()
	}

	user, _ := user.Current()

	cfg.UserConfig.SetMailAddress(GetString(*cmd, "mailaddress"))
	cfg.UserConfig.SetCompany(GetString(*cmd, "company"))
	cfg.UserConfig.SetCopyright(GetString(*cmd, "copyright"))
	cfg.UserConfig.SetLicense(GetString(*cmd, "license"))
	cfg.UserConfig.SetUser(GetString(*cmd, "account"), user.Username)
	cfg.UserConfig.SetUsername(GetString(*cmd, "username"))

	cfg.Save()
}

func init() {
	ConfigCmd.AddCommand(InitConfigCmd)
	InitConfigCmd.Flags().StringP("mailaddress", "m", "", "Mail address")
	InitConfigCmd.Flags().StringP("company", "c", "", "Company name")
	InitConfigCmd.Flags().StringP("copyright", "r", "", "Copyright statement")
	InitConfigCmd.Flags().StringP("license", "l", "MIT", "License statement")
	InitConfigCmd.Flags().StringP("account", "u", "", "Account name")
	InitConfigCmd.Flags().StringP("username", "n", "", "User name")
	InitConfigCmd.Flags().BoolP("force", "f", false, "Force (re)creation")
}
