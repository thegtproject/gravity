package gravitygl

func resolveDataType(data interface{}) DataType {
	switch data.(type) {
	case []float32:
		return GLFloat
	case []uint32:
		return GLUnsignedInt32
	default:
		panic("invalid datatype")
	}
}

func getInterfaceSliceLen(data interface{}, ty DataType) int {
	switch ty {
	case GLFloat:
		return len(data.([]float32))
	case GLUnsignedInt32:
		return len(data.([]uint32))
	default:
		panic("unsupported data type")
	}
}

func makeBufferSlice(data interface{}, ty DataType) interface{} {
	switch ty {
	case GLFloat:
		d := data.([]float32)
		b := make([]float32, len(d))
		copy(b, d)
		return b
	case GLUnsignedInt32:
		d := data.([]uint32)
		b := make([]uint32, len(d))
		copy(b, d)
		return b
	default:
		panic("unsupported data type")
	}
}
