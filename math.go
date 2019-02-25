package gravity

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
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
func Atan2(y, x float32) float32 {
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

// Sincos ...
func Sincos(x float32) (sin, cos float32) {
	s, c := math.Sincos(float64(x))
	return float32(s), float32(c)
}

// Norm ...
func Norm(q mgl32.Quat) float32 {
	return float32(math.Sqrt(float64(
		q.W*q.W + q.V[0]*q.V[0] + q.V[1]*q.V[1] + q.V[2]*q.V[2],
	)))
}

// Unit ...
func Unit(q mgl32.Quat) mgl32.Quat {
	k := Norm(q)
	return mgl32.Quat{
		W: q.W / k,
		V: mgl32.Vec3{
			q.V[0] / k,
			q.V[1] / k,
			q.V[2] / k,
		},
	}
}

// QuatToEuler ...
func QuatToEuler(q mgl32.Quat) (float32, float32, float32) {
	r := Unit(q)
	phi := Atan2(2*(r.W*r.V[0]+r.V[1]*r.V[2]), 1-2*(r.V[0]*r.V[0]+r.V[1]*r.V[1]))
	theta := Asin(2 * (r.W*r.V[1] - r.V[2]*r.V[0]))
	psi := Atan2(2*(r.V[0]*r.V[1]+r.W*r.V[2]), 1-2*(r.V[1]*r.V[1]+r.V[2]*r.V[2]))
	return phi, theta, psi
}

// QuatToEulerDeg ...
func QuatToEulerDeg(q mgl32.Quat) (float32, float32, float32) {
	r := Unit(q)
	phi := Atan2(2*(r.W*r.V[0]+r.V[1]*r.V[2]), 1-2*(r.V[0]*r.V[0]+r.V[1]*r.V[1]))
	theta := Asin(2 * (r.W*r.V[1] - r.V[2]*r.V[0]))
	psi := Atan2(2*(r.V[0]*r.V[1]+r.W*r.V[2]), 1-2*(r.V[1]*r.V[1]+r.V[2]*r.V[2]))
	return R2D(phi), R2D(theta), R2D(psi)
}

// Deg2Quat ...
func Deg2Quat(yaw, pitch float32) mgl32.Quat {
	var s [3]float64
	var c [3]float64

	s[0], c[0] = 0, 1
	s[1], c[1] = math.Sincos(float64(yaw * float32(math.Pi) / 180 / 2))
	s[2], c[2] = math.Sincos(float64(pitch * float32(math.Pi) / 180 / 2))

	return mgl32.Quat{
		W: float32(c[0]*c[1]*c[2] - s[0]*c[1]*s[2]),
		V: mgl32.Vec3{
			float32(c[0]*c[1]*s[2] + s[0]*c[1]*c[2]),
			float32(c[0]*s[1]*s[2] - s[0]*s[1]*c[2]),
			float32(c[0]*s[1]*c[2] + s[0]*s[1]*s[2]),
		},
	}
}
