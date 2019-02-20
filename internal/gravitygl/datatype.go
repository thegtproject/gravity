package gravitygl

// Data types
const (
	GLFloat         DataType = FLOAT
	GLUnsignedInt32 DataType = UNSIGNED_INT
)

// DataType ...
type DataType uint32

// Size ...
func (dt DataType) Size() int {
	switch dt {
	case GLFloat:
		return 4
	case GLUnsignedInt32:
		return 4
	default:
		panic("invalid datatype")
	}
}
