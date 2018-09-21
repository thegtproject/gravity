package meshbuilders

import "github.com/thegtproject/gravity/internal/mesh"

// Quad ...
func Quad(width, height float32) *mesh.Mesh {
	halfw, halfh := width/2, height/2
	m := mesh.New()

	A := mesh.V(-halfw, -halfh, 0.0) // A
	B := mesh.V(halfw, -halfh, 0.0)  // B
	C := mesh.V(-halfw, halfh, 0.0)  // C
	D := mesh.V(halfw, halfh, 0.0)   // D

	F1 := mesh.F(A, B, C)
	F2 := mesh.F(C, B, D)

	m.AddVertex(A, B, C, D)
	m.AddFace(F1, F2)

	return m
}
