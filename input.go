package gravity

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

var (
	inputTable            = make(map[Button]bool)
	inputJustPressedTable = make(map[Button]bool)
)

// Mouse ...
var Mouse MouseInfo

// MouseInfo ...
type MouseInfo struct {
	Position [2]float32
	Delta    [2]float32
	Scroll   [2]float32

	capture bool
}

var _mouse MouseInfo

// OnMouseMove ...
var OnMouseMove = func(_ *glfw.Window, x, y float64) {
	_x, _y := float32(x), window.Height-float32(y)

	_mouse.Delta[0] = _x - Mouse.Position[0]
	_mouse.Delta[1] = _y - Mouse.Position[1]
	_mouse.Position[0] = _x
	_mouse.Position[1] = _y
}

// OnMouseScroll ...
var OnMouseScroll = func(_ *glfw.Window, xoff float64, yoff float64) {
	_mouse.Scroll[0] += float32(xoff)
	_mouse.Scroll[1] += float32(yoff)
}

// Unpress ...
func Unpress(btn Button) {
	inputTable[btn] = false
	inputJustPressedTable[btn] = false
}

// Pressed ...
func Pressed(btn Button) bool {
	return inputTable[btn]
}

// JustPressed ...
func JustPressed(btn Button) bool {
	if Pressed(btn) && inputJustPressedTable[btn] {
		inputJustPressedTable[btn] = false
		return true
	}
	return false
}

// UpdateInput ...
func UpdateInput() {
	Mouse = _mouse
	_mouse.Delta[0], _mouse.Delta[1] = 0, 0
	_mouse.Scroll[0], _mouse.Scroll[1] = 0, 0
}

func onButtonPress(btn int) {
	inputTable[Button(btn)] = true
	inputJustPressedTable[Button(btn)] = true
}

func onButtonRelease(btn int) {
	inputTable[Button(btn)] = false
	inputJustPressedTable[Button(btn)] = false
}
