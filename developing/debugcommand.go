package main

import (
	"fmt"

	"github.com/thegtproject/gravity"
)

var debugCommandMode bool
var lastDebugCommand gravity.Button

func processDebugCommandKey(btn gravity.Button) {
	checkLast := func(msg string) {
		if btn != lastDebugCommand {
			fmt.Println(msg)
		}
	}

	switch btn {
	case gravity.KeyGrave:
		checkLast("quick test `")

	case gravity.Key1:
		checkLast("quick test 1")
		// v := DefaultScene.QueryObject("terrain").Position()
		// cam.LookAt(v[0], v[1], v[2])
	case gravity.KeyC:
		checkLast("print camera information")
		pos := cam.Position()
		p := cam.ProjectionMatrix
		v := cam.ViewMatrix
		fmt.Printf(`
Position:   %v
ProjectionMatrix:
%s
ViewMatrix:
%s
`, pos, p, v)
	case gravity.KeyEscape:
		checkLast("exit debug command mode")
		debugCommandMode = false
		gravity.Unpress(btn)
	case gravity.KeyI:
		checkLast("print object information")
		// fmt.Println("Position: ", DefaultScene.QueryObject("terrain").Position())
	default:
		fmt.Println("unknown")
		fmt.Print("debug cmd: ")
		gravity.Unpress(btn)
	}

	lastDebugCommand = btn
}

func checkDebugCommand() (escape bool) {
	if !debugCommandMode {
		return false
	}

	for btn := gravity.Button(0); btn <= gravity.LastKey; btn++ {
		if gravity.JustPressed(btn) {
			processDebugCommandKey(btn)
			break
		}
	}
	return true
}
