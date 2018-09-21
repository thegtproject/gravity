package mesh

import (
	"fmt"

	"github.com/thegtproject/gravity/internal/guid"
)

// Mesh ...
type Mesh struct {
	Faces    []*Face
	Vertices []*Vertex
	id       id.GUID
}

// New ...
func New() *Mesh {
	return &Mesh{id: id.NextID()}
}

// AddVertex ...
func (mesh *Mesh) AddVertex(vert ...*Vertex) {
	mesh.Vertices = append(mesh.Vertices, vert...)
}

// AddFace ...
func (mesh *Mesh) AddFace(face ...*Face) {
	mesh.Faces = append(mesh.Faces, face...)
}

// Print ...
func (mesh *Mesh) Print() {

	fmt.Printf("----------------\n")
	fmt.Printf("Mesh %d\n", mesh.id)
	fmt.Printf("----------------\n")
	for i := 0; i < len(mesh.Faces); i++ {
		face := mesh.Faces[i]
		fmt.Printf("  | Face %d\n", face.id)
		fmt.Printf("  |.......................................\n")
		fmt.Printf("  | vert id   | vert position            |\n")
		fmt.Printf("  |```````````|``````````````````````````|\n")
		for j := 0; j < len(face.Vertices); j++ {
			fmt.Printf("  | %d         |    %v\n",
				face.Vertices[j].id, face.Vertices[j].Position,
			)
		}
		fmt.Printf("  |...........|..........................|\n")
	}
}
