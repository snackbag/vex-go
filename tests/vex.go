package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetTraceLogLevel(rl.LogError)

	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	vex.Cool()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
