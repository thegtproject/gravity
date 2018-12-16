package meshbuilder

import (
	"github.com/thegtproject/gravity/internal/mesh"
)

// TesselatedPlane ...
func TesselatedPlane() *mesh.Mesh {
	mb := New()

	return mb.Build()
}
