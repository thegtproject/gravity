package gravity

import (
	"github.com/thegtproject/gravity/components/transformer"
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
	Primitive   PrimitiveType
	Transformer *transformer.Transformer
	vao         *gravitygl.VertexArray
	id          uint32
	tag         string
}

// NewBaseObject ...
func NewBaseObject() *BaseObject {
	obj := &BaseObject{
		Primitive: Triangles,
		id:        0,
	}
	obj.Transformer = transformer.NewTransformer()
	return obj
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
