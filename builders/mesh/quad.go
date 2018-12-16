package meshbuilder

import (
	"fmt"

	"github.com/thegtproject/gravity/mesh"
)

// Quad ...
func Quad(width, height float32) *mesh.Mesh {
	halfw, halfh := width/2, height/2
	mb := New()

	A := V(-halfw, -halfh, 0.0) // A
	B := V(halfw, -halfh, 0.0)  // B
	C := V(-halfw, halfh, 0.0)  // C
	D := V(halfw, halfh, 0.0)   // D

	F1 := F(A, B, C)
	F2 := F(C, B, D)

	println(mb.indexer.Add(A.Position))
	println(mb.indexer.Add(B.Position))
	println(mb.indexer.Add(C.Position))
	println(mb.indexer.Add(D.Position))
	println(mb.indexer.Add(B.Position))
	println(mb.indexer.Add(B.Position))
	println(mb.indexer.Add(D.Position))
	println(mb.indexer.Add(D.Position))

	mb.AddVertex(A, B, C, D)
	mb.AddFace(F1, F2)

	fmt.Println("Indexer:\n", mb.indexer.table)

	return mb.Build()
}
