package main

import (
	"github.com/thegtproject/gravity"
)

func handleInput(dt float32) {
	if gravity.Pressed(gravity.KeyEscape) {
		gravity.Stop()
	}
	if gravity.Pressed(gravity.Key1) {
		terrain.RotateZ(dt * 7)
	}
	if gravity.Pressed(gravity.Key2) {
		terrain.RotateZ(dt * -7)
	}

	if gravity.Pressed(gravity.KeyW) {
		cam.MoveForward(dt * 10)
	}
	if gravity.Pressed(gravity.KeyS) {
		cam.MoveBackward(dt * 10)

	}
	if gravity.Pressed(gravity.KeyA) {
		cam.MoveLeft(dt * 10)
	}
	if gravity.Pressed(gravity.KeyD) {
		cam.MoveRight(dt * 10)
	}
	if gravity.Pressed(gravity.KeySpace) {
		cam.MoveUp(dt * 10)
	}
	if gravity.Pressed(gravity.KeyLeftControl) {
		cam.MoveDown(dt * 10)
	}
	if gravity.Pressed(gravity.KeyLeft) {
		cam.Turn(dt * 10)
	}
	if gravity.Pressed(gravity.KeyRight) {
		cam.Turn(-dt * 10)
	}
	if gravity.Pressed(gravity.KeyUp) {
		cam.Roll(dt * 10)
	}
	if gravity.Pressed(gravity.KeyDown) {
		cam.Roll(-dt * 10)
	}
}
