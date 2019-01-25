package gravity

import (
	"github.com/thegtproject/gravity/components/transformer"
	"github.com/thegtproject/gravity/math/glmath"
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
	return &Camera{
		ProjectionMatrix: glmath.Perspective(glmath.DegToRad(80), 800/600, 0.1, 10000),
		Transformer:      transformer.NewTransformer(),
	}
}

// Update ...
func (cam *Camera) Update() {
	cam.Transformer.Compose()
	cam.ViewMatrix = *cam.Transformer.Mat.CreateInverse()
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
	return cam.Transformer.Orientation.RotateVec3(forward)
}

// GetRight ...
func (cam *Camera) GetRight() Vec3 {
	return cam.Transformer.Orientation.RotateVec3(xAxis)
}

// GetUp ...
func (cam *Camera) GetUp() Vec3 {
	return cam.Transformer.Orientation.RotateVec3(yAxis)
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
	v := cam.GetForward()
	v.Multiply(speed)
	cam.Transformer.TranslateV(&v)
}

// MoveBackward ...
func (cam *Camera) MoveBackward(speed float32) {
	v := cam.GetForward()
	v.Multiply(-speed)
	cam.Transformer.TranslateV(&v)
}

// MoveLeft ...
func (cam *Camera) MoveLeft(speed float32) {
	v := cam.GetRight()
	v.Multiply(-speed)
	cam.Transformer.TranslateV(&v)
}

// MoveRight ...
func (cam *Camera) MoveRight(speed float32) {
	v := cam.GetRight()
	v.Multiply(speed)
	cam.Transformer.TranslateV(&v)
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
func (cam *Camera) Rotate(angle float32, axis Vec3) {
	rot := glmath.CreateQuatAngle(angle, axis)
	cam.RotateQ(&rot)
}

// RotateQ ...
func (cam *Camera) RotateQ(rotation *Quat) {
	cam.Transformer.Orientation.Multiply(rotation)
}
