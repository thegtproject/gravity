package main

import (
	"github.com/thegtproject/gravity"
	"github.com/thegtproject/gravity/internal/gravitygl"
)

// DefaultScene ...
var DefaultScene gravity.Scene

var cam *gravity.Camera

var terrain *gravity.Model
var skybox *gravity.Model
var linewidget *gravity.Model

func setupscene() {
	setgloptions()
	setcallbacks()

	cam = DefaultScene.SetCamera(
		gravity.NewCamera(
			// camera options
			gravity.Position(0, 0, 950),
		))
	// cam.Push(90, 0)

	configureSkybox()
	DefaultScene.Import(skybox)

	configureTerrain()
	DefaultScene.Import(terrain)

	configureLinewidget()
	DefaultScene.Import(linewidget)

	run()
}

func setgloptions() {

	gravitygl.ClearDepth(5.0)
	gravitygl.Enable(gravitygl.GL_LINE_SMOOTH)
	gravitygl.Hint(gravitygl.GL_LINE_SMOOTH_HINT, gravitygl.NICEST)
	gravitygl.Enable(gravitygl.DEPTH_TEST)
	gravitygl.DepthFunc(gravitygl.LEQUAL)
	gravitygl.Enable(gravitygl.BLEND)
	gravitygl.Enable(gravitygl.SCISSOR_TEST)
	gravitygl.BlendEquation(gravitygl.FUNC_ADD)
	gravitygl.BlendFunc(gravitygl.ONE, gravitygl.ONE_MINUS_SRC_ALPHA)
	gravitygl.LineWidth(2)

}

func setcallbacks() {
	gravity.Callbacks.CaptureModeOnChange = func(b bool) {
		GUIConsole.Println("Mouse Capture Mode:", b)
	}
}
