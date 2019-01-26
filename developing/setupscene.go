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
var cube *gravity.Model
var cubeb *gravity.BaseObject
var terrain *gravity.Model
var terrainb *gravity.BaseObject
var linewidget *gravity.Model
var linewidgetb *gravity.BaseObject

func setupscene() {
	setgloptions()
	cam = gravity.NewCamera()
	DefaultScene.SetCamera(cam)
	cam.SetPosition(0, 0, 3)

	linewidget = gravity.NewModel(
		mesh.FromGob("assets/linewidget.gmesh"),
		materials.NewNone(),
		cam,
	)
	linewidget.Primitive = gravity.Lines
	linewidgetb = linewidget.Base()

	terrain = gravity.NewModel(
		mesh.FromGob("assets/terrain.gmesh"),
		materials.NewNone(),
		cam,
	)
	terrainb = terrain.Base()

	cube = gravity.NewModel(mesh.NewCube().Scale(15), materials.NewNone(), cam)
	cubeb = cube.Base()

	cube.Primitive = gravitygl.TRIANGLES

	DefaultScene.Import(linewidget)
	DefaultScene.Import(terrain)
	DefaultScene.Import(cube)

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
