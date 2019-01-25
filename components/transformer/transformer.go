package transformer

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/thegtproject/gravity/common"
)

// Transformer ...
type Transformer struct {
	Mat         mgl32.Mat4
	Orientation mgl32.Quat
	Position    mgl32.Vec3
	Scale       mgl32.Vec3

	XMat common.XMat4
}

var (
	xAxis = mgl32.Vec3{1, 0, 0}
	yAxis = mgl32.Vec3{0, 1, 0}
	zAxis = mgl32.Vec3{0, 0, 1}
)

// NewTransformer ...
func NewTransformer() *Transformer {
	return &Transformer{
		Mat:         mgl32.Ident4(),
		Position:    mgl32.Vec3{0, 0, 0},
		Orientation: mgl32.QuatIdent(),
		Scale:       mgl32.Vec3{1, 1, 1},
	}
}

// Compose ...
func (tf *Transformer) Compose() *mgl32.Mat4 {
	tf.Mat = quatToMat4(tf.Orientation)
	scale(&tf.Mat, tf.Scale)
	setPosition(&tf.Mat, tf.Position)
	return &tf.Mat
}

// TranslateV ...
func (tf *Transformer) TranslateV(v mgl32.Vec3) {
	tf.Position = tf.Position.Add(v)
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
func (tf *Transformer) Rotate(q mgl32.Quat) {
	tf.Orientation = tf.Orientation.Mul(q)
	//tf.Compose()
}

func (tf *Transformer) rotateQ(rad float32, axis mgl32.Vec3) {
	tf.Orientation = tf.Orientation.Mul(mgl32.QuatRotate(rad, axis))
	//tf.Compose()
}

func scale(out *mgl32.Mat4, v mgl32.Vec3) *mgl32.Mat4 {
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

func setPosition(out *mgl32.Mat4, v mgl32.Vec3) {
	out[12] = v[0]
	out[13] = v[1]
	out[14] = v[2]
}

func quatToMat4(q mgl32.Quat) mgl32.Mat4 {
	n := q.Dot(q)

	var s float32

	if !mgl32.FloatEqual(n, 0) {
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

	m := mgl32.Mat4{
		1.0 - (yy + zz), xy + wz, xz - wy, 0,
		xy - wz, 1.0 - (xx + zz), yz + wx, 0,
		xz + wy, yz - wx, 1.0 - (xx + yy), 0,
		0, 0, 0, 1,
	}

	return m
}
