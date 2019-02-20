package gravitygl

import "fmt"

// Attribute ...
type Attribute struct {
	buffer *Buffer
	usage  uint32
	data   *datainfo
	size   int32
}

type datainfo struct {
	data     interface{}
	dataType DataType
}

// NewAttribute ...
func NewAttribute(usage uint32, componentsPerElement int32, data interface{}) *Attribute {
	return &Attribute{
		usage: usage,
		data:  validateDataInterface(data),
		size:  componentsPerElement,
	}
}

// Dump ...
func (attr *Attribute) Dump() {
	fmt.Println("buffer:", attr.buffer)
	fmt.Println("data:", attr.data)
}

func validateDataInterface(data interface{}) *datainfo {
	switch data.(type) {
	case []float32:
		return newData(data, GLFloat)
	case []uint32:
		return newData(data, GLUnsignedInt32)
	default:
		panic("invalid data type for attribute")
	}
}

func newData(d interface{}, dataType DataType) *datainfo {
	return &datainfo{
		data:     d,
		dataType: dataType,
	}
}
