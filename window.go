package gravity

import (
	"fmt"

	"github.com/go-gl/glfw/v3.2/glfw"
)

// Window ...
type Window struct {
	Width, Height float32
	glfwWin       *glfw.Window

	focused     bool
	capture     bool
	captureHold bool
}

// Captured ...
func Captured() bool {
	return window.Captured()
}

// Captured ...
func (win *Window) Captured() bool {
	return win.capture
}

// SetCaptureMode ...
func SetCaptureMode(b bool) {
	if b {
		window.glfwWin.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	} else {
		window.glfwWin.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
	}
	window.capture = b
	Mouse.capture = b
}

func (win *Window) onFocusChange(w *glfw.Window, focused bool) {
	win.focused = focused

	if !win.focused && win.capture {
		fmt.Println("pausing capture mode")
		win.captureHold = true
		SetCaptureMode(false)
		return
	}

	if win.focused && win.captureHold {
		fmt.Println("resuming capture mode")
		win.captureHold = false
		SetCaptureMode(true)
	}
}
