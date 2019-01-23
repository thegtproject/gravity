package gravitygl

func resolveDataType(data interface{}) DataType {
	switch data.(type) {
	case []float32:
		return GLFloat
	case []uint16:
		return GLUnsignedShort
	default:
		panic("invalid datatype")
	}
}

func getInterfaceSliceLen(data interface{}, ty DataType) int {
	switch ty {
	case GLFloat:
		return len(data.([]float32))
	case GLUnsignedShort:
		return len(data.([]uint16))
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
	case GLUnsignedShort:
		d := data.([]uint16)
		b := make([]uint16, len(d))
		copy(b, d)
		return b
	default:
		panic("unsupported data type")
	}
}
