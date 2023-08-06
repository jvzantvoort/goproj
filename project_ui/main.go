package project_ui

import (
	"github.com/jvzantvoort/goproj/project"
	"github.com/rivo/tview"
)

type ProjectUI struct {
	Project project.Project
}

var (
	app   *tview.Application
	pages *tview.Pages
)

func (ui *ProjectUI) MainWindows() {
	app = tview.NewApplication()

	form := tview.NewForm().
		AddInputField("Name:", ui.Project.Name(), 30, nil, nil).
		AddInputField("Description:", ui.Project.Description(), 30, nil, nil).
		AddInputField("Root:", ui.Project.Locations.RootDir, 30, nil, nil).
		AddInputField("Config Dir:", ui.Project.Locations.ConfigDir(), 30, nil, nil).
		AddButton("Save", ui.Project.Save).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("Project Info").SetTitleAlign(tview.AlignLeft)
	if err := app.EnableMouse(true).SetRoot(form, true).Run(); err != nil {
		panic(err)
	}
}

func NewProjectUI(p *project.Project) *ProjectUI {
	retv := &ProjectUI{}
	retv.Project = *p
	return retv
}
