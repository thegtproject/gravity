package main

import (
	"fmt"
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/thegtproject/gravity"
)

func run() {
	currentcol := mgl32.Vec4{0, 0.85, 0, 1}
	gravity.SetClearColor(currentcol)

	plane := gravity.GeometryBuilders.TesselatedPlane()
	fmt.Println(plane)

	for gravity.Running() {
		if gravity.Pressed(gravity.KeyEscape) {
			gravity.Stop()
		}

		// Hold space to generate a bunch of MTX calls for profiling
		if gravity.Pressed(gravity.KeySpace) {
			currentcol[0] = rand.Float32()
			currentcol[1] = rand.Float32()
			currentcol[2] = rand.Float32()
			gravity.SetClearColor(currentcol)
		}

		gravity.Update()
	}
}

func main() {
	cfg := gravity.Config{
		Title: "Gravity Developing Application",
		Width: 800, Height: 600,
		MTXCallQueueCap: 3,
	}
	gravity.Init(cfg)
	gravity.Run(run)
}
