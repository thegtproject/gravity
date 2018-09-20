package opengl

// Types ...
type (
	ShaderLoc    = uint32
	ProgramLoc   = uint32
	AttributeLoc = uint32
	UniformLoc   = uint32
	ObjectID     = uint32
)

// DataType ...
type DataType int

// Size ...
func (dt DataType) Size() int {
	switch dt {
	case Float:
		return 4
	case Uint16:
		return 2
	default:
		panic("invalid datatype")
	}
}

// Data types
const (
	Float DataType = iota + 1
	Uint16
)

func resolveDataType(data interface{}) DataType {
	switch data.(type) {
	case []float32:
		return Float
	case []uint16:
		return Uint16
	default:
		panic("invalid datatype")
	}
}
