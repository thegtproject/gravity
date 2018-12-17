package main

import (
	"runtime"
	"time"

	"github.com/go-gl/mathgl/mgl32"
	imgui "github.com/inkyblackness/imgui-go"
	"github.com/thegtproject/gravity"
	gl "github.com/thegtproject/gravitygl"
)

var mm = mgl32.Ident4()
var val float32

func run() {
	context := imgui.CreateContext(nil)
	defer context.Destroy()
	impl := imguiGlfw3Init(gravity.Window)
	defer impl.Shutdown()
	cam := gravity.NewCamera()

	currentcol := gravity.Vec4{0, 0.55, 0, 1}
	gl.SetClearColor(currentcol[0], currentcol[1], currentcol[2], currentcol[3])

	indices := gl.NewUShortBuffer(gl.ELEMENT_ARRAY_BUFFER, gl.STATIC_DRAW, cubeindices)
	vertices := gl.NewFloatBuffer(gl.ARRAY_BUFFER, gl.STATIC_DRAW, cubevertices)
	cubevao := gl.NewVertexArrayObject(indices, vertices)

	program, err := gl.MakeProgram(vsNormalShading, fsNormalShading)
	if err != nil {
		panic(err)
	}

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)
	gl.ClearDepth(1.0)
	gl.ViewPort(0, 0, int32(800), int32(600))
	showDemoWin := true
	// start := time.Now()
	last := time.Now()
	for gravity.Running() {
		dt := float32(time.Since(last).Seconds())
		last = time.Now()
		cam.Update()
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.BindVertexArray(cubevao.ID)
		gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, 0)
		gl.EnableVertexAttribArray(0)

		gl.UseProgram(program)
		programMM := gl.GetUniformLocation(program, "uModelMatrix")
		programVM := gl.GetUniformLocation(program, "uViewMatrix")
		programPM := gl.GetUniformLocation(program, "uProjectionMatrix")

		gl.UniformMatrix4fv(programMM, mm)
		gl.UniformMatrix4fv(programVM, cam.ViewMatrix.Inv())
		gl.UniformMatrix4fv(programPM, cam.ProjectionMatrix)

		gl.BindVertexArray(cubevao.ID)

		gl.DrawElements(gl.TRIANGLES, len(cubeindices), gl.UNSIGNED_SHORT, 0)

		if gravity.Pressed(gravity.KeyEscape) {
			gravity.Stop()
		}

		if gravity.Pressed(gravity.KeyW) {
			cam.MoveForward(dt)

		}
		if gravity.Pressed(gravity.KeyS) {
			cam.MoveBackward(dt)

		}
		if gravity.Pressed(gravity.KeyA) {
			cam.MoveLeft(dt)

		}
		if gravity.Pressed(gravity.KeyD) {
			cam.MoveRight(dt)

		}
		if gravity.Pressed(gravity.KeySpace) {
			cam.MoveUp(dt)

		}
		if gravity.Pressed(gravity.KeyLeftControl) {
			cam.MoveDown(dt)

		}
		impl.NewFrame()
		imgui.ShowDemoWindow(&showDemoWin)
		imgui.Render()
		impl.Render(imgui.RenderedDrawData())
		gravity.Update()
	}
}

func main() {
	cfg := gravity.Config{
		Title: "Gravity Developing Application",
		Width: 800, Height: 600,
		VSync: true,
	}

	gravity.Init(cfg)
	gravity.Run(run)
}

func init() {
	runtime.LockOSThread()
}
