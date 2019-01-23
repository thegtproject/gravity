package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/thegtproject/gravity"
	gl "github.com/thegtproject/gravity/internal/gravitygl"

	"github.com/thegtproject/gravity/mesh"
)

// DefaultScene ...
var DefaultScene gravity.Scene

func setupscene() {
	setgloptions()

	basicMaterial, err := gl.MakeProgram(vsBasicShader, fsBasicShader)
	if err != nil {
		panic(err)
	}

	cam = gravity.NewCamera()
	cam.SetPosition(4, -13, 8)
	cam.LookAt(0, 0, 0)

	DefaultScene.SetCamera(cam)

	// pre generated terrain
	terrain = DefaultScene.Import("terrain", mesh.FromGob("assets/terrain.gmesh"), basicMaterial)

	// xyz line widget thing
	widget := DefaultScene.Import("linewidget", mesh.FromGob("assets/linewidget.gmesh"), basicMaterial)
	widget.Primitive = gl.LINES
	widget.Translate(0, 0, 5)

	run()
}

func setgloptions() {
	gl.ClearColor(mgl32.Vec4{0.05, 0.05, 0.05, 1})
	gl.ClearDepth(5.0)
	gl.ViewPort(0, 0, int32(800), int32(600))
	gl.Scissor(0, 0, int32(800), int32(600))
	gl.Enable(gl.GL_LINE_SMOOTH)
	gl.Hint(gl.GL_LINE_SMOOTH_HINT, gl.NICEST)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)
	gl.Enable(gl.BLEND)
	gl.Enable(gl.SCISSOR_TEST)
	gl.BlendEquation(gl.FUNC_ADD)
	gl.BlendFunc(gl.ONE, gl.ONE_MINUS_SRC_ALPHA)
	gl.LineWidth(2)
}
