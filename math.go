package gravity

import (
	"math"

	"github.com/thegtproject/gravity/math/mgl32"
)

// Cos ...
func Cos(f float32) float32 {
	return float32(math.Cos(float64(f)))
}

// Sin ...
func Sin(f float32) float32 {
	return float32(math.Sin(float64(f)))
}

// Mod ...
func Mod(x, y float32) float32 {
	return float32(math.Mod(float64(x), float64(y)))
}

// Atan2 ...
func Atan2(x, y float32) float32 {
	return float32(math.Atan2(float64(y), float64(x)))
}

// Asin ...
func Asin(x float32) float32 {
	return float32(math.Asin(float64(x)))
}

// D2R ...
func D2R(d float32) float32 {
	return mgl32.DegToRad(d)
}

// R2D ...
func R2D(r float32) float32 {
	return mgl32.RadToDeg(r)
}

// SetQuatFromEuler ...
func SetQuatFromEuler(q *mgl32.Quat, v mgl32.Vec3) *mgl32.Quat {
	c1 := Cos(v[0] / 2)
	c2 := Cos(v[1] / 2)
	c3 := Cos(v[2] / 2)
	s1 := Sin(v[0] / 2)
	s2 := Sin(v[1] / 2)
	s3 := Sin(v[2] / 2)

	q.V[0] = s1*c2*c3 - c1*s2*s3
	q.V[1] = c1*s2*c3 + s1*c2*s3
	q.V[2] = c1*c2*s3 - s1*s2*c3
	q.W = c1*c2*c3 + s1*s2*s3

	return q
}

// Min ...
func Min(a, b float32) float32 {
	if a <= b {
		return a
	}
	return b
}

// Max ...
func Max(a, b float32) float32 {
	if a >= b {
		return a
	}
	return b
}

// Abs ...
func Abs(f float32) float32 {
	return float32(math.Abs(float64(f)))
}

// ScreenToGLCoords ...
func ScreenToGLCoords(width, height, x, y float32) (float32, float32) {
	return x / (width - 1), -1.0*y/(height-1) + 1.0
}
