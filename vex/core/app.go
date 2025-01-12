package core

import "github.com/snackbag/vex-go/vex/widgets"

type VApp struct {
	x      int
	y      int
	width  int
	height int

	title   string
	widgets []widgets.VWidget
}

func NewApp(title string) *VApp {
	return &VApp{title: title, x: 0, y: 0, width: 640, height: 480}
}
