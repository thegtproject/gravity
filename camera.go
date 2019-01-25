package gravity

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/thegtproject/gravity/components/transformer"
)

var (
	xAxis = Vec3{1, 0, 0}
	yAxis = Vec3{0, 1, 0}
	zAxis = Vec3{0, 0, 1}
)

// Camera ...
type Camera struct {
	Transformer      *transformer.Transformer
	ProjectionMatrix Mat4
	ViewMatrix       Mat4
}

// NewCamera ...
func NewCamera() *Camera {
	cam := &Camera{
		ProjectionMatrix: mgl32.Perspective(mgl32.DegToRad(80), 800/600, 0.1, 10000),
		Transformer:      transformer.NewTransformer(),
	}
	return cam
}

// Update ...
func (cam *Camera) Update() {
	cam.Transformer.Compose()
	cam.ViewMatrix = cam.Transformer.Mat.Inv()
}

// Position ...
func (cam *Camera) Position() Vec3 {
	return cam.Transformer.Position
}

// SetPosition ...
func (cam *Camera) SetPosition(x, y, z float32) {
	cam.Transformer.Position[0] = x
	cam.Transformer.Position[1] = y
	cam.Transformer.Position[2] = z
}

// GetForward ...
func (cam *Camera) GetForward() Vec3 {
	forward := Vec3{0, 0, -1}
	return cam.Transformer.Orientation.Rotate(forward)
}

// GetRight ...
func (cam *Camera) GetRight() Vec3 {
	return cam.Transformer.Orientation.Rotate(xAxis)
}

// GetUp ...
func (cam *Camera) GetUp() Vec3 {
	return cam.Transformer.Orientation.Rotate(yAxis)
}

// MoveUp ...
func (cam *Camera) MoveUp(speed float32) {
	d := zAxis[2] * speed
	cam.Transformer.TranslateZ(d)
}

// MoveDown ...
func (cam *Camera) MoveDown(speed float32) {
	d := zAxis[2] * -speed
	cam.Transformer.TranslateZ(d)
}

// MoveForward ...
func (cam *Camera) MoveForward(speed float32) {
	v := cam.GetForward().Mul(speed)
	cam.Transformer.TranslateV(v)
}

// MoveBackward ...
func (cam *Camera) MoveBackward(speed float32) {
	v := cam.GetForward().Mul(-speed)
	cam.Transformer.TranslateV(v)
}

// MoveLeft ...
func (cam *Camera) MoveLeft(speed float32) {
	v := cam.GetRight().Mul(-speed)
	cam.Transformer.TranslateV(v)
}

// MoveRight ...
func (cam *Camera) MoveRight(speed float32) {
	v := cam.GetRight().Mul(speed)
	cam.Transformer.TranslateV(v)
}

// Yaw ...
func (cam *Camera) Yaw(rad float32) {
	cam.Rotate(rad, yAxis)
}

// Roll ...
func (cam *Camera) Roll(rad float32) {
	cam.Rotate(rad, xAxis)
}

// Pitch ...
func (cam *Camera) Pitch(rad float32) {
	cam.Rotate(rad, cam.GetRight())
}

// Turn ...
func (cam *Camera) Turn(rad float32) {
	cam.Rotate(rad, yAxis)
}

// Rotate ...
func (cam *Camera) Rotate(rad float32, axis Vec3) {
	cam.RotateQ(mgl32.QuatRotate(rad, axis))
}

// RotateQ ...
func (cam *Camera) RotateQ(rotation Quat) {
	cam.Transformer.Orientation = cam.Transformer.Orientation.Mul(rotation)
}

// SetProjection ...
func (cam *Camera) SetProjection(mat Mat4) {
	cam.ProjectionMatrix = mat
}
