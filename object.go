package gravity

import (
	"github.com/thegtproject/gravity/core/components"
	"github.com/thegtproject/gravity/internal/gravitygl"
	"github.com/thegtproject/gravity/math/mgl32"
)

// Object ...
type Object interface {
	Base() *BaseObject
	Renderable() bool
	Prepare()
}

// BaseObject ...
type BaseObject struct {
	components.Transformer

	Primitive PrimitiveType
	vao       *gravitygl.VertexArray
	id        uint32
	tag       string
}

// NewBaseObject ...
func NewBaseObject() *BaseObject {
	return &BaseObject{
		Transformer: components.NewTransformer(mgl32.QuatIdent()),
		Primitive:   Triangles,
		id:          0,
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
