package geometry

import (
	"github.com/go-gl/mathgl/mgl32"
)

// Vec2 ...
type Vec2 = mgl32.Vec2

// V2Z returns a zero Vec2
var V2Z = Vec2{0, 0}

// V2 creates a new Vec2
func V2(x, y float32) Vec2 {
	return Vec2{x, y}
}

// Vec3 ...
type Vec3 = mgl32.Vec3

// V3Z returns a zero Vec3
var V3Z = Vec3{0, 0, 0}

// V3 creates a new Vec3
func V3(x, y, z float32) Vec3 {
	return Vec3{x, y, z}
}

// Vec4 ...
type Vec4 = mgl32.Vec4

// V4Z returns a zero Vec4
var V4Z = Vec4{0, 0, 0, 0}

// V4 creates a new Vec4
func V4(x, y, z, w float32) Vec4 {
	return Vec4{x, y, z, w}
}
