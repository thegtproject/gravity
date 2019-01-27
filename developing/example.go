package main

import (
	"runtime"

	"github.com/thegtproject/gravity"
)

func main() {

	cfg := gravity.Config{
		Title: "Gravity Developing Application",
		Width: 1366, Height: 768,
		VSync: true,
	}

	gravity.Init(cfg)
	gravity.Run(setupscene)
}

func init() {
	runtime.LockOSThread()
}
