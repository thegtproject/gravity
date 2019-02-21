package materials

import (
	"github.com/thegtproject/gravity"
)

// SkyBox ...
type SkyBox struct {
	*gravity.BaseMaterial
}

var _ gravity.Material = &SkyBox{}

// NewSkyBox ...
func NewSkyBox() *SkyBox {
	return &SkyBox{
		BaseMaterial: gravity.NewBaseMaterial("skybox"),
	}
}

// GetBaseMaterial ...
func (mat *SkyBox) GetBaseMaterial() *gravity.BaseMaterial {
	return mat.BaseMaterial
}

// PreRender ...
func (mat *SkyBox) PreRender() {
	mat.Program.Use()
}

// Render ...
func (mat *SkyBox) Render() {

}
