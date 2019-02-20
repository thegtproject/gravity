package gravity

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Window ...
type Window struct {
	Width, Height float32
	GlfwWin       *glfw.Window

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
		window.GlfwWin.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	} else {
		window.GlfwWin.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
	}
	Callbacks.CaptureModeOnChange(b)
}

func (win *Window) onFocusChange(w *glfw.Window, focused bool) {
	if !win.focused && win.Captured() {
		SetCaptureMode(false)
	}
	win.focused = focused
}

func (win *Window) glfwCursorDisabled() bool {
	return win.GlfwWin.GetInputMode(glfw.CursorMode) == glfw.CursorDisabled
}
