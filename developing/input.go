package main

import (
	"github.com/thegtproject/gravity"
	"github.com/thegtproject/gravity/internal/gravitygl"
)

var lastx, lasty float32

var moveFactor = float32(55.0)
var polygonMode = gravitygl.FILL

var skipinput = false

var yaww, pitchh float32

func handleInput(dt float32) {
	if gravity.Pressed(gravity.KeyLeftShift) {
		moveFactor = 255
	} else {
		moveFactor = 55
	}

	if gravity.Pressed(gravity.MouseButton2) {
		PX = gravity.Mouse.Position[0]
		PY = -(gravity.Mouse.Position[1] - gravity.GetWindow().Height)
		GUICamera.posOpened.X = PX
		GUICamera.posOpened.Y = PY
		Log.Println(PX, PY)
	}

	if !gravity.Captured() {
		if gravity.Pressed(gravity.MouseButton1) &&
			(lastx != gravity.Mouse.Delta[0] || lasty != gravity.Mouse.Delta[1]) {
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
			GUIConsole.Println("Wireframe: Off")
		} else {
			gravitygl.PolygonMode(gravitygl.FRONT_AND_BACK, gravitygl.LINE)
			polygonMode = gravitygl.LINE
			GUIConsole.Println("Wireframe: On")
		}
	}

	if gravity.Pressed(gravity.Key8) {
		linewidget.TranslateZ(75)
	}

	if gravity.Pressed(gravity.Key9) {
		linewidget.TranslateZ(-75)
	}
	if gravity.JustPressed(gravity.KeyEscape) {
		if gravity.Captured() {
			gravity.SetCaptureMode(!gravity.Captured())

		} else {
			gravity.Stop()
		}
	}

	if gravity.Pressed(gravity.Key1) {
		terrain.Scalef(1.1)
	}
	if gravity.Pressed(gravity.Key2) {
		terrain.Scalef(0.9)
	}
	if gravity.JustPressed(gravity.Key3) {
		Log.Println(cam.GetPosition())
		Log.Println(cam.GetRotation())
		Log.Println(terrain.GetScale())
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

	lastx = gravity.Mouse.Delta[0]
	lasty = gravity.Mouse.Delta[1]
}
