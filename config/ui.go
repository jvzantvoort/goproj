package config

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func UserConfigUI(cfg *MainConfig) {

	app := tview.NewApplication()

	form := tview.NewForm().
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetFieldTextColor(tcell.ColorWhite).
		SetButtonBackgroundColor(tcell.ColorYellow).
		SetButtonTextColor(tcell.ColorBlack).
		AddInputField("Mail Address", cfg.UserConfig.MailAddress, 50, nil,
			func(answer string) {
				cfg.UserConfig.MailAddress = answer
			}).
		AddInputField("Company", cfg.UserConfig.Company, 50, nil,
			func(answer string) {
				cfg.UserConfig.Company = answer
			}).
		AddInputField("Copyright", cfg.UserConfig.Copyright, 50, nil,
			func(answer string) {
				cfg.UserConfig.Copyright = answer
			}).
		AddInputField("License", cfg.UserConfig.License, 50, nil,
			func(answer string) {
				cfg.UserConfig.License = answer
			}).
		AddInputField("User", cfg.UserConfig.User, 50, nil,
			func(answer string) {
				cfg.UserConfig.User = answer
			}).
		AddInputField("Username", cfg.UserConfig.Username, 50, nil,
			func(answer string) {
				cfg.UserConfig.Username = answer
			}).
		AddButton("Save", cfg.Save).
		AddButton("Quit", func() {
			app.Stop()
		})

	form.SetBorder(true).SetTitle("User Config").SetTitleAlign(tview.AlignLeft)
	if err := app.EnableMouse(true).SetRoot(form, true).Run(); err != nil {
		panic(err)
	}

}
