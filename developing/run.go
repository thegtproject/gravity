package main

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/thegtproject/gravity"
	gl "github.com/thegtproject/gravity/internal/gravitygl"
	imgui "github.com/thegtproject/gravity/internal/imgui-go"
	"github.com/thegtproject/gravity/pkg/math/mgl32"
)

var (
	colorWindowBg           = imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0.4}
	colorTitleBar           = imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0.6}
	colorCollapseBtn        = imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0.6}
	colorCollapseBtnHovered = imgui.Vec4{X: 0.25, Y: 0, Z: 0, W: 0.9}
)

func run() {
	context := imgui.CreateContext(nil)
	defer context.Destroy()
	impl := imguiGlfw3Init(
		gravity.GetWindow().GlfwWin,
	)
	defer impl.Shutdown()
	imgui.CurrentIO().Fonts().AddFontFromFileTTF("assets/fonts/SourceCodePro-Regular.ttf", 14)

	last := time.Now()
	start := time.Now()

	var frames uint64
	var timing time.Duration

	io := imgui.CurrentIO()
	imgui.PushStyleColor(imgui.StyleColorWindowBg, colorWindowBg)
	imgui.PushStyleColor(imgui.StyleColorTitleBg, colorTitleBar)
	imgui.PushStyleColor(imgui.StyleColorTitleBgActive, colorTitleBar)
	imgui.PushStyleColor(imgui.StyleColorTitleBgCollapsed, colorTitleBar)
	imgui.PushStyleColor(imgui.StyleColorFrameBgHovered, colorCollapseBtnHovered)

	for gravity.Running() {
		gl.ClearColor(mgl32.Vec4{0.06, 0.06, 0.06, 1})
		dt := float32(time.Since(last).Seconds())
		last = time.Now()
		cam.Update()
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		if io.WantCaptureMouse() == false || gravity.Captured() {
			handleInput(dt)
		}

		DefaultScene.Render()

		processgui(impl)

		gravity.Update()
		frames++

		timing = time.Since(start)
		// if  timing.Seconds() >= 3 {
		// 	gravity.Stop()
		// }
	}

	stats := &debug.GCStats{}
	debug.ReadGCStats(stats)

	fmt.Println("")
	fmt.Println("Run Stats:")
	fmt.Println("-----------")
	fmt.Println("NumGC:     ", stats.NumGC)
	fmt.Println("PauseTotal:", stats.PauseTotal)
	fmt.Println("Frames:    ", frames)
	fmt.Println("Timing:    ", timing.Seconds())
	fmt.Println("Avg FPS:   ", float32(frames)/float32(timing.Seconds()))
}

func processgui(impl *imguiGlfw3) {
	impl.NewFrame()
	euler := QTE(cam.GetRotation())
	yaw, pitch := cam.GetYawPitch()
	GUICamera.buf.Reset()
	GUICamera.Printf(`Camera Orientation:

  Yaw:  %0.2f  Pitch: %0.2f
  Quaternion: 
  W: %0.2f X: %0.2f Y: %0.2f Z: %0.2f
  Euler:
  %0.2f Y: %0.2f Z: %0.2f

Camera Position:
  X: %0.2f Y: %0.2f Z: %0.2f
`,
		yaw, pitch, cam.GetRotation().W, cam.GetRotation().X(), cam.GetRotation().Y(), cam.GetRotation().Z(), euler[0], euler[1], euler[2], cam.GetPosition()[0], cam.GetPosition()[1], cam.GetPosition()[2],
	)

	GUIOutput.Render()
	GUIConsole.Render()
	GUICamera.Render()

	imgui.EndFrame()
	imgui.Render()
	impl.Render(imgui.RenderedDrawData())
}
