package gravity

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

var (
	xaxis = Vec3{1, 0, 0}
	yaxis = Vec3{0, 1, 0}
	zaxis = Vec3{0, 0, 1}
)

// Camera ...
type Camera struct {
	ProjectionMatrix Mat4
	ViewMatrix       Mat4
	CamMatrix        Mat4
	position         Vec3
	orientation      Quat
}

// NewCamera ...
func NewCamera() *Camera {
	cam := &Camera{
		ProjectionMatrix: mgl32.Perspective(mgl32.DegToRad(80), 800/600, 0.1, 10000),
		CamMatrix:        mgl32.Ident4(),
		position:         Vec3{0, 0, 5},
		orientation:      mgl32.QuatIdent(),
	}

	cam.LookAt(0, 0, 0)
	cam.Update()
	return cam
}

// LookAt ...
func (cam *Camera) LookAt(x, y, z float32) {
	cam.orientation = mgl32.QuatLookAtV(cam.position, Vec3{x, y, z}, zaxis).Inverse()
	cam.Update()
}

// Update ...
func (cam *Camera) Update() {
	cam.CamMatrix = mgl32.Ident4()
	fromRotationTranslation(&cam.CamMatrix, cam.orientation, cam.position)
	cam.ViewMatrix = cam.CamMatrix.Inv()
}

// Position ...
func (cam *Camera) Position() Vec3 {
	return cam.position
}

// SetPosition ...
func (cam *Camera) SetPosition(x, y, z float32) {
	cam.position[0] = x
	cam.position[1] = y
	cam.position[2] = z
}

// GetForward ...
func (cam *Camera) GetForward() Vec3 {
	v := Vec3{0, 0, -1}
	return cam.orientation.Rotate(v)
}

// GetLeft ...
func (cam *Camera) GetLeft() Vec3 {
	v := Vec3{1, 0, 0}
	return cam.orientation.Rotate(v)
}

// GetUp ...
func (cam *Camera) GetUp() Vec3 {
	v := Vec3{0, 1, 0}
	return cam.orientation.Rotate(v)
}

// MoveUp ...
func (cam *Camera) MoveUp(speed float32) {
	v := zaxis.Mul(speed)
	cam.position = cam.position.Add(v)
}

// MoveDown ...
func (cam *Camera) MoveDown(speed float32) {
	v := zaxis.Mul(speed)
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
	cam.Rotate(rad, cam.GetLeft())
}

// Turn ...
func (cam *Camera) Turn(rad float32) {
	cam.Rotate(rad, yaxis)
}

// Rotate ...
func (cam *Camera) Rotate(rad float32, axis Vec3) {
	cam.RotateQ(mgl32.QuatRotate(rad, axis))
}

// RotateQ ...
func (cam *Camera) RotateQ(rotation Quat) {
	cam.orientation = cam.orientation.Mul(rotation)
}

// Projection ...
func (cam *Camera) Projection(mat Mat4) {
	cam.ProjectionMatrix = mat
}

func translate(out *Mat4, v Vec3) {
	x, y, z := v[0], v[1], v[2]

	out[12] = out[0]*x + out[4]*y + out[8]*z + out[12]
	out[13] = out[1]*x + out[5]*y + out[9]*z + out[13]
	out[14] = out[2]*x + out[6]*y + out[10]*z + out[14]
	out[15] = out[3]*x + out[7]*y + out[11]*z + out[15]

}

// Translate ...
func Translate(out *Mat4, v Vec3) {
	x, y, z := v[0], v[1], v[2]

	out[12] = out[0]*x + out[4]*y + out[8]*z + out[12]
	out[13] = out[1]*x + out[5]*y + out[9]*z + out[13]
	out[14] = out[2]*x + out[6]*y + out[10]*z + out[14]
	out[15] = out[3]*x + out[7]*y + out[11]*z + out[15]
}

func setAxisAngle(out *Quat, axis Vec3, rad float32) {
	rad = rad * 0.5
	s := float32(math.Sin(float64(rad)))
	out.V[0] = s * axis[0]
	out.V[1] = s * axis[1]
	out.V[2] = s * axis[2]
	out.W = float32(math.Cos(float64(rad)))
}
