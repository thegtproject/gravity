package opengl

import (
	"unsafe"

	"github.com/go-gl/gl/v4.3-core/gl"
)

// VAO ...
type VAO struct {
	ID ObjectID
}

// NewVertexArrayObject ...
func NewVertexArrayObject() *VAO {
	return &VAO{
		ID: CreateNewVertexArray(),
	}
}

// CreateNewVertexArray ...
func CreateNewVertexArray() ObjectID {
	var id ObjectID
	gl.CreateVertexArrays(1, &id)
	return id
}

// Bind ...
func (vao *VAO) Bind() {
	gl.BindVertexArray(vao.ID)
}

// VBO ...
type VBO struct {
	ID     ObjectID
	Target uint32
	Usage  uint32
	Size   int
	data   []float32
}

// NewVertexBufferObject ...
func NewVertexBufferObject(target uint32, usage uint32) *VBO {
	return &VBO{
		ID:     CreateNewVertexBuffer(),
		Target: target,
		Usage:  usage,
	}
}

// CreateNewVertexBuffer ...
func CreateNewVertexBuffer() ObjectID {
	var id ObjectID
	gl.CreateBuffers(1, &id)
	return id
}

// UpdateAllData ...
func (vbo *VBO) UpdateAllData(data []float32) {
	d := make([]float32, len(data))
	copy(d, data)
	vbo.Size = len(d) * 4
	vbo.data = d
	gl.BindBuffer(vbo.Target, vbo.ID)
	gl.BufferData(
		vbo.Target, vbo.Size,
		unsafe.Pointer(&vbo.data[0]),
		vbo.Usage,
	)
}

// Bind ...
func (vbo *VBO) Bind() {
	gl.BindBuffer(vbo.Target, vbo.ID)
}
