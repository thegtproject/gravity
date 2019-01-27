package gravity

import (
	"fmt"

	"github.com/thegtproject/gravity/components/transformer"
	"github.com/thegtproject/gravity/math/mgl32"
)

// Camera ...
type Camera struct {
	Transformer      *transformer.Transformer
	ProjectionMatrix mgl32.Mat4
	ViewMatrix       mgl32.Mat4
	pitch, dpitch    float32
	yaw, dyaw        float32

	forward mgl32.Vec3
	up      mgl32.Vec3
	right   mgl32.Vec3
	worldUp mgl32.Vec3
}

// NewCamera ...
func NewCamera() *Camera {
	cam := &Camera{
		ProjectionMatrix: mgl32.Perspective(mgl32.DegToRad(45), 800/600, 0.1, 10000),
		Transformer:      transformer.NewTransformer(),
		pitch:            90.0,
		yaw:              0.0,
		up:               mgl32.Vec3{0, 1, 0},
		worldUp:          mgl32.Vec3{0, 0, 1},
	}
	return cam
}

// Update ...
func (cam *Camera) Update() {
	cam.updateVectors()
	cam.updateDirection()
	cam.Transformer.Orientation = mgl32.AnglesToQuat(mgl32.DegToRad(cam.yaw), mgl32.DegToRad(cam.pitch), 0, mgl32.ZXY)
	cam.Transformer.Compose()
	cam.ViewMatrix = cam.Transformer.Mat.Inv()
}

// Debug ...
func (cam *Camera) Debug() {
	fmt.Println("Pitch  Deg:", cam.pitch, "  Rad:", mgl32.DegToRad(cam.pitch))
	fmt.Println("Yaw    Deg:", cam.yaw, "  Rad:", mgl32.DegToRad(cam.yaw))
}

var lastForward mgl32.Vec3

func (cam *Camera) updateVectors() {
	cam.forward = cam.Transformer.Orientation.Rotate(mgl32.Vec3{0, 0, -1})
	cam.up = cam.Transformer.Orientation.Rotate(cam.worldUp)
	cam.right = cam.forward.Cross(cam.worldUp)

}

// LookAt ...
func (cam *Camera) LookAt(target mgl32.Vec3) {
	rot := mgl32.QuatLookAtV(cam.Transformer.Position, target, cam.worldUp)
	// mgl32.QuatBetweenVectors(start, dest)
	cam.Transformer.Orientation = rot.Mul(cam.Transformer.Orientation)

	o := cam.Transformer.Orientation
	/////////
	// // Calculate Euler angles (roll, pitch, and yaw) from the unit quaternion.
	// float roll = atan2(2.0f * (quat.w() * quat.x() + quat.y() * quat.z()),
	// 	1.0f - 2.0f * (quat.x() * quat.x() + quat.y() * quat.y()));
	// float pitch = asin(max(-1.0f, min(1.0f, 2.0f * (quat.w() * quat.y() - quat.z() * quat.x()))));
	// float yaw = atan2(2.0f * (quat.w() * quat.z() + quat.x() * quat.y()),
	// 	1.0f - 2.0f * (quat.y() * quat.y() + quat.z() * quat.z()));
	/////////

	// roll := Atan2(2*o.Y()*o.W+2*o.X()*o.Z(), 1-2*o.Y()*o.Y()-2*o.Z()*o.Z())
	pitch := Atan2(2*o.X()*o.W+2*o.Y()*o.Z(), 1-2*o.X()*o.X()-2*o.Z()*o.Z())

	yaw := Asin(2*o.X()*o.Y() + 2*o.Z()*o.W)

	cam.pitch = pitch
	cam.yaw = yaw

}

func (cam *Camera) updateDirection() {
	cam.pitch += cam.dpitch
	if cam.pitch > 89.0 {
		cam.pitch = 89.0
	} else if cam.pitch < -89.0 {
		cam.pitch = -89.0
	}

	cam.yaw = Mod(cam.yaw+cam.dyaw, 360)
	cam.dpitch, cam.dyaw = 0, 0
}

// Rotate ...
func (cam *Camera) Rotate(pitch, yaw float32) {
	cam.dpitch += pitch
	cam.dyaw += yaw
}

// MoveUp ...
func (cam *Camera) MoveUp(speed float32) {
	v := cam.worldUp.Mul(speed)
	cam.Transformer.TranslateV(v)
}

// MoveDown ...
func (cam *Camera) MoveDown(speed float32) {
	v := cam.worldUp.Mul(-speed)
	cam.Transformer.TranslateV(v)
}

// MoveForward ...
func (cam *Camera) MoveForward(speed float32) {
	v := cam.forward.Mul(speed)
	cam.Transformer.TranslateV(v)
}

// MoveBackward ...
func (cam *Camera) MoveBackward(speed float32) {
	v := cam.forward.Mul(-speed)
	cam.Transformer.TranslateV(v)
}

// MoveLeft ...
func (cam *Camera) MoveLeft(speed float32) {
	v := cam.right.Mul(-speed)
	cam.Transformer.TranslateV(v)
}

// MoveRight ...
func (cam *Camera) MoveRight(speed float32) {
	v := cam.right.Mul(speed)
	cam.Transformer.TranslateV(v)
}
