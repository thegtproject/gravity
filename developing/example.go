package main

import (
	"runtime"

	"github.com/thegtproject/gravity"
)

var Log = gravity.Log

func main() {

	cfg := gravity.Config{
		Title: "Gravity Developing Application",
		Width: 1366, Height: 768,
		VSync: true,
	}

	gravity.Log.Print = GUIOutput.Print
	gravity.Log.Printf = GUIOutput.Printf
	gravity.Log.Println = GUIOutput.Println
	Log = gravity.Log

	gravity.Init(cfg)
	gravity.Run(setupscene)
}

func init() {
	runtime.LockOSThread()
}
