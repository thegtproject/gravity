package gravitygl

import (
	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/thegtproject/gravity/internal/gravitygl/opengl"
	"github.com/thegtproject/gravity/internal/schedulers/mtx"
)

// Init ...
func Init() {
	println("GravityGL<Desktop>.Init()")
	err := gl.Init()
	if err != nil {
		panic(err)
	}
}

// SetClearColor ...
func SetClearColor(col mgl32.Vec4) {
	mtx.Call(func() {
		gl.ClearColor(col[0], col[1], col[2], col[3])
	})
}

// Enable ...
func Enable(option uint32) {
	println("GravityGL.Enable()")
	mtx.Call(func() {
		gl.Enable(option)
	})
}

// Disable ...
func Disable(option uint32) {
	println("GravityGL.Disable()")
	mtx.Call(func() {
		gl.Disable(option)
	})
}

// DepthFunc ...
func DepthFunc(xfunc uint32) {
	println("GravityGL.DepthFunc()")
	mtx.Call(func() {
		gl.DepthFunc(xfunc)
	})
}

// BlendFunc ...
func BlendFunc(sfactor, dfactor uint32) {
	println("GravityGL.BlendFunc()")
	mtx.Call(func() {
		gl.BlendFunc(sfactor, dfactor)
	})
}

// ViewPort ...
func ViewPort(x, y, width, height int32) {
	println("GravityGL.ViewPort()")
	mtx.Call(func() {
		gl.Viewport(x, y, width, height)
	})
}

// GetGLVersion ...
func GetGLVersion() string {
	return opengl.GetString(
		opengl.VERSION,
	)
}
