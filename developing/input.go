package main

import (
	"fmt"

	"github.com/thegtproject/gravity"
	"github.com/thegtproject/gravity/internal/gravitygl"
	"github.com/thegtproject/gravity/math/mgl32"
)

var lastx, lasty float32
var fovy float32 = 55

var moveFactor = float32(55.0)
var polygonMode = gravitygl.FILL

func handleInput(dt float32) {
	if checkDebugCommand() {
		return
	}

	if !gravity.Captured() {
		if gravity.Pressed(gravity.MouseButton1) &&
			(lastx != gravity.Mouse.Delta[0] ||
				lasty != gravity.Mouse.Delta[1]) {
			cam.Rotate(dt*55*gravity.Mouse.Delta[1], -dt*55*gravity.Mouse.Delta[0])
		}
	} else {
		cam.Rotate(dt*15*gravity.Mouse.Delta[1], -dt*15*gravity.Mouse.Delta[0])
	}

	if gravity.JustPressed(gravity.KeyF1) {
		gravity.SetCaptureMode(!gravity.Captured())
	}

	if gravity.JustPressed(gravity.KeyF2) {
		if polygonMode == gravitygl.LINE {
			gravitygl.PolygonMode(gravitygl.FRONT_AND_BACK, gravitygl.FILL)
			polygonMode = gravitygl.FILL
		} else {
			gravitygl.PolygonMode(gravitygl.FRONT_AND_BACK, gravitygl.LINE)
			polygonMode = gravitygl.LINE
		}
	}

	if gravity.Mouse.Scroll[1] != 0 {
		if gravity.Pressed(gravity.KeyLeftAlt) {
			moveFactor += gravity.Mouse.Scroll[1] * 2
		} else {
			fovy += gravity.Mouse.Scroll[1]
			cam.ChangePerspective(fovy, 800/600, 0.1, 10000)
		}
	}

	if gravity.Pressed(gravity.KeyEscape) {
		gravity.Stop()
	}

	if gravity.Pressed(gravity.KeyE) {
		linewidget.Transformer.TranslateZ(1)
		fmt.Println(linewidget.Transformer.Position)
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
		cam.MoveForward(dt * moveFactor)
	}
	if gravity.Pressed(gravity.KeyS) {
		cam.MoveBackward(dt * moveFactor)

	}
	if gravity.Pressed(gravity.KeyA) {
		cam.MoveLeft(dt * moveFactor)
	}
	if gravity.Pressed(gravity.KeyD) {
		cam.MoveRight(dt * moveFactor)
	}
	if gravity.Pressed(gravity.KeySpace) {
		cam.MoveUp(dt * moveFactor)
	}
	if gravity.Pressed(gravity.KeyLeftControl) {
		cam.MoveDown(dt * moveFactor)
	}

	if gravity.Pressed(gravity.KeyLeft) {
		cam.Rotate(0, dt*120)
	}
	if gravity.Pressed(gravity.KeyRight) {
		cam.Rotate(0, dt*-120)
	}
	if gravity.Pressed(gravity.KeyUp) {
		cam.Rotate(dt*120, 0)
	}
	if gravity.Pressed(gravity.KeyDown) {
		cam.Rotate(dt*-120, 0)
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

	lastx = gravity.Mouse.Delta[0]
	lasty = gravity.Mouse.Delta[1]
}
