package builtins

import (
	"github.com/thegtproject/gravity/geometry"
	gb "github.com/thegtproject/gravity/geometry/builders"
)

// GeometryBuilders ...
var GeometryBuilders = struct {
	TesselatedPlane TesselatedPlane
	Sphere          Sphere
}{
	TesselatedPlane: gb.TesselatedPlane,
	Sphere:          gb.Sphere,
}

// Geometry builder function signatures
type (
	// TesselatedPlane ...
	TesselatedPlane = func() *geometry.Geometry
	// Sphere ...
	Sphere = func(detail int) *geometry.Geometry
)
