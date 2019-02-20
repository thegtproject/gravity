package gravitygl

// Buffer ...
type Buffer struct {
	DataType
	ID     uint32
	target GLUint32
	usage  GLUint32
	data   interface{}
}

// Data ...
func (b *Buffer) Data() interface{} {
	return b.data
}

// Target ...
func (b *Buffer) Target() GLUint32 {
	return b.target
}

// Usage ...
func (b *Buffer) Usage() GLUint32 {
	return b.usage
}

// GLType ...
func (b *Buffer) GLType() GLUint32 {
	return GLUint32(b.DataType)
}

// NewBuffer ...
func NewBuffer(target GLUint32, usage GLUint32, data interface{}) *Buffer {
	var xtype uint32
	switch data.(type) {
	case []float32:
		xtype = FLOAT
	case []uint32:

		xtype = UNSIGNED_INT
	default:

		panic("unknown data type")
	}

	id := CreateBuffer()
	BindBuffer(target, id)
	BufferData(target, usage, data)

	return &Buffer{
		DataType: DataType(xtype),
		data:     data,
		target:   target,
		usage:    usage,
	}
}
