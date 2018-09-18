package opengl

import (
	"unsafe"

	"github.com/go-gl/gl/v4.3-core/gl"
)

// Buffer ...
type Buffer struct {
	loc  BufferLoc
	data []float32
}

// NewBuffer ...
func NewBuffer() *Buffer {
	buffer := &Buffer{}
	gl.CreateBuffers(1, &buffer.loc)
	return buffer
}

// UpdateDataFull ...
func (b *Buffer) UpdateDataFull(data []float32) {
	gl.BindBuffer(ARRAY_BUFFER, b.loc)
	gl.BufferData(
		ARRAY_BUFFER, len(data)*4,
		unsafe.Pointer(&data[0]),
		STATIC_DRAW,
	)
}
