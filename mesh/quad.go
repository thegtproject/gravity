package mesh

// NewQuad ...
func NewQuad() *Mesh {
	return &Mesh{
		Position: quadPositions,
		Indices:  quadTriangles,
		Colors:   quadColorsTemp,
	}
}

var quadColorsTemp = []float32{
	0.5, 0.5, 0.5, 1.0,
	0.5, 0.5, 0.5, 1.0,
	0.5, 0.5, 0.5, 1.0,
	0.5, 0.5, 0.5, 1.0,
}

var quadPositions = []float32{
	-0.5, -0.5, 0.0,
	0.5, -0.5, 0.0,
	-0.5, 0.5, 0.0,
	0.5, 0.5, 0.0,
}

var quadTriangles = []uint16{
	0, 1, 2, 2, 1, 3,
}
