package gravity

import (
	"fmt"

	ggl "github.com/thegtproject/gravitygl"

	"github.com/go-gl/mathgl/mgl32"
)

var (
	Platform ggl.Platform
)

type Config struct {
	Title         string
	Width, Height int
	VSync         bool
}

// Init ...
func Init(cfg Config) {
	fmt.Print("Gravity - ", Version, "\n\n")
	println("Gravity.Init()")

	Platform = ggl.Init(ggl.Config{
		Title:        cfg.Title,
		Width:        cfg.Width,
		Height:       cfg.Height,
		VSync:        cfg.VSync,
		OnKeyPress:   onButtonPress,
		OnKeyRelease: onButtonRelease,
	})
}

// Run ...
func Run(run func()) {
	println("Gravity.Run()")

	Platform.Start(run)
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
	ggl.Update()
}

// SetClearColor ...
func SetClearColor(color mgl32.Vec4) {
	ggl.SetClearColor(color[0], color[1], color[2], color[3])
}

// SetTitle ...
func SetTitle(val string) {

}
