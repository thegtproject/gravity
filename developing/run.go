package main

import (
	"time"

	"github.com/thegtproject/gravity"
	gl "github.com/thegtproject/gravitygl"
)

var (
	cam *gravity.Camera

	terrain *gravity.Renderable
)

func run() {
	last := time.Now()
	for gravity.Running() {
		dt := float32(time.Since(last).Seconds())
		last = time.Now()
		cam.Update()
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		handleInput(dt)

		DefaultScene.Render()

		gravity.Update()
	}
}
