package gravity

import (
	"github.com/thegtproject/gravity/internal/gravitygl"
	"github.com/thegtproject/gravity/pkg/math/mgl32"
	"github.com/thegtproject/gravity/pkg/mesh"
)

// Model ...
type Model struct {
	*BaseObject

	uniformSubmitList []UniformSubmission

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

	// vao and attributes
	model.vao = gravitygl.NewVertexArray()
	model.vao.Triangles = &model.Mesh.Indices
	model.vao.AddAttributes(attrib(model.Mesh.Positions, gravitygl.STATIC_DRAW, 3))
	if len(model.Mesh.Colors) > 0 {
		model.vao.AddAttributes(attrib(model.Mesh.Colors, gravitygl.STATIC_DRAW, 4))
	}
	if len(model.Mesh.Coords) > 0 {
		model.vao.AddAttributes(attrib(model.Mesh.Coords, gravitygl.STATIC_DRAW, 2))
	}

	model.vao.Init()

	// default uniforms

	model.AddUniform("uProjectionMatrix", model.p)
	model.AddUniform("uViewMatrix", model.v)
	model.AddUniform("uModelMatrix", model.GetTransformMatrix())

	return model
}

// AddUniform ...
func (model *Model) AddUniform(name string, data interface{}) {
	//	fmt.Println("Adding:", name, "\n", data, "\n", "---------------------------")
	for _, u := range model.Mat.GetBaseMaterial().Program.Uniforms {
		if u.Name == name && u.Loc > -1 {
			model.uniformSubmitList = append(model.uniformSubmitList,
				UniformSubmission{
					Type: u.Type,
					Loc:  u.Loc,
					Data: data,
				},
			)
			//	fmt.Println("Added:", u.Name, "Loc:", u.Loc, "data:\n", data)
			return
		}
	}
	//	fmt.Printf("model.adduniform: error: \"%s\" does not exist\n", name)
}

// Renderable ...
func (model *Model) Renderable() bool {
	return true
}

// Prepare ...
func (model *Model) Prepare() {
	model.Mat.PreRender()
	model.UpdateTransform()

	model.Mat.GetBaseMaterial().SubmitUniforms(model.uniformSubmitList)
}

// Base ...
func (model *Model) Base() *BaseObject {
	return model.BaseObject
}
