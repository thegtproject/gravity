package geometry

// Vertex ...
type Vertex struct {
	Position Vec3
}

// V returns a new Vertex with position initialized
// to given paramters
func V(x, y, z float32) Vertex {
	return Vertex{
		Position: Vec3{x, y, z},
	}
}
