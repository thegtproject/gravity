package main

import (
	"math/rand"

	"github.com/thegtproject/gravity"
	gl "github.com/thegtproject/gravitygl"
)

func run() {
	currentcol := gravity.Vec4{0, 0.55, 0, 1}
	gl.SetClearColor(currentcol[0], currentcol[1], currentcol[2], currentcol[3])

	for gravity.Running() {

		if gravity.Pressed(gravity.KeyEscape) {
			gravity.Stop()
		}

		if gravity.Pressed(gravity.KeyUp) {
			gl.SetClearColor(rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32())
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
