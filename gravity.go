package gravity

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/thegtproject/gravity/internal/platform"
)

// Platform ...
var Platform platform.Platform

// Init ...
func Init(title string, width, height int) {
	println("Gravity.Init()")
	Platform = platform.New(title, width, height)
}

// Run ...
func Run(run func()) {
	println("Gravity.Run()")

	Platform.Run(run)
}

// Running ...
func Running() bool {
	return Platform.Running()
}

// Stop ...
func Stop() {
	Platform.Stop()
}

// Update ...
func Update() {
	Platform.Update()
}

// SetClearColor ...
func SetClearColor(color mgl32.Vec4) {
	Platform.SetClearColor(color)
}
