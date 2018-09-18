package gravity

import (
	"github.com/thegtproject/gravity/geometry"
)

// Mesh ...
type Mesh struct {
	Geometry geometry.Geometry
	Material Material
}

// NewMesh ...
func NewMesh(geom geometry.Geometry, mat Material) *Mesh {
	return &Mesh{
		Geometry: geom,
		Material: mat,
	}
}
