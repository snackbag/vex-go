package core

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/snackbag/vex-go/vex/widgets"
)

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
	if visibility == !app.visibility {
		app.visibility = visibility

		if visibility {
			app.startRaylibRendering()
		} else {
			rl.CloseWindow()
		}
	}
}

func (app *VApp) startRaylibRendering() {
	rl.SetTraceLogLevel(rl.LogError)
	rl.InitWindow(int32(app.width), int32(app.height), app.title)

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() && app.visibility {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		for _, element := range app.widgets {
			element.Render()
		}

		rl.EndDrawing()
	}

	if rl.WindowShouldClose() || !app.visibility {
		rl.CloseWindow()
	}
}

func NewApp(title string) *VApp {
	return &VApp{title: title, x: 0, y: 0, width: 640, height: 480, visibility: false}
}
