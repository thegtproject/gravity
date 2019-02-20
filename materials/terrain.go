package materials

import (
	"github.com/thegtproject/gravity"
)

// Terrain ...
type Terrain struct {
	*gravity.BaseMaterial
}

var _ gravity.Material = &Terrain{}

// NewTerrain ...
func NewTerrain() *Terrain {
	return &Terrain{
		BaseMaterial: gravity.NewBaseMaterial("terrain"),
	}
}

// GetBaseMaterial ...
func (mat *Terrain) GetBaseMaterial() *gravity.BaseMaterial {
	return mat.BaseMaterial
}

// PreRender ...
func (mat *Terrain) PreRender() {
	mat.Program.Use()
}

// Render ...
func (mat *Terrain) Render() {

}
