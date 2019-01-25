package main

import (
	"fmt"

	"github.com/thegtproject/gravity"
)

func handleInput(dt float32) {
	if checkDebugCommand() {
		return
	}

	if gravity.Pressed(gravity.KeyEscape) {
		gravity.Stop()
	}
	if gravity.Pressed(gravity.Key1) {
		terrainb.Transformer.RotateZ(dt * 7)
	}
	if gravity.Pressed(gravity.Key2) {
		terrainb.Transformer.RotateZ(dt * -7)
	}
	if gravity.Pressed(gravity.Key3) {
		scaleDelta := gravity.Vec3{dt * 5, dt * 5, dt * 5}
		terrainb.Transformer.Scale.Add(&scaleDelta)
	}
	if gravity.Pressed(gravity.Key4) {
		scaleDelta := gravity.Vec3{dt * 5, dt * 5, dt * 5}
		terrainb.Transformer.Scale.Sub(&scaleDelta)
	}
	if gravity.Pressed(gravity.Key5) {
		scaleDelta := gravity.Vec3{dt * 15, dt * 15, dt * 15}
		linewidgetb.Transformer.Scale.Add(&scaleDelta)
	}
	if gravity.Pressed(gravity.Key6) {
		scaleDelta := gravity.Vec3{dt * 15, dt * 15, dt * 15}
		linewidgetb.Transformer.Scale.Sub(&scaleDelta)
	}

	if gravity.Pressed(gravity.KeyW) {
		cam.MoveForward(dt * 25)
	}
	if gravity.Pressed(gravity.KeyS) {
		cam.MoveBackward(dt * 25)

	}
	if gravity.Pressed(gravity.KeyA) {
		cam.MoveLeft(dt * 25)
	}
	if gravity.Pressed(gravity.KeyD) {
		cam.MoveRight(dt * 25)
	}
	if gravity.Pressed(gravity.KeySpace) {
		cam.MoveUp(dt * 25)
	}
	if gravity.Pressed(gravity.KeyLeftControl) {
		cam.MoveDown(dt * 25)
	}

	if gravity.Pressed(gravity.KeyLeft) {
		cam.Turn(dt * 25)
	}
	if gravity.Pressed(gravity.KeyRight) {
		cam.Turn(-dt * 25)
	}
	if gravity.Pressed(gravity.KeyUp) {
		cam.Roll(dt * 25)
	}
	if gravity.Pressed(gravity.KeyDown) {
		cam.Roll(-dt * 25)
	}

	if gravity.Pressed(gravity.KeyI) {
		terrainb.Transformer.RotateX(dt * -7)
	}
	if gravity.Pressed(gravity.KeyK) {
		terrainb.Transformer.RotateX(dt * 7)
	}
	if gravity.Pressed(gravity.KeyJ) {
		terrainb.Transformer.RotateY(dt * -7)
	}
	if gravity.Pressed(gravity.KeyL) {
		terrainb.Transformer.RotateY(dt * 7)
	}

	if gravity.JustPressed(gravity.KeyGrave) {
		debugCommandMode = true
		fmt.Print("debug command: ")
	}
}
