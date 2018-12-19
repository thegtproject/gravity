package gravity

import (
	"github.com/go-gl/mathgl/mgl32"
)

var (
	xaxis = mgl32.Vec3{1, 0, 0}
	yaxis = mgl32.Vec3{0, 1, 0}
	zaxis = mgl32.Vec3{0, 0, 1}
)

// Camera ...
type Camera struct {
	ProjectionMatrix mgl32.Mat4
	ViewMatrix       mgl32.Mat4
	position         mgl32.Vec3
	orientation      mgl32.Quat
}

// NewCamera ...
func NewCamera() *Camera {
	cam := &Camera{
		ProjectionMatrix: mgl32.Perspective(mgl32.DegToRad(75), 1, 0.1, 1000),
		ViewMatrix:       mgl32.Ident4(),
		position:         mgl32.Vec3{0, 4, 0},
		orientation:      mgl32.QuatIdent(),
	}

	cam.LookAt(mgl32.Vec3{0, 0, 0})

	return cam
}

// LookAt ...
func (cam *Camera) LookAt(v mgl32.Vec3) {
	cam.orientation = mgl32.QuatLookAtV(cam.position, v, yaxis).Inverse()

}

// Update ...
func (cam *Camera) Update() {
	m := mgl32.Ident4()

	qm := cam.orientation.Mat4()
	translate(&m, cam.position)

	cam.ViewMatrix = m.Mul4(qm)
}

// GetForward ...
func (cam *Camera) GetForward() mgl32.Vec3 {
	v := mgl32.Vec3{0, 0, -1}
	return cam.orientation.Rotate(v)
}

// GetLeft ...
func (cam *Camera) GetLeft() mgl32.Vec3 {
	v := mgl32.Vec3{-1, 0, 0}
	return cam.orientation.Rotate(v)
}

// GetUp ...
func (cam *Camera) GetUp() mgl32.Vec3 {
	v := mgl32.Vec3{0, 1, 0}
	return cam.orientation.Rotate(v)
}

// MoveUp ...
func (cam *Camera) MoveUp(speed float32) {
	v := cam.GetUp().Mul(speed)
	cam.position = cam.position.Add(v)
}

// MoveDown ...
func (cam *Camera) MoveDown(speed float32) {
	v := cam.GetUp().Mul(speed)
	cam.position = cam.position.Sub(v)
}

// MoveForward ...
func (cam *Camera) MoveForward(speed float32) {
	v := cam.GetForward().Mul(speed)
	cam.position = cam.position.Add(v)
}

// MoveBackward ...
func (cam *Camera) MoveBackward(speed float32) {
	v := cam.GetForward().Mul(speed)
	cam.position = cam.position.Sub(v)
}

// MoveLeft ...
func (cam *Camera) MoveLeft(speed float32) {
	v := cam.GetLeft().Mul(speed)
	cam.position = cam.position.Add(v)
}

// MoveRight ...
func (cam *Camera) MoveRight(speed float32) {
	v := cam.GetLeft().Mul(speed)
	cam.position = cam.position.Sub(v)
}

// Yaw ...
func (cam *Camera) Yaw(rad float32) {
	cam.Rotate(rad, yaxis)
}

// Roll ...
func (cam *Camera) Roll(rad float32) {
	cam.Rotate(rad, xaxis)
}

// Pitch ...
func (cam *Camera) Pitch(rad float32) {
	cam.Rotate(rad, zaxis)
}

// Turn ...
func (cam *Camera) Turn(rad float32) {
	cam.Rotate(rad, yaxis)
}

// Rotate ...
func (cam *Camera) Rotate(rad float32, axis mgl32.Vec3) {
	cam.RotateQ(mgl32.QuatRotate(rad, axis))
}

// RotateQ ...
func (cam *Camera) RotateQ(rotation mgl32.Quat) {
	cam.orientation = cam.orientation.Mul(rotation)
}

// Projection ...
func (cam *Camera) Projection(mat mgl32.Mat4) {
	cam.ProjectionMatrix = mat
}

func translate(out *mgl32.Mat4, v mgl32.Vec3) {
	x, y, z := v[0], v[1], v[2]

	out[12] = out[0]*x + out[4]*y + out[8]*z + out[12]
	out[13] = out[1]*x + out[5]*y + out[9]*z + out[13]
	out[14] = out[2]*x + out[6]*y + out[10]*z + out[14]
	out[15] = out[3]*x + out[7]*y + out[11]*z + out[15]

}
