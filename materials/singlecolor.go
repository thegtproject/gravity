package materials

import (
	"github.com/thegtproject/gravity"
)

// SingleColor ...
type SingleColor struct {
	*gravity.BaseMaterial

	Ambient  gravity.Vec3
	Diffuse  gravity.Vec3
	Specular gravity.Vec3
	Emissive gravity.Vec3
}

// NewSingleColor ...
func NewSingleColor(Ambient, Diffuse, Specular, Emissive gravity.Vec3) *SingleColor {
	return &SingleColor{
		BaseMaterial: gravity.NewBaseMaterial("singlecolor"),
		Ambient:      Ambient,
		Diffuse:      Diffuse,
		Specular:     Specular,
		Emissive:     Emissive,
	}
}

// GetBaseMaterial ...
func (mat *SingleColor) GetBaseMaterial() *gravity.BaseMaterial {
	return mat.BaseMaterial
}

// PreRender ...
func (mat *SingleColor) PreRender() {

}

// Render ...
func (mat *SingleColor) Render() {

}
