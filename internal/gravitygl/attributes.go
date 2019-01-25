package gravitygl

import "fmt"

// Attribute ...
type Attribute struct {
	buffer *Buffer
	usage  uint32
	data   *data
	size   int32
}

type data struct {
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

func validateDataInterface(data interface{}) *data {
	switch data.(type) {
	case []float32:
		return newData(data, GLFloat)
	case []uint16:
		return newData(data, GLUnsignedShort)
	default:
		panic("invalid data type for attribute")
	}
}

func newData(d interface{}, dataType DataType) *data {
	return &data{
		data:     d,
		dataType: dataType,
	}
}
