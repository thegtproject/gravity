package meshbuilders

import (
	"github.com/thegtproject/gravity/internal/mesh"
)

// TesselatedPlane ...
func TesselatedPlane() *mesh.Mesh {
	m := mesh.New()

	// C-----D
	// | \   |
	// |   \ |
	// A-----B

	m.AddVertex(
		mesh.V(-1.0, -1.0, 0.0), // A
		mesh.V(1.0, -1.0, 0.0),  // B
		mesh.V(-1.0, 1.0, 0.0),  // C
		mesh.V(1.0, 1.0, 0.0),   // D
	)

	return m
}
