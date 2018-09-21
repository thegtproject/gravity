package main

import (
	"fmt"
	"time"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/thegtproject/gravity"
)

func run() {
	currentcol := mgl32.Vec4{0, 0.55, 0, 1}
	gravity.SetClearColor(currentcol)

	var frames uint64
	start := time.Now()
	lastfpscheck := time.Now()
	for gravity.Running() {

		if gravity.Pressed(gravity.KeyEscape) {
			gravity.Stop()
		}

		gravity.Update()
		frames++
		if time.Since(lastfpscheck).Seconds() >= 1 {
			x := time.Since(start).Seconds()
			gravity.SetTitle(
				fmt.Sprintf("Gravity Development - FPS: %d",
					uint(frames/uint64(x)),
				),
			)
			lastfpscheck = time.Now()
		}
	}
}

func main() {
	cfg := gravity.Config{
		Title: "Gravity Developing Application",
		Width: 800, Height: 600,
		MTXCallQueueCap: 3,
		VSync:           true,
	}
	gravity.Init(cfg)
	gravity.Run(run)
}
