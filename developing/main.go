package main

import (
	"math/rand"

	"github.com/thegtproject/gravity"
)

func run() {
	currentcol := gravity.Vec4{0, 0.55, 0, 1}
	gravity.SetClearColor(currentcol)

	for gravity.Running() {

		if gravity.Pressed(gravity.KeyEscape) {
			gravity.Stop()
		}

		if gravity.Pressed(gravity.KeyUp) {
			gravity.SetClearColor(gravity.Vec4{rand.Float32(), rand.Float32(), rand.Float32(), 1.0})
		}

		gravity.Update()
	}
}

func main() {
	cfg := gravity.Config{
		Title: "Gravity Developing Application",
		Width: 800, Height: 600,
		VSync: true,
	}
	gravity.Init(cfg)
	gravity.Run(run)
}
