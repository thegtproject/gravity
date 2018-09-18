//+build !js

package input

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

// MouseButtonCallbackHandler ...
func MouseButtonCallbackHandler(_ *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	switch action {
	case glfw.Press:
		InputTable[Button(button)] = true
		InputJustPressedTable[Button(button)] = true
	case glfw.Release:
		InputTable[Button(button)] = false
		InputJustPressedTable[Button(button)] = false
	}
}

// KeyCallbackHandler ...
func KeyCallbackHandler(_ *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mod glfw.ModifierKey) {
	if key == glfw.KeyUnknown {
		return
	}
	switch action {
	case glfw.Press:
		InputTable[Button(key)] = true
		InputJustPressedTable[Button(key)] = true
	case glfw.Release:
		InputTable[Button(key)] = false
		InputJustPressedTable[Button(key)] = false
	}
}

// MousePositionHandler ...
func MousePositionHandler(_ *glfw.Window, x, y float64) {
	MousePosition[0] = float32(x)
	MousePosition[1] = float32(y)
}
