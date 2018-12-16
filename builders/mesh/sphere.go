package meshbuilder

import (
	"github.com/thegtproject/gravity/mesh"
)

// Sphere ...
func Sphere(detail int) *mesh.Mesh {
	mb := New()

	return mb.Build()
}
