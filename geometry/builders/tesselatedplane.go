package geometrybuilders

import (
	"github.com/thegtproject/gravity/geometry"
)

// TesselatedPlane ...
func TesselatedPlane() *geometry.Geometry {
	geom := geometry.New()

	// C-----D
	// | \   |
	// |   \ |
	// A-----B

	geom.AddVertex(
		vert(-1.0, -1.0, 0.0), // A
		vert(1.0, -1.0, 0.0),  // B
		vert(-1.0, 1.0, 0.0),  // C
		vert(1.0, 1.0, 0.0),   // D
	)

	return geom
}
