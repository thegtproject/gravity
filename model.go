package gravity

import (
	"github.com/thegtproject/gravity/internal/gravitygl"
	"github.com/thegtproject/gravity/math/mgl32"
	"github.com/thegtproject/gravity/mesh"
)

// Model ...
type Model struct {
	*BaseObject

	Mesh *mesh.Mesh
	Mat  Material

	v, p *mgl32.Mat4
}

var attrib = func(d interface{}, u, size uint32) *gravitygl.Attribute {
	return gravitygl.NewAttribute(u, int32(size), d)
}

// NewModel ...
func NewModel(m *mesh.Mesh, material Material, cam *Camera) *Model {
	model := &Model{
		BaseObject: NewBaseObject(),
		Mesh:       m,
		Mat:        material,
	}

	model.v = &cam.ViewMatrix
	model.p = &cam.ProjectionMatrix

	model.vao = gravitygl.NewVertexArray()
	model.vao.Triangles = &model.Mesh.Indices

	model.vao.AddAttributes(attrib(model.Mesh.Positions, gravitygl.STATIC_DRAW, 3))

	if len(model.Mesh.Colors) > 0 {
		model.vao.AddAttributes(attrib(model.Mesh.Colors, gravitygl.STATIC_DRAW, 4))
	}

	model.vao.Init()

	return model
}

// Renderable ...
func (model *Model) Renderable() bool {
	return true
}

// Prepare ...
func (model *Model) Prepare() {
	model.Mat.PreRender()

	model.Transformer.Compose()
	gravitygl.UniformMatrix4fv(0, model.Transformer.Mat)
	gravitygl.UniformMatrix4fv(1, *model.v)
	gravitygl.UniformMatrix4fv(2, *model.p)
}

// Base ...
func (model *Model) Base() *BaseObject {
	return model.BaseObject
}
