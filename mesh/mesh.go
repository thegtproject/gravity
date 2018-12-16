package mesh

// Mesh ...
type Mesh struct {
	Position  []float32
	Normal    []float32
	TexCoords []float32
	Indices   []uint16
}

// NewMesh ...
func NewMesh() *Mesh {
	return &Mesh{}
}
