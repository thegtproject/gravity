package mesh

import (
	"github.com/thegtproject/gravity/internal/geometry"
)

// Vertex ...
type Vertex struct {
	Position geometry.Vec3
}

// V returns a new Vertex with position initialized
// to given parameters
func V(x, y, z float32) Vertex {
	return Vertex{
		Position: geometry.Vec3{x, y, z},
	}
}
