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

// // QuatFromEuler ...
// func QuatFromEuler(alpha, beta, gamma float32) mgl32.Quat {
// 	return mgl32.Quat{
// 		W: Cos(alpha/2)*Cos(beta/2)*Cos(gamma/2) + Sin(alpha/2)*Sin(beta/2)*Sin(gamma/2),
// 		V: mgl32.Vec3{
// 			Sin(alpha/2)*Cos(beta/2)*Cos(gamma/2) - Cos(alpha/2)*Sin(beta/2)*Sin(gamma/2),
// 			Cos(alpha/2)*Sin(beta/2)*Cos(gamma/2) + Sin(alpha/2)*Cos(beta/2)*Sin(gamma/2),
// 			Cos(alpha/2)*Cos(beta/2)*Sin(gamma/2) - Sin(alpha/2)*Sin(beta/2)*Cos(gamma/2),
// 		},
// 	}
// }

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

// // Calculate Euler angles (roll, pitch, and yaw) from the unit quaternion.
// float roll = atan2(2.0f * (quat.w() * quat.x() + quat.y() * quat.z()),
// 	1.0f - 2.0f * (quat.x() * quat.x() + quat.y() * quat.y()));
// float pitch = asin(max(-1.0f, min(1.0f, 2.0f * (quat.w() * quat.y() - quat.z() * quat.x()))));
// float yaw = atan2(2.0f * (quat.w() * quat.z() + quat.x() * quat.y()),
// 	1.0f - 2.0f * (quat.y() * quat.y() + quat.z() * quat.z()));

// cam.front[0] = Cos(mgl32.DegToRad(cam.pitch)) * Cos(mgl32.DegToRad(cam.yaw))
// cam.front[1] = Sin(mgl32.DegToRad(cam.pitch))
// cam.front[2] = Cos(mgl32.DegToRad(cam.pitch)) * Sin(mgl32.DegToRad(cam.yaw))
