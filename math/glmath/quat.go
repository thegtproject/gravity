package glmath

import (
	"math"
)

// Quat ...
type Quat struct {
	W float32
	V Vec3
}

// Q ...
func Q() Quat {
	return Quat{
		W: 1,
		V: *V3(),
	}
}

// Multiply ...
func (quat *Quat) Multiply(q *Quat) *Quat {
	quat.V[0] = quat.V[0]*q.W + quat.W*q.V[0] + quat.V[1]*q.V[2] - quat.V[2]*q.V[1]
	quat.V[1] = quat.V[1]*q.W + quat.W*q.V[1] + quat.V[2]*q.V[0] - quat.V[0]*q.V[2]
	quat.V[2] = quat.V[2]*q.W + quat.W*q.V[2] + quat.V[0]*q.V[1] - quat.V[1]*q.V[0]
	quat.W = quat.W*q.W - quat.V[0]*q.V[0] - quat.V[1]*q.V[1] - quat.V[2]*q.V[2]
	return quat
}

// Dot ...
func (quat *Quat) Dot(q *Quat) float32 {
	return quat.W*q.W + quat.V[0]*q.V[0] + quat.V[1]*q.V[1] + quat.V[2]*q.V[2]
}

// RotateVec3 ...
func (quat *Quat) RotateVec3(v Vec3) Vec3 {
	tempv := quat.V.Clone()
	cross := tempv.CreateCross(v)
	c2 := tempv.Multiply(2).CreateCross(cross)
	return *v.Add(cross.Multiply(2 * quat.W)).Add(&c2)
}

// CreateQuatAngle ...
func CreateQuatAngle(angle float32, axis Vec3) Quat {
	q := Quat{}
	s := float32(math.Sin(float64(angle / 2)))
	q.V[0] = s * axis[0]
	q.V[1] = s * axis[1]
	q.V[2] = s * axis[2]
	q.W = float32(math.Cos(float64(angle / 2)))
	return q
}
