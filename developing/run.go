package main

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/thegtproject/gravity"
	gl "github.com/thegtproject/gravity/internal/gravitygl"
)

var (
	cam *gravity.Camera

	terrain *gravity.Renderable
)

func run() {

	last := time.Now()
	start := time.Now()

	var frames uint64
	var timing time.Duration

	for gravity.Running() {

		dt := float32(time.Since(last).Seconds())
		last = time.Now()
		cam.Update()
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		handleInput(dt)

		DefaultScene.Render()

		gravity.Update()
		frames++

		timing = time.Since(start)
		// if  timing.Seconds() >= 3 {
		// 	gravity.Stop()
		// }
	}

	stats := &debug.GCStats{}
	debug.ReadGCStats(stats)
	fmt.Println("NumGC:     ", stats.NumGC)
	fmt.Println("PauseTotal:", stats.PauseTotal)
	fmt.Println("Frames:    ", frames)
	fmt.Println("Timing:    ", timing.Seconds())
	fmt.Println("Avg FPS:   ", float32(frames)/float32(timing.Seconds()))
}
