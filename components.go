package gravity

import (
	"github.com/go-gl/mathgl/mgl32"
)

// Transformer ...
type Transformer struct {
	matrix mgl32.Mat4
	rot    mgl32.Quat
	pos    mgl32.Vec3
	scale  mgl32.Vec3
}

var (
	xAxis = mgl32.Vec3{1, 0, 0}
	yAxis = mgl32.Vec3{0, 1, 0}
	zAxis = mgl32.Vec3{0, 0, 1}
)

// NewTransformer ...
func NewTransformer() *Transformer {
	return &Transformer{
		matrix: mgl32.Ident4(),
		pos:    mgl32.Vec3{0, 0, 0},
		scale:  mgl32.Vec3{1, 1, 1},
		rot:    mgl32.QuatIdent(),
	}
}

// UpdateTransform ...
func (t *Transformer) UpdateTransform() {
	t.matrix = t.rot.Mat4()
	t.matrix[12] = t.pos.X()
	t.matrix[13] = t.pos.Y()
	t.matrix[14] = t.pos.Z()
	scale(&t.matrix, t.scale)
}

// GetTransformMatrix ...
func (t *Transformer) GetTransformMatrix() *mgl32.Mat4 {
	return &t.matrix
}

// GetRotation ...
func (t *Transformer) GetRotation() mgl32.Quat {
	return t.rot
}

// GetPosition ...
func (t *Transformer) GetPosition() mgl32.Vec3 {
	return t.pos
}

// SetPosition ...
func (t *Transformer) SetPosition(x, y, z float32) {
	t.pos[0], t.pos[1], t.pos[2] = x, y, z
}

// SetRotation ...
func (t *Transformer) SetRotation(q mgl32.Quat) {
	t.rot = q
}

// Translate ...
func (t *Transformer) Translate(v mgl32.Vec3) {
	t.pos = t.pos.Add(v)
}

// TranslateX ...
func (t *Transformer) TranslateX(d float32) {
	t.pos[0] += d
}

// TranslateY ...
func (t *Transformer) TranslateY(d float32) {
	t.pos[1] += d
}

// TranslateZ ...
func (t *Transformer) TranslateZ(d float32) {
	t.pos[2] += d
}

// // RotateX ...
// func (t *Transformer) RotateX(angle float32) {

// }

// // RotateY ...
// func (t *Transformer) RotateY(angle float32) {

// }

// // RotateZ ...
// func (t *Transformer) RotateZ(angle float32) {

// }

// Rotate ...
// func (t *Transformer) Rotate(q mgl32.Quat) {
// 	t.rot = q.Mul(t.rot)
// }

// GetScale ...
func (t *Transformer) GetScale() mgl32.Vec3 {
	return t.scale
}

// Scale ...
func (t *Transformer) Scale(v mgl32.Vec3) {
	t.scale = v
}

// Scalef ...
func (t *Transformer) Scalef(f float32) {
	t.scale = t.scale.Mul(f)
}

// ScaleX ...
func (t *Transformer) ScaleX(f float32) {
	t.scale[0] *= f
}

// ScaleY ...
func (t *Transformer) ScaleY(f float32) {
	t.scale[1] *= f
}

// ScaleZ ...
func (t *Transformer) ScaleZ(f float32) {
	t.scale[2] *= f
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
