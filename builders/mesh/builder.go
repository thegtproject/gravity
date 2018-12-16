package meshbuilder

import (
	"fmt"

	"github.com/thegtproject/gravity/mesh"

	"github.com/thegtproject/gravity/geometry"
	"github.com/thegtproject/gravity/guid"
)

// MeshBuilder ...
type MeshBuilder struct {
	Faces    []*Face
	Vertices []*Vertex

	indexer Indexer
	id      id.GUID
}

// Build ...
func (mb *MeshBuilder) Build() *mesh.Mesh {
	m := mesh.NewMesh()

	var (
		positions []float32
		// normals   []float32
		texcoords []float32
		indices   []uint16
	)

	for i := 0; i < len(mb.Faces); i++ {
		fverts := mb.Faces[i].Vertices
		for j := 0; j < len(fverts); j++ {
			v := *fverts[j]
			positions = append(positions,
				v.Position[0], v.Position[1], v.Position[2],
			)
			texcoords = append(texcoords,
				v.TexCoord[0], v.TexCoord[1],
			)
		}
		indices = append(indices, uint16(i*3))
	}

	m.Position = make([]float32, len(positions))
	m.TexCoords = make([]float32, len(texcoords))
	m.Indices = make([]uint16, len(indices))
	copy(m.Position, positions)
	copy(m.TexCoords, texcoords)
	copy(m.Indices, indices)

	return m
}

// New ...
func New() *MeshBuilder {
	return &MeshBuilder{
		id: id.NextID(),
		indexer: Indexer{
			table: make(map[geometry.Vec3]int),
		},
	}
}

// AddVertex ...
func (mb *MeshBuilder) AddVertex(vert ...*Vertex) {
	mb.Vertices = append(mb.Vertices, vert...)
}

// AddFace ...
func (mb *MeshBuilder) AddFace(face ...*Face) {
	mb.Faces = append(mb.Faces, face...)
}

// Print ...
func (mb *MeshBuilder) Print() {

	fmt.Printf("----------------\n")
	fmt.Printf("Mesh %d\n", mb.id)
	fmt.Printf("----------------\n")
	for i := 0; i < len(mb.Faces); i++ {
		face := mb.Faces[i]
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
