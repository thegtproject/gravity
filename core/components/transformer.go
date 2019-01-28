package components

import (
	"github.com/thegtproject/gravity/math/mgl32"
)

// Transformer ...
type Transformer struct {
	Mat         mgl32.Mat4
	Orientation mgl32.Quat
	Position    mgl32.Vec3
	Scale       mgl32.Vec3

	originalOrientation mgl32.Quat
}

var (
	xAxis = mgl32.Vec3{1, 0, 0}
	yAxis = mgl32.Vec3{0, 1, 0}
	zAxis = mgl32.Vec3{0, 0, 1}
)

// NewTransformer ...
func NewTransformer(base mgl32.Quat) Transformer {
	return Transformer{
		Mat:                 mgl32.Ident4(),
		Position:            mgl32.Vec3{0, 0, 0},
		Orientation:         mgl32.QuatIdent(),
		Scale:               mgl32.Vec3{1, 1, 1},
		originalOrientation: base,
	}
}

// Compose ...
func (t *Transformer) Compose() {
	t.Mat = t.Orientation.Mat4()
	scale(&t.Mat, t.Scale)
	setPosition(&t.Mat, t.Position)
}

// TranslateV ...
func (t *Transformer) TranslateV(v mgl32.Vec3) {
	t.Position = t.Position.Add(v)
}

// TranslateX ...
func (t *Transformer) TranslateX(d float32) {
	t.Position[0] += d
}

// TranslateY ...
func (t *Transformer) TranslateY(d float32) {
	t.Position[1] += d
}

// TranslateZ ...
func (t *Transformer) TranslateZ(d float32) {
	t.Position[2] += d
}

// RotateX ...
func (t *Transformer) RotateX(angle float32) {
	t.rotateQ(angle, xAxis)
}

// RotateY ...
func (t *Transformer) RotateY(angle float32) {
	t.rotateQ(angle, yAxis)
}

// RotateZ ...
func (t *Transformer) RotateZ(angle float32) {
	t.rotateQ(angle, zAxis)
}

// Rotate ...
func (t *Transformer) Rotate(q mgl32.Quat) {
	t.Orientation = q.Mul(t.originalOrientation)
}

func (t *Transformer) rotateQ(angle float32, axis mgl32.Vec3) {
	q := mgl32.QuatRotate(angle, axis)
	t.Orientation = t.Orientation.Mul(q)
}

func scale(out *mgl32.Mat4, v mgl32.Vec3) {
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
