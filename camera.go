package gravity

import (
	"github.com/go-gl/mathgl/mgl32"
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
	return &Camera{
		ProjectionMatrix: mgl32.Perspective(mgl32.DegToRad(75), 1, 0.1, 1000),
		ViewMatrix:       mgl32.Ident4(),
		position:         mgl32.Vec3{0, 0, 2},
		orientation:      mgl32.QuatIdent(),
	}
}

// Update ...
func (cam *Camera) Update() {
	m := mgl32.Ident4()
	translate(&m, cam.position)
	cam.ViewMatrix = m
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
