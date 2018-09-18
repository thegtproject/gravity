package platform

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/thegtproject/gravity/internal/gravitygl"
)

// Platform ...
type Platform interface {
	Run(run func())
	Running() bool
	Update()
	SetClearColor(mgl32.Vec4)
	Stop()
}

// New ...
func New(title string, width int, height int) Platform {
	println("Platform.New()")
	platform := newPlatform(title, width, height)
	gravitygl.Init(platform.getMainThreadCallQueue())
	return platform
}
