package gravity

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/thegtproject/gravity/internal/platform"
	"github.com/thegtproject/gravity/internal/schedulers/mtx"
)

// Platform ...
var Platform platform.Platform

// Config ...
type Config struct {
	Title         string
	Width, Height int

	// Limit the size of the MTX main thread scheduler queue
	// Default is 3
	MTXCallQueueCap int
}

// Init ...
func Init(cfg Config) {
	fmt.Print("Gravity - ", Version, "\n\n")
	println("Gravity.Init()")
	mtx.Init(cfg.MTXCallQueueCap)
	Platform = platform.New(cfg.Title, cfg.Width, cfg.Height)
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
