package main

import (
	"fmt"

	"github.com/thegtproject/gravity"
	"github.com/thegtproject/gravity/math/mgl32"
)

func handleInput(dt float32) {
	if checkDebugCommand() {
		return
	}

	if gravity.Pressed(gravity.KeyEscape) {
		gravity.Stop()
	}

	if gravity.Pressed(gravity.KeyE) {
		cam.Debug()
	}

	if gravity.Pressed(gravity.Key1) {
		cubeb.Transformer.RotateZ(dt * 7)
	}
	if gravity.Pressed(gravity.Key2) {
		cubeb.Transformer.RotateZ(dt * -7)
	}
	if gravity.Pressed(gravity.Key3) {
		scaleDelta := mgl32.Vec3{dt * 5, dt * 5, dt * 5}
		cubeb.Transformer.Scale = cubeb.Transformer.Scale.Sub(scaleDelta)
	}
	if gravity.Pressed(gravity.Key4) {
		scaleDelta := mgl32.Vec3{dt * 5, dt * 5, dt * 5}
		cubeb.Transformer.Scale = cubeb.Transformer.Scale.Add(scaleDelta)
	}

	if gravity.Pressed(gravity.KeyW) {
		cam.MoveForward(dt * 15)
	}
	if gravity.Pressed(gravity.KeyS) {
		cam.MoveBackward(dt * 15)

	}
	if gravity.Pressed(gravity.KeyA) {
		cam.MoveLeft(dt * 15)
	}
	if gravity.Pressed(gravity.KeyD) {
		cam.MoveRight(dt * 15)
	}
	if gravity.Pressed(gravity.KeySpace) {
		cam.MoveUp(dt * 15)
	}
	if gravity.Pressed(gravity.KeyLeftControl) {
		cam.MoveDown(dt * 15)
	}

	if gravity.Pressed(gravity.KeyLeft) {
		cam.Rotate(0, dt*-75)
	}
	if gravity.Pressed(gravity.KeyRight) {
		cam.Rotate(0, dt*75)
	}
	if gravity.Pressed(gravity.KeyUp) {
		cam.Rotate(dt*75, 0)
	}
	if gravity.Pressed(gravity.KeyDown) {
		cam.Rotate(dt*-75, 0)
	}

	if gravity.Pressed(gravity.KeyI) {
		cubeb.Transformer.RotateX(dt * -7)
	}
	if gravity.Pressed(gravity.KeyK) {
		cubeb.Transformer.RotateX(dt * 7)
	}
	if gravity.Pressed(gravity.KeyJ) {
		cubeb.Transformer.RotateY(dt * -7)
	}
	if gravity.Pressed(gravity.KeyL) {
		cubeb.Transformer.RotateY(dt * 7)
	}
	if gravity.JustPressed(gravity.KeyGrave) {
		debugCommandMode = true
		fmt.Print("debug command: ")
	}
}
