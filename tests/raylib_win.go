package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func TestRaylibWin() {
	fmt.Println("[Test] Running test raylib_win.go")
	rl.SetTraceLogLevel(rl.LogError)

	rl.InitWindow(800, 450, "Window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("text", 10, 10, 20, rl.Black)

		rl.EndDrawing()
	}
}
