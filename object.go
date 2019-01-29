package gravity

import (
	"github.com/thegtproject/gravity/core/components"
	"github.com/thegtproject/gravity/internal/gravitygl"
)

// Object ...
type Object interface {
	Base() *BaseObject
	Renderable() bool
	Prepare()
}

// BaseObject ...
type BaseObject struct {
	*Transformer

	Primitive PrimitiveType
	vao       *gravitygl.VertexArray
	id        uint32
	tag       string
}

// NewBaseObject ...
func NewBaseObject() *BaseObject {
	return &BaseObject{
		Primitive:   Triangles,
		id:          0,
		Transformer: components.NewTransformer(),
	}
}

// Prepare ...
func (obj *BaseObject) Prepare() {
	// no op on BaseObject
}

// LinkVAO ...
func (obj *BaseObject) LinkVAO(vao *gravitygl.VertexArray) {
	obj.vao = vao
}

// Tag ...
func (obj *BaseObject) Tag() string {
	return obj.tag
}

// SetTag ...
func (obj *BaseObject) SetTag(tag string) {
	obj.tag = tag
}
