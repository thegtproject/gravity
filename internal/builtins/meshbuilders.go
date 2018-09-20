package builtins

import (
	mb "github.com/thegtproject/gravity/internal/builders/mesh"
	"github.com/thegtproject/gravity/internal/mesh"
)

// MeshBuilders ...
var MeshBuilders = struct {
	TesselatedPlane TesselatedPlane
	Sphere          Sphere
}{
	TesselatedPlane: mb.TesselatedPlane,
	Sphere:          mb.Sphere,
}

// Mesh builder function signatures
type (
	// TesselatedPlane ...
	TesselatedPlane = func() *mesh.Mesh
	// Sphere ...
	Sphere = func(detail int) *mesh.Mesh
)
