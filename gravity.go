package gravity

import (
	"fmt"

	"github.com/go-gl/glfw/v3.2/glfw"

	gl "github.com/thegtproject/gravity/internal/gravitygl"
)

// Window ...
var Window *glfw.Window

var running bool

// Config ...
type Config struct {
	Title         string
	Width, Height int
	VSync         bool
}

// Init ...
func Init(cfg Config) {
	fmt.Print("Gravity - ", Version, "\n\n")
	println("Gravity.Init()")

	initglfw()
	createWindow(cfg.Title, cfg.Width, cfg.Height)
	initgl()
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
	err := gl.Init()
	if err != nil {
		panic(err)
	}
}

// Run ...
func Run(run func()) {
	println("Gravity.Run()")
	run()
}

// Running ...
func Running() bool {
	return !Window.ShouldClose()
}

// Stop ...
func Stop() {
	Window.SetShouldClose(true)
}

// Update ...
func Update() {
	Window.SwapBuffers()
	glfw.PollEvents()
}

func initCallbacks() {

	Window.SetKeyCallback(
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

	Window = win
}
