package gravitygl

// VertexArray ...
type VertexArray struct {
	Triangles  *[]uint16
	Attributes []*Attribute
	id         uint32
	buffer     *Buffer
	length     int
}

// NewVertexArray ...
func NewVertexArray() *VertexArray {
	return &VertexArray{
		id: CreateVertexArray(),
	}
}

// Init ...
func (vao *VertexArray) Init() {
	vao.Bind()
	vao.initTrianglesBuffer()

	for i, attrib := range vao.Attributes {
		vao.initAttribBuffer(i)
		VertexAttribPointer(uint32(i), attrib.size, uint32(attrib.data.dataType), false, 0, 0)
		EnableVertexAttribArray(uint32(i))
	}
}

// AddAttributes ...
func (vao *VertexArray) AddAttributes(attr ...*Attribute) {
	vao.Attributes = append(vao.Attributes, attr...)
}

// Bind ...
func (vao *VertexArray) Bind() {
	BindVertexArray(vao.id)
}

// Length ...
func (vao *VertexArray) Length() int {
	return vao.length
}

func (vao *VertexArray) initAttribBuffer(index int) {
	a := vao.Attributes[index]
	if a.buffer != nil {
		panic("initAttribBuffer(): buffer already exists")
	}
	a.buffer = NewBuffer(ARRAY_BUFFER, a.usage, a.data.data)
}

func (vao *VertexArray) initTrianglesBuffer() {
	if vao.buffer != nil {
		panic("initTrianglesBuffer(): buffer already exists")
	}
	vao.buffer = NewBuffer(ELEMENT_ARRAY_BUFFER, STATIC_DRAW, *vao.Triangles)
	vao.length = len(*vao.Triangles)
}
