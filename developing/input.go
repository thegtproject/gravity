package main

import (
	"math"

	"github.com/westphae/quaternion"

	"github.com/thegtproject/gravity"
	"github.com/thegtproject/gravity/internal/gravitygl"
	"github.com/thegtproject/gravity/pkg/math/mgl32"
)

var lastx, lasty float32

var moveFactor = float32(55.0)
var polygonMode = gravitygl.FILL

var skipinput = false

var yaww, pitchh float32

func handleInput(dt float32) {
	if checkDebugCommand() {
		return
	}

	if gravity.Pressed(gravity.KeyLeftShift) {
		moveFactor = 255
	} else {
		moveFactor = 55
	}

	dothing := func(yaw, pitch float32) {
		y, p := cam.GetYawPitch()
		cam.DebugSetYawPitch(yaw+y, 0, pitch+p, 0)
	}

	if gravity.JustPressed(gravity.KeyF8) {
		dothing(45, 0)
	}
	if gravity.JustPressed(gravity.KeyF9) {
		dothing(-45, 0)
	}
	if gravity.JustPressed(gravity.KeyF10) {
		dothing(0, 45)
	}
	if gravity.JustPressed(gravity.KeyF11) {
		dothing(0, -45)
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
			GUIConsole.Println("Wireframe: Off")
		} else {
			gravitygl.PolygonMode(gravitygl.FRONT_AND_BACK, gravitygl.LINE)
			polygonMode = gravitygl.LINE
			GUIConsole.Println("Wireframe: On")
		}
	}

	if gravity.Pressed(gravity.Key8) {
		linewidget.TranslateZ(25)
	}

	if gravity.Pressed(gravity.Key9) {
		linewidget.TranslateZ(-25)
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

	if gravity.JustPressed(gravity.KeyT) {
		gravity.ShowUniform = true
	}

	if gravity.JustPressed(gravity.KeyQ) {
		Log.Println("--------")
		Log.Println("cam rot:  ", cam.GetRotation(), "  ", QTE(cam.GetRotation()))
		// Log.Println("quad rot: ", quad.GetRotation(), "  ", QTE(quad.GetRotation()))
	}

	//	quadrot = mgl32.AnglesToQuat(yaw, pitch, 0, mgl32.ZXZ).Mul(quadrot)
	//	terrain.SetRotation(quadrot)
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

	if gravity.JustPressed(gravity.KeyGrave) {
		debugCommandMode = true
		Log.Print("debug command: ")
	}

	lastx = gravity.Mouse.Delta[0]
	lasty = gravity.Mouse.Delta[1]
}

func deg2quat(yaw, pitch float32) mgl32.Quat {
	var s [3]float64
	var c [3]float64

	s[0], c[0] = 0, 1
	s[1], c[1] = math.Sincos(float64(yaw * float32(math.Pi) / 180 / 2))
	s[2], c[2] = math.Sincos(float64(pitch * float32(math.Pi) / 180 / 2))

	return mgl32.Quat{
		W: float32(c[0]*c[1]*c[2] - s[0]*c[1]*s[2]),
		V: mgl32.Vec3{
			float32(c[0]*c[1]*s[2] + s[0]*c[1]*c[2]),
			float32(c[0]*s[1]*s[2] - s[0]*s[1]*c[2]),
			float32(c[0]*s[1]*c[2] + s[0]*s[1]*s[2]),
		},
	}
}

func QTE(q mgl32.Quat) [3]float32 {
	qq := quaternion.Quaternion{
		W: float64(q.W),
		X: float64(q.V[0]),
		Y: float64(q.V[1]),
		Z: float64(q.V[2]),
	}
	a, b, c := quaternion.Euler(qq)
	return [3]float32{gravity.R2D(float32(a)), gravity.R2D(float32(b)), gravity.R2D(float32(c))}
}
