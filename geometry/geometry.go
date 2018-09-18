package geometry

// Geometry ...
type Geometry struct {
	Faces              []Face
	Vertices           []Vertex
	TextureCoordinates []Vec2
}

// New ...
func New() *Geometry {
	return &Geometry{}
}

// AddVertex ...
func (geom *Geometry) AddVertex(vert ...Vertex) {
	geom.Vertices = append(geom.Vertices, vert...)
}

// AddFace ...
func (geom *Geometry) AddFace(face ...Face) {
	geom.Faces = append(geom.Faces, face...)
}

// AddTextureCoordinate ...
func (geom *Geometry) AddTextureCoordinate(coord ...Vec2) {
	geom.TextureCoordinates = append(geom.TextureCoordinates, coord...)
}
