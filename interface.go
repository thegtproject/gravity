package gravity

import "github.com/thegtproject/gravity/pkg/math/mgl32"

// Transformable ...
type Transformable interface {
	Translatable
	Rotatable
	Scaleable

	GetTransformMatrix() mgl32.Mat4
	UpdateTransform()
}

// Translatable ...
type Translatable interface {
	Translate(mgl32.Vec3)
	TranslateX(float32)
	TranslateY(float32)
	TranslateZ(float32)
}

// Rotatable ...
type Rotatable interface {
	RotateX(float32)
	RotateY(float32)
	RotateZ(float32)
	Rotate(mgl32.Quat)
}

// Scaleable ...
type Scaleable interface {
	Scale(mgl32.Vec3)
	Scalef(float32)
	ScaleX(float32)
	ScaleY(float32)
	ScaleZ(float32)
}
