package meshbuilder

import (
	"github.com/thegtproject/gravity/mesh"
)

// TesselatedPlane ...
func TesselatedPlane() *mesh.Mesh {
	mb := New()

	return mb.Build()
}
