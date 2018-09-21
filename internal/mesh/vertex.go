package mesh

import (
	"github.com/thegtproject/gravity/internal/geometry"
	"github.com/thegtproject/gravity/internal/guid"
)

// Vertex ...
type Vertex struct {
	Position geometry.Vec3
	TexCoord geometry.Vec2
	Faces    []*Face
	id       id.GUID
}

// NewVertex ...
func NewVertex() *Vertex {
	return &Vertex{
		id: id.NextID(),
	}
}

// V returns a new Vertex with position initialized
// to given parameters
func V(x, y, z float32) *Vertex {
	v := NewVertex()
	v.Position = geometry.V3(x, y, z)
	return v
}

// VZ returns a new zero'd Vertex
func VZ() *Vertex {
	return NewVertex()
}

func (v *Vertex) pushFace(f *Face) {
	v.Faces = append(v.Faces, f)
}
