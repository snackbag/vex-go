package main

import (
	"fmt"
	"github.com/snackbag/vex-go/vex/core"
)

func TestVexOpenWin() {
	fmt.Println("[Test] Running test vex_openwin.go")

	app := core.NewApp("Test for Vex app")
	app.Show()
}
