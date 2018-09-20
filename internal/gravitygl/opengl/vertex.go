package opengl

import (
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
