package mesh

// NewQuad ...
func NewQuad() *Mesh {
	return &Mesh{
		Indices:   quadIndices,
		Positions: quadPositions,
		Colors:    quadColors,
		// Coords:    quadCoords,
	}
}

var quadIndices = []uint16{
	0, 1, 2, 2, 1, 3,
}

var quadPositions = []float32{
	-0.5, -0.5, 0.0,
	0.5, -0.5, 0.0,
	-0.5, 0.5, 0.0,
	0.5, 0.5, 0.0,
}

var quadColors = []float32{
	0.5, 0.5, 0.5, 1.0,
	0.5, 0.5, 0.5, 1.0,
	0.5, 0.5, 0.5, 1.0,
	0.5, 0.5, 0.5, 1.0,
}

var quadCoords = []float32{
	0.0, 0.0,
	1.0, 0.0,
	1.0, 1.0,
	0.0, 1.0,
}
