package core

import "github.com/thegtproject/gravity/math/mgl32"

// Positionable ...
type Positionable interface {
	Translate(v mgl32.Vec3)
	TranslateX(d float32)
	TranslateY(d float32)
	TranslateZ(d float32)
}

// Rotatable ...
type Rotatable interface {
	RotateX(angle float32)
	RotateY(angle float32)
	RotateZ(angle float32)
	Rotate(q mgl32.Quat)
}
