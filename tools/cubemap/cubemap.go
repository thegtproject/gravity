package main

import (  
	"runtime"
 

	gl "github.com/thegtproject/gravity/internal/gravitygl"
	imgui "github.com/inkyblackness/imgui-go"
	"github.com/thegtproject/gravity"
	"github.com/thegtproject/gravity/pkg/math/mgl32"
)

var Log = gravity.Log

var (
	colorWindowBg           = imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0.4}
	colorTitleBar           = imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0.6}
	colorCollapseBtn        = imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0.6}
	colorCollapseBtnHovered = imgui.Vec4{X: 0.25, Y: 0, Z: 0, W: 0.9}
)

func runloop() {
	context := imgui.CreateContext(nil)
	imgui.CurrentIO().SetIniFilename("")
	defer context.Destroy()
	impl := imguiGlfw3Init(gravity.GetWindow().GlfwWin)
	defer impl.Shutdown()
	imgui.CurrentIO().Fonts().AddFontFromFileTTF("assets/fonts/SourceCodePro-Black.ttf", 16)
 

	io := imgui.CurrentIO()
	imgui.PushStyleColor(imgui.StyleColorWindowBg, colorWindowBg)
	imgui.PushStyleColor(imgui.StyleColorTitleBg, colorTitleBar)
	imgui.PushStyleColor(imgui.StyleColorTitleBgActive, colorTitleBar)
	imgui.PushStyleColor(imgui.StyleColorTitleBgCollapsed, colorTitleBar)
	imgui.PushStyleColor(imgui.StyleColorFrameBgHovered, colorCollapseBtnHovered)

	gl.ClearColor(mgl32.Vec4{0.06, 0.06, 0.06, 1})

	for gravity.Running() {

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		if io.WantCaptureMouse() == false || gravity.Captured() {
			// handleInput(dt)
		}

		// DefaultScene.Render()

		processgui(impl)

		gravity.Update()
		 
	}

	gravity.Update()
}

func processgui(impl *imguiGlfw3) {
	impl.NewFrame()

	imgui.EndFrame()
	imgui.Render()
	impl.Render(imgui.RenderedDrawData())
}

func setup() {

	runloop()
}

func main() {
	cfg := gravity.Config{
		Title: "Gravity Developing Application",
		Width: 1366, Height: 768,
		VSync: true,
	}

	Log = gravity.Log

	gravity.Init(cfg)
	gravity.Run(setup)
}

func init() {
	runtime.LockOSThread()
}
