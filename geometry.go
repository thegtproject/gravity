package gravity

import (
	"github.com/thegtproject/gravity/geometry"
	"github.com/thegtproject/gravity/geometry/builders"
)

// GeometryBuilders ...
var GeometryBuilders = GeometryBuilderList{
	TesselatedPlane: builders.TesselatedPlane,
}

// GeometryBuilderList ...
type GeometryBuilderList struct {
	TesselatedPlane TesselatedPlane
}

// Geometry builder function signatures
type (
	// TesselatedPlane ...
	TesselatedPlane = func() *geometry.Geometry
)
