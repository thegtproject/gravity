package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/thegtproject/gravity"

	gl "github.com/thegtproject/gravity/internal/gravitygl"
	"github.com/thegtproject/gravity/materials"
	"github.com/thegtproject/gravity/mesh"
)

// DefaultScene ...
var DefaultScene gravity.Scene

var cam *gravity.Camera
var terrain gravity.Object

func setupscene() {
	setgloptions()

	cam = gravity.NewCamera()
	cam.SetPosition(0, -13, 8)

	DefaultScene.SetCamera(cam)

	testquad := gravity.NewModel(
		mesh.NewQuad(),
		materials.NewSingleColor(
			gravity.Vec3{0, 0, 0},
			gravity.Vec3{0.7, 0.1, 0.1},
			gravity.Vec3{0, 0, 0},
			gravity.Vec3{0, 0, 0},
		),
		cam,
	)

	linewidget := gravity.NewModel(
		mesh.FromGob("assets/linewidget.gmesh"),
		materials.NewNone(),
		cam,
	)

	terrain = gravity.NewModel(
		mesh.FromGob("assets/terrain.gmesh"),
		materials.NewNone(),
		cam,
	)

	linewidget.Primitive = gravity.Lines

	DefaultScene.Import(testquad)
	DefaultScene.Import(linewidget)
	DefaultScene.Import(terrain)

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
