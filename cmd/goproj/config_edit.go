package main

import (
	"os/user"

	"github.com/jvzantvoort/goproj/config"
	"github.com/jvzantvoort/goproj/messages"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// EditConfigCmd represents the edit command
var EditConfigCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the application configuration",
	Long:  messages.GetLong("config/edit"),
	Run:   handleEditConfigCmd,
}

func handleEditConfigCmd(cmd *cobra.Command, args []string) {
	if verbose {
		log.SetLevel(log.DebugLevel)
	}
	log.Debugf("%s: start", cmd.Use)
	defer log.Debugf("%s: end", cmd.Use)

	cfg := config.NewMainConfig()

	interactive := cmd.Flags().Changed("interactive")
	user, _ := user.Current()

	cfg.UserConfig.SetMailAddress(GetString(*cmd, "mailaddress"))
	cfg.UserConfig.SetCompany(GetString(*cmd, "company"))
	cfg.UserConfig.SetCopyright(GetString(*cmd, "copyright"))
	cfg.UserConfig.SetLicense(GetString(*cmd, "license"))
	cfg.UserConfig.SetUser(GetString(*cmd, "account"), user.Username)
	cfg.UserConfig.SetUsername(GetString(*cmd, "username"))

	if interactive {
		config.UserConfigUI(cfg)
	}

	cfg.Save()

}

func init() {
	ConfigCmd.AddCommand(EditConfigCmd)
	EditConfigCmd.Flags().StringP("mailaddress", "m", "", "Mail address")
	EditConfigCmd.Flags().StringP("company", "c", "", "Company name")
	EditConfigCmd.Flags().StringP("copyright", "r", "", "Copyright statement")
	EditConfigCmd.Flags().StringP("license", "l", "MIT", "License statement")
	EditConfigCmd.Flags().StringP("account", "u", "", "Account name")
	EditConfigCmd.Flags().StringP("username", "n", "", "User name")
	EditConfigCmd.Flags().BoolP("interactive", "i", false, "Interactive")

}
