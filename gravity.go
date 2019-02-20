package gravity

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/thegtproject/gravity/internal/gravitygl"
	"github.com/thegtproject/gravity/pkg/log"
)

// Window ...
var window *Window
var running bool
var Log = log.GetRedirector()

// Config ...
type Config struct {
	Title         string
	Width, Height int
	VSync         bool
}

// Init ...
func Init(cfg Config) {
	Log.Print("Gravity - ", Version, "\n\n")
	Log.Println("Gravity.Init()")

	initglfw()
	createWindow(cfg.Title, cfg.Width, cfg.Height)
	window.Width = float32(cfg.Width)
	window.Height = float32(cfg.Height)
	initgl()

	gravitygl.ViewPort(0, 0, int32(cfg.Width), int32(cfg.Height))
	gravitygl.Scissor(0, 0, int32(cfg.Width), int32(cfg.Height))

	initCallbacks()
	if cfg.VSync {
		glfw.SwapInterval(1)
	} else {
		glfw.SwapInterval(0)
	}

	loadDefaultMaterialPrograms()
}

func initglfw() {
	if err := glfw.Init(); err != nil {
		panic("failed to initialize glfw: " + err.Error())
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCompatProfile)
	glfw.WindowHint(glfw.OpenGLDebugContext, glfw.True)
	glfw.WindowHint(glfw.Samples, 1)

}

func initgl() {
	err := gravitygl.Init()
	if err != nil {
		panic(err)
	}
}

// GetWindow ...
func GetWindow() *Window {
	return window
}

// Run ...
func Run(run func()) {
	Log.Println("Gravity.Run()")
	run()
}

// Running ...
func Running() bool {
	return !window.GlfwWin.ShouldClose()
}

// Stop ...
func Stop() {
	window.GlfwWin.SetShouldClose(true)
}

// Update ...
func Update() {
	window.GlfwWin.SwapBuffers()
	glfw.PollEvents()
	UpdateInput()
}

func initCallbacks() {
	window.GlfwWin.SetKeyCallback(
		func(_ *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mod glfw.ModifierKey) {
			if key == glfw.KeyUnknown {
				return
			}
			switch action {
			case glfw.Press:
				onButtonPress(int(key))
			case glfw.Release:
				onButtonRelease(int(key))
			}
		})
	window.GlfwWin.SetFocusCallback(window.onFocusChange)
	window.GlfwWin.SetCursorPosCallback(OnMouseMove)
	window.GlfwWin.SetMouseButtonCallback(func(_ *glfw.Window, btn glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
		switch action {
		case glfw.Press:
			onButtonPress(int(btn))
		case glfw.Release:
			onButtonRelease(int(btn))
		}
	})

	window.GlfwWin.SetScrollCallback(OnMouseScroll)

}
func createWindow(title string, width int, height int) {
	win, err := glfw.CreateWindow(
		width,
		height,
		title,
		nil,
		nil,
	)
	if err != nil {
		panic(err)
	}

	win.MakeContextCurrent()

	window = &Window{
		Width:   float32(width),
		Height:  float32(height),
		GlfwWin: win,
	}
}
