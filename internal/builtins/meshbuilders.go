package builtins

import (
	mb "github.com/thegtproject/gravity/internal/builders/mesh"
	"github.com/thegtproject/gravity/internal/mesh"
)

// MeshBuilders ...
var MeshBuilders = struct {
	TesselatedPlane TesselatedPlane
	Sphere          Sphere
	Quad            Quad
}{
	TesselatedPlane: mb.TesselatedPlane,
	Sphere:          mb.Sphere,
	Quad:            mb.Quad,
}

// Mesh builder function signatures
type (
	// TesselatedPlane ...
	TesselatedPlane = func() *mesh.Mesh
	// Sphere ...
	Sphere = func(detail int) *mesh.Mesh
	// Quad ...
	Quad = func(width, height float32) *mesh.Mesh
)
