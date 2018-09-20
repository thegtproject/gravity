package opengl

import (
	"unsafe"

	"github.com/go-gl/gl/v4.3-core/gl"
)

// Buffer ...
type Buffer struct {
	ID        ObjectID
	Target    uint32
	Usage     uint32
	Size      int
	firstElem unsafe.Pointer
	dataType  DataType
	data      interface{}
}

// NewBuffer ...
func NewBuffer(target uint32, usage uint32) *Buffer {
	return &Buffer{
		ID:     CreateNewBuffer(),
		Target: target,
		Usage:  usage,
	}
}

// CreateNewBuffer ...
func CreateNewBuffer() ObjectID {
	var id ObjectID
	gl.CreateBuffers(1, &id)
	return id
}

// BufferData ...
func (b *Buffer) BufferData(data interface{}, ty DataType) {
	d := makeBufferSlice(data, ty)
	b.Size = getInterfaceSliceLen(d, ty) * ty.Size()
	b.data = d
	b.firstElem = getInterfaceFirstElementPtr(b.data, ty)
	b.dataType = ty
	gl.BindBuffer(b.Target, b.ID)
	gl.BufferData(
		b.Target, b.Size,
		b.firstElem,
		b.Usage,
	)
}

// Bind ...
func (b *Buffer) Bind() {
	gl.BindBuffer(b.Target, b.ID)
}

func getInterfaceFirstElementPtr(data interface{}, ty DataType) unsafe.Pointer {
	switch ty {
	case Float:
		return unsafe.Pointer(&data.([]float32)[0])
	case Uint16:
		return unsafe.Pointer(&data.([]uint16)[0])
	default:
		panic("unsupported data type")
	}
}

func getInterfaceSliceLen(data interface{}, ty DataType) int {
	switch ty {
	case Float:
		return len(data.([]float32))
	case Uint16:
		return len(data.([]uint16))
	default:
		panic("unsupported data type")
	}
}

func makeBufferSlice(data interface{}, ty DataType) interface{} {
	switch ty {
	case Float:
		d := data.([]float32)
		b := make([]float32, len(d))
		copy(b, d)
		return b
	case Uint16:
		d := data.([]uint16)
		b := make([]uint16, len(d))
		copy(b, d)
		return b
	default:
		panic("unsupported data type")
	}
}
