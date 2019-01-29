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
	case gravity.KeyC:
		checkLast("Camera Location")
		fmt.Println("Position:")
		fmt.Println(cam.GetPosition())
		fmt.Println("Orientation:")

		fmt.Println("--------------------")
		fmt.Println(terrain.GetPosition())
		debugCommandMode = false
		gravity.Unpress(btn)
		btn = gravity.Button(0)
	case gravity.KeyEscape:
		checkLast("exit debug command mode")
		debugCommandMode = false
		gravity.Unpress(btn)
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
