package main

import (
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/thegtproject/gravity"
)

func run() {
	currentcol := mgl32.Vec4{0, 0.85, 0, 1}

	gravity.SetClearColor(currentcol)
	for gravity.Running() {

		if gravity.Pressed(gravity.KeyEscape) {
			gravity.Stop()
		}

		if gravity.JustPressed(gravity.KeyC) {
			currentcol[0] = rand.Float32()
			currentcol[1] = rand.Float32()
			currentcol[2] = rand.Float32()
			currentcol[3] = 1.0
			gravity.SetClearColor(currentcol)
		}

		if gravity.Pressed(gravity.MouseButton1) {
			mousepos := gravity.MousePosition()
			currentcol[1] = mousepos[0] / 800
			gravity.SetClearColor(currentcol)
		}

		gravity.Update()
	}
}

func main() {
	gravity.Init("Gravity Developing Application", 800, 600)
	gravity.Run(run)
}
