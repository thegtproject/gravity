package main

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/thegtproject/gravity"
)

func run() {
	currentcol := mgl32.Vec4{0, 0.85, 0, 1}
	gravity.SetClearColor(currentcol)

	{
		geom := gravity.GeometryBuilders.TesselatedPlane()
		fmt.Println("Geom:\n", geom)
	}

	gravity.Stop()

	for gravity.Running() {
		if gravity.Pressed(gravity.KeyEscape) {
			gravity.Stop()
		}

		gravity.Update()
	}
}

func main() {
	gravity.Init("Gravity Developing Application", 800, 600)
	gravity.Run(run)
}
