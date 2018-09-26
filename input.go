package gravity

var (
	// MousePosition ...
	MousePosition = [2]float32{}

	inputTable            = make(map[Button]bool)
	inputJustPressedTable = make(map[Button]bool)
)

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

func onButtonPress(btn int) {
	inputTable[Button(btn)] = true
	inputJustPressedTable[Button(btn)] = true
}

func onButtonRelease(btn int) {
	inputTable[Button(btn)] = false
	inputJustPressedTable[Button(btn)] = false
}
