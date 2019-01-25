package transformer

import (
	"github.com/thegtproject/gravity/math/glmath"
)

// Transformer ...
type Transformer struct {
	Mat         glmath.Mat4
	Orientation glmath.Quat
	Position    glmath.Vec3
	Scale       glmath.Vec3
}

var (
	xAxis = glmath.Vec3{1, 0, 0}
	yAxis = glmath.Vec3{0, 1, 0}
	zAxis = glmath.Vec3{0, 0, 1}
)

// NewTransformer ...
func NewTransformer() *Transformer {
	return &Transformer{
		Mat:         glmath.M4(),
		Position:    glmath.Vec3{0, 0, 0},
		Orientation: glmath.Q(),
		Scale:       glmath.Vec3{1, 1, 1},
	}
}

// Compose ...
func (tf *Transformer) Compose() *glmath.Mat4 {
	tf.Mat = quatToMat4(&tf.Orientation)
	scale(&tf.Mat, tf.Scale)
	setPosition(&tf.Mat, tf.Position)
	return &tf.Mat
}

// TranslateV ...
func (tf *Transformer) TranslateV(v *glmath.Vec3) {
	tf.Position.Add(v)
}

// TranslateX ...
func (tf *Transformer) TranslateX(d float32) {
	tf.Position[0] += d
}

// TranslateY ...
func (tf *Transformer) TranslateY(d float32) {
	tf.Position[1] += d
}

// TranslateZ ...
func (tf *Transformer) TranslateZ(d float32) {
	tf.Position[2] += d
}

// RotateX ...
func (tf *Transformer) RotateX(angle float32) {
	tf.rotateQ(angle, xAxis)
}

// RotateY ...
func (tf *Transformer) RotateY(angle float32) {
	tf.rotateQ(angle, yAxis)
}

// RotateZ ...
func (tf *Transformer) RotateZ(angle float32) {
	tf.rotateQ(angle, zAxis)
}

// Rotate ...
func (tf *Transformer) Rotate(q *glmath.Quat) {
	tf.Orientation.Multiply(q)
}

func (tf *Transformer) rotateQ(angle float32, axis glmath.Vec3) {
	q := glmath.CreateQuatAngle(angle, axis)
	tf.Orientation.Multiply(&q)
}

func scale(out *glmath.Mat4, v glmath.Vec3) *glmath.Mat4 {
	out[0] *= v[0]
	out[1] *= v[0]
	out[2] *= v[0]
	out[3] *= v[0]
	out[4] *= v[1]
	out[5] *= v[1]
	out[6] *= v[1]
	out[7] *= v[1]
	out[8] *= v[2]
	out[9] *= v[2]
	out[10] *= v[2]
	out[11] *= v[2]
	return out
}

func setPosition(out *glmath.Mat4, v glmath.Vec3) {
	out[12] = v[0]
	out[13] = v[1]
	out[14] = v[2]
}

func quatToMat4(q *glmath.Quat) glmath.Mat4 {
	n := q.Dot(q)

	var s float32

	if !glmath.Equal(n, 0) {
		s = 2.0 / n
	}

	xs := q.V[0] * s
	ys := q.V[1] * s
	zs := q.V[2] * s
	wx := q.W * xs
	wy := q.W * ys
	wz := q.W * zs
	xx := q.V[0] * xs
	xy := q.V[0] * ys
	xz := q.V[0] * zs
	yy := q.V[1] * ys
	yz := q.V[1] * zs
	zz := q.V[2] * zs

	m := glmath.Mat4{
		1.0 - (yy + zz), xy + wz, xz - wy, 0,
		xy - wz, 1.0 - (xx + zz), yz + wx, 0,
		xz + wy, yz - wx, 1.0 - (xx + yy), 0,
		0, 0, 0, 1,
	}

	return m
}
