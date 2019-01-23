package gravitygl

// Data types
const (
	GLFloat         DataType = FLOAT
	GLUnsignedShort DataType = UNSIGNED_SHORT
)

// DataType ...
type DataType uint32

// Size ...
func (dt DataType) Size() int {
	switch dt {
	case GLFloat:
		return 4
	case GLUnsignedShort:
		return 2
	default:
		panic("invalid datatype")
	}
}
