package main

import (
	"github.com/thegtproject/gravity"
	"github.com/thegtproject/gravity/internal/gravitygl"
	"github.com/thegtproject/gravity/materials"
	"github.com/thegtproject/gravity/math/mgl32"
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
	cam.Transformer.Position = mgl32.Vec3{-50, -110, 245}

	linewidget = gravity.NewModel(
		mesh.FromGob("assets/linewidget.gmesh").Scale(15),
		materials.NewNone(),
		cam,
	)
	linewidget.Transformer.Position = mgl32.Vec3{0, 0, 142}
	linewidget.Primitive = gravity.Lines
	linewidgetb = linewidget.Base()

	terrain = gravity.NewModel(
		mesh.FromGob("assets/terrain.gmesh").ScaleXYZ(50, 50, 50),
		materials.NewNone(),
		cam,
	)
	terrainb = terrain.Base()

	// cube = gravity.NewModel(mesh.NewCube().Scale(15), materials.NewNone(), cam)
	// cubeb = cube.Base()
	// cube.Primitive = gravitygl.TRIANGLES

	DefaultScene.Import(linewidget)
	DefaultScene.Import(terrain)
	// DefaultScene.Import(cube)

	run()
}

func setgloptions() {
	gravitygl.ClearColor(mgl32.Vec4{0.05, 0.05, 0.05, 1})
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
