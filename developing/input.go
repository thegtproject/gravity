package main

import (
	"fmt"

	"github.com/thegtproject/gravity"
	"github.com/thegtproject/gravity/internal/gravitygl"
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
			cam.Push(dt*55*gravity.Mouse.Delta[1], -dt*55*gravity.Mouse.Delta[0])
		}
	} else {
		cam.Push(dt*15*gravity.Mouse.Delta[1], -dt*15*gravity.Mouse.Delta[0])
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

	if gravity.JustPressed(gravity.KeyEscape) {
		if gravity.Captured() {
			gravity.SetCaptureMode(!gravity.Captured())

		} else {
			gravity.Stop()
		}
	}

	if gravity.Pressed(gravity.KeyE) {
		linewidget.TranslateZ(1)
	}

	if gravity.JustPressed(gravity.KeyQ) {

	}

	if gravity.Pressed(gravity.Key1) {

	}
	if gravity.Pressed(gravity.Key2) {

	}
	if gravity.Pressed(gravity.Key3) {
		cubeb.Scalef(-0.1)
	}
	if gravity.Pressed(gravity.Key4) {
		cubeb.Scalef(0.1)
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

	var zz float32 = 10
	if gravity.JustPressed(gravity.KeyLeft) {
		cam.Push(0, zz)
	}
	if gravity.JustPressed(gravity.KeyRight) {
		cam.Push(0, -zz)
	}
	if gravity.JustPressed(gravity.KeyUp) {
		cam.Push(-zz, 0)
	}
	if gravity.JustPressed(gravity.KeyDown) {
		cam.Push(zz, 0)
	}

	if gravity.JustPressed(gravity.KeyGrave) {
		debugCommandMode = true
		fmt.Print("debug command: ")
	}

	lastx = gravity.Mouse.Delta[0]
	lasty = gravity.Mouse.Delta[1]
}
