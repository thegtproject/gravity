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

// QuatFromRotationMat4 ...
func QuatFromRotationMat4(m Mat4) Quat {
	q := Quat{}
	m11 := m[0]
	m12 := m[4]
	m13 := m[8]
	m21 := m[1]
	m22 := m[5]
	m23 := m[9]
	m31 := m[2]
	m32 := m[6]
	m33 := m[10]
	trace := m11 + m22 + m33

	var s float32
	if trace > 0 {
		s = 0.5 / Sqrt(trace+1.0)
		q.W = 0.25 / s
		q.V[0] = (m32 - m23) * s
		q.V[1] = (m13 - m31) * s
		q.V[2] = (m21 - m12) * s
	} else if m11 > m22 && m11 > m33 {
		s = 2.0 * Sqrt(1.0+m11-m22-m33)
		q.W = (m32 - m23) / s
		q.V[0] = 0.25 * s
		q.V[1] = (m12 + m21) / s
		q.V[2] = (m13 + m31) / s
	} else if m22 > m33 {
		s = 2.0 * Sqrt(1.0+m22-m11-m33)
		q.W = (m13 - m31) / s
		q.V[0] = (m12 + m21) / s
		q.V[1] = 0.25 * s
		q.V[2] = (m23 + m32) / s
	} else {
		s = 2.0 * Sqrt(1.0+m33-m11-m22)
		q.W = (m21 - m12) / s
		q.V[0] = (m13 + m31) / s
		q.V[1] = (m23 + m32) / s
		q.V[2] = 0.25 * s
	}
	return q
}
