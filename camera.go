package gravity

import (
	"github.com/thegtproject/gravity/pkg/core/components"
	"github.com/thegtproject/gravity/pkg/math/mgl32"
)

// Camera ...
type Camera struct {
	*Transformer

	ProjectionMatrix mgl32.Mat4
	ViewMatrix       mgl32.Mat4
	pitch, dpitch    float32
	yaw, dyaw        float32

	forward mgl32.Vec3
	up      mgl32.Vec3
	right   mgl32.Vec3
	worldUp mgl32.Vec3
}

func (cam *Camera) GetYawPitch() (float32, float32) {
	return cam.yaw, cam.pitch
}

// NewCamera ...
func NewCamera(options ...CameraOption) *Camera {
	cam := &Camera{
		ProjectionMatrix: mgl32.Perspective(D2R(55), 1920/1080, 0.1, 10000),
		Transformer:      components.NewTransformer(),
		up:               mgl32.Vec3{0, 1, 0},
		worldUp:          mgl32.Vec3{0, 0, 1},
	}

	// defaults
	cam.SetPosition(0, 0, 0)
	// ---
	for _, opt := range options {
		opt(cam)
	}
	return cam
}

// ChangePerspective ...
func (cam *Camera) ChangePerspective(fovy, aspect, near, far float32) {
	cam.ProjectionMatrix = mgl32.Perspective(mgl32.DegToRad(fovy), aspect, near, far)
}

// Update ...
func (cam *Camera) Update() {
	cam.update()
}

// GetViewMatrix ...
func (cam *Camera) GetViewMatrix() mgl32.Mat4 {
	return cam.ViewMatrix
}

// Push ...
func (cam *Camera) Push(pitch, yaw float32) {
	cam.dpitch += pitch
	cam.dyaw += yaw
}

// MoveUp ...
func (cam *Camera) MoveUp(speed float32) {
	v := cam.worldUp.Mul(speed)
	cam.Translate(v)
}

// MoveDown ...
func (cam *Camera) MoveDown(speed float32) {
	v := cam.worldUp.Mul(-speed)
	cam.Translate(v)
}

// MoveForward ...
func (cam *Camera) MoveForward(speed float32) {
	v := cam.forward.Mul(speed)
	cam.Translate(v)
}

// MoveBackward ...
func (cam *Camera) MoveBackward(speed float32) {
	v := cam.forward.Mul(-speed)
	cam.Translate(v)
}

// MoveLeft ...
func (cam *Camera) MoveLeft(speed float32) {
	v := cam.right.Mul(-speed)
	cam.Translate(v)
}

// MoveRight ...
func (cam *Camera) MoveRight(speed float32) {
	v := cam.right.Mul(speed)
	cam.Translate(v)
}

//
//
//
//
// HEY GENIUS,SWAP IT TO A RADIANS SYSTEM, NOT DEGREES
//
//

func (cam *Camera) processYawPitchDeltas() {
	cam.pitch += cam.dpitch
	if cam.pitch > 180.0 {
		cam.pitch = 180.0
	} else if cam.pitch < 0.0 {
		cam.pitch = 0.0
	}
	cam.yaw = Mod(cam.yaw+cam.dyaw, 360)
	cam.dpitch, cam.dyaw = 0, 0
}

func (cam *Camera) processYawPitchDeltasNoLock() {
	cam.pitch = Mod(cam.pitch+cam.dpitch, 360)
	cam.yaw = Mod(cam.yaw+cam.dyaw, 360)
	cam.dpitch, cam.dyaw = 0, 0
}

var (
	localAxisUp = mgl32.Vec3{0, 0, -1}
	tmpRot      = mgl32.Quat{}
)

func (cam *Camera) calculateLocalAxis() {
	cam.forward = tmpRot.Rotate(localAxisUp)
	cam.up = tmpRot.Rotate(cam.worldUp)
	cam.right = cam.forward.Cross(cam.worldUp)
}

func (cam *Camera) update() {
	cam.processYawPitchDeltas()
	//cam.processYawPitchDeltasNoLock()
	tmpRot = Deg2Quat(cam.yaw, cam.pitch)
	cam.calculateLocalAxis()
	cam.SetRotation(tmpRot)

	cam.UpdateTransform()
	cam.ViewMatrix = cam.GetTransformMatrix().Inv()
}

// CameraOption ...
type CameraOption func(c *Camera)

// Position ...
func Position(x, y, z float32) CameraOption {
	return func(c *Camera) {
		c.SetPosition(x, y, z)
	}
}

// Rotate  ...
func Rotate(yaw, pitch float32) CameraOption {
	return func(c *Camera) {
		pitch = Mod(pitch, 180)
		if pitch > 180.0 {
			pitch = 180.0
		} else if pitch < 0.0 {
			pitch = 0.0
		}
		c.dyaw = Mod(yaw, 360)
		c.dpitch = pitch
	}
}
