package materials

import (
	"github.com/thegtproject/gravity"
	"github.com/thegtproject/gravity/internal/gravitygl"
	"github.com/thegtproject/gravity/pkg/math/mgl32"
)

// SingleColor ...
type SingleColor struct {
	*gravity.BaseMaterial

	Ambient  mgl32.Vec3
	Diffuse  mgl32.Vec3
	Specular mgl32.Vec3
	Emissive mgl32.Vec3
}

var _ gravity.Material = &SingleColor{}

// NewSingleColor ...
func NewSingleColor(Ambient mgl32.Vec3, Diffuse mgl32.Vec3, Specular mgl32.Vec3, Emissive mgl32.Vec3) *SingleColor {
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
	mat.Program.Use()
	gravitygl.Uniform3fv(3, mat.Diffuse)
}

// Render ...
func (mat *SingleColor) Render() {

}
