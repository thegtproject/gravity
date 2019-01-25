package main

import (
	"github.com/thegtproject/gravity"
	"github.com/thegtproject/gravity/internal/gravitygl"
	"github.com/thegtproject/gravity/materials"
	"github.com/thegtproject/gravity/math/glmath"
	"github.com/thegtproject/gravity/mesh"
)

// DefaultScene ...
var DefaultScene gravity.Scene

var cam *gravity.Camera

// cheat variables for testing
var terrain *gravity.Model
var terrainb *gravity.BaseObject
var linewidget *gravity.Model
var linewidgetb *gravity.BaseObject

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

	linewidget = gravity.NewModel(
		mesh.FromGob("assets/linewidget.gmesh"),
		materials.NewNone(),
		cam,
	)
	linewidgetb = linewidget.Base()

	terrain = gravity.NewModel(
		mesh.FromGob("assets/terrain.gmesh"),
		materials.NewNone(),
		cam,
	)
	terrainb = terrain.Base()

	linewidget.Primitive = gravity.Lines

	DefaultScene.Import(testquad)
	DefaultScene.Import(linewidget)
	DefaultScene.Import(terrain)

	run()
}

func setgloptions() {
	gravitygl.ClearColor(glmath.Vec4{0.05, 0.05, 0.05, 1})
	gravitygl.ClearDepth(5.0)
	gravitygl.ViewPort(0, 0, int32(800), int32(600))
	gravitygl.Scissor(0, 0, int32(800), int32(600))
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
