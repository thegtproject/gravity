package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/thegtproject/gravity"
	gl "github.com/thegtproject/gravitygl"

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

	DefaultScene.Camera(cam)
	terrain := DefaultScene.Import("terrain", mesh.FromGob("assets\\terrain.gmesh"), basicMaterial)

	DefaultScene.Import(
		"linewidget", mesh.FromGob("assets\\linewidget.gmesh"), basicMaterial,
	).Primitive = gl.LINES

	terrainBoundingBoxMesh := terrain.Mesh.GenerateBoundingBoxMeshWireframe()

	// gravity.MapFloat4Array(
	// 	&terrainBoundingBoxMesh.Colors,
	// 	func(r, g, b, a *float32) {
	// 		*a = 0.2
	// 	},
	// )

	DefaultScene.Import("terrainBoundingBoxMesh", terrainBoundingBoxMesh, basicMaterial).Primitive = gl.LINES

	terrainbb = DefaultScene.QueryObject("terrainBoundingBoxMesh")

	run()
}

func setgloptions() {
	gl.ClearColor(mgl32.Vec4{0.05, 0.05, 0.05, 1})
	gl.ClearDepth(1.0)
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
