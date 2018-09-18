package input

import (
	"github.com/go-gl/mathgl/mgl32"
)

// InputTable ...
var InputTable = make(map[Button]bool)

// InputJustPressedTable ...
var InputJustPressedTable = make(map[Button]bool)

// MousePosition ...
var MousePosition = mgl32.Vec2{0, 0}

// Pressed ...
func Pressed(btn Button) bool {
	return InputTable[btn]
}

// JustPressed ...
func JustPressed(btn Button) bool {
	if Pressed(btn) && InputJustPressedTable[btn] {
		InputJustPressedTable[btn] = false
		return true
	}
	return false
}
