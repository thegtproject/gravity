package main

import (
	"fmt"

	"github.com/thegtproject/gravity/math/glmath"

	"github.com/thegtproject/gravity"
)

func handleInput(dt float32) {
	if checkDebugCommand() {
		return
	}

	if gravity.Pressed(gravity.KeyEscape) {
		gravity.Stop()
	}

	if gravity.JustPressed(gravity.KeyE) {
		m := cam.Transformer.Mat.TargetTo(cam.Transformer.Position, glmath.Vec3{0, 0, 0}, cam.GetUp())
		q := glmath.QuatFromRotationMat4(m)
		cam.Transformer.Orientation = q
		cam.Transformer.Compose()
	}

	if gravity.Pressed(gravity.Key1) {
		cubeb.Transformer.RotateZ(dt * 7)
	}
	if gravity.Pressed(gravity.Key2) {
		cubeb.Transformer.RotateZ(dt * -7)
	}
	if gravity.Pressed(gravity.Key3) {
		scaleDelta := gravity.Vec3{dt * 5, dt * 5, dt * 5}
		cubeb.Transformer.Scale.Add(&scaleDelta)
	}
	if gravity.Pressed(gravity.Key4) {
		scaleDelta := gravity.Vec3{dt * 5, dt * 5, dt * 5}
		cubeb.Transformer.Scale.Sub(&scaleDelta)
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
