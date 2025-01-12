package core

import "github.com/snackbag/vex-go/vex/widgets"

type VApp struct {
	x      int
	y      int
	width  int
	height int

	visibility bool

	title   string
	widgets []widgets.VWidget
}

func (app *VApp) Show() {
	app.SetVisibility(true)
}

func (app *VApp) Hide() {
	app.SetVisibility(false)
}

func (app *VApp) SetVisibility(visibility bool) {
	app.visibility = visibility
}

func NewApp(title string) *VApp {
	return &VApp{title: title, x: 0, y: 0, width: 640, height: 480}
}
