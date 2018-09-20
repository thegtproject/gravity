package mesh

import (
	"github.com/thegtproject/gravity/internal/geometry"
)

// Mesh ...
type Mesh struct {
	Faces              []Face
	Vertices           []Vertex
	TextureCoordinates []geometry.Vec2
}

// New ...
func New() *Mesh {
	return &Mesh{}
}

// AddVertex ...
func (mesh *Mesh) AddVertex(vert ...Vertex) {
	mesh.Vertices = append(mesh.Vertices, vert...)
}

// AddFace ...
func (mesh *Mesh) AddFace(face ...Face) {
	mesh.Faces = append(mesh.Faces, face...)
}

// AddTextureCoordinate ...
func (mesh *Mesh) AddTextureCoordinate(coord ...geometry.Vec2) {
	mesh.TextureCoordinates = append(mesh.TextureCoordinates, coord...)
}
