package gravity

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Window ...
type Window struct {
	Width, Height float32
	glfwWin       *glfw.Window

	focused bool
}

// Captured ...
func Captured() bool {
	return window.glfwCursorDisabled()
}

// Captured ...
func (win *Window) Captured() bool {
	return window.glfwCursorDisabled()
}

// SetCaptureMode ...
func SetCaptureMode(b bool) {
	if b {
		window.glfwWin.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	} else {
		window.glfwWin.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
	}
}

func (win *Window) onFocusChange(w *glfw.Window, focused bool) {
	if !win.focused && win.Captured() {
		SetCaptureMode(false)
	}
	win.focused = focused
}

func (win *Window) glfwCursorDisabled() bool {
	return win.glfwWin.GetInputMode(glfw.CursorMode) == glfw.CursorDisabled
}
