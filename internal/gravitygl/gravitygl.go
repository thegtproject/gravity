//+build !debug

package gravitygl

import (
	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var mainThreadCallQueue chan func()

// Init ...
func Init(callQueue chan func()) {
	err := gl.Init()
	if err != nil {
		panic(err)
	}
	mainThreadCallQueue = callQueue
}

// SetClearColor ...
func SetClearColor(col mgl32.Vec4) {
	mainThreadCallQueue <- func() {
		gl.ClearColor(col[0], col[1], col[2], col[3])
	}
}

// Enable ...
func Enable(option uint32) {
	mainThreadCallQueue <- func() {
		gl.Enable(option)
	}
}

// Disable ...
func Disable(option uint32) {
	mainThreadCallQueue <- func() {
		gl.Disable(option)
	}
}

// DepthFunc ...
func DepthFunc(xfunc uint32) {
	mainThreadCallQueue <- func() {
		gl.DepthFunc(xfunc)
	}
}

// BlendFunc ...
func BlendFunc(sfactor, dfactor uint32) {
	mainThreadCallQueue <- func() {
		gl.BlendFunc(sfactor, dfactor)
	}
}

// ViewPort ...
func ViewPort(x, y, width, height int32) {
	mainThreadCallQueue <- func() {
		gl.Viewport(x, y, width, height)
	}
}
