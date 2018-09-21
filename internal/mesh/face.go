package mesh

import (
	"github.com/thegtproject/gravity/internal/guid"
)

// Face ...
type Face struct {
	Vertices []*Vertex
	id       id.GUID
}

// NewFace ...
func NewFace() *Face {
	return &Face{id: id.NextID()}
}

// F ...
func F(vertices ...*Vertex) *Face {
	f := NewFace()
	f.PushVertex(vertices...)
	return f
}

// PushVertex ...
func (f *Face) PushVertex(vertices ...*Vertex) {
	for i := 0; i < len(vertices); i++ {
		f.pushVertex(vertices[i])
	}
}

func (f *Face) pushVertex(v *Vertex) {
	f.Vertices = append(f.Vertices, v)
}
