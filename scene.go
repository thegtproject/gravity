package gravity

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/thegtproject/gravity/mesh"
	gl "github.com/thegtproject/gravitygl"
)

// Scene ...
type Scene struct {
	Objects []*Renderable
	cam     *Camera
}

// Render ...
func (s *Scene) Render() {
	for _, obj := range s.Objects {
		gl.BindVertexArray(obj.vao)
		gl.UseProgram(obj.material)

		gl.UniformMatrix4fv(obj.m, obj.transform)
		gl.UniformMatrix4fv(obj.v, s.cam.ViewMatrix)
		gl.UniformMatrix4fv(obj.p, s.cam.ProjectionMatrix)

		gl.DrawElements(obj.Primitive, obj.length, gl.UNSIGNED_SHORT, 0)
	}
}

// Camera ...
func (s *Scene) Camera(c *Camera) {
	s.cam = c
}

// Import ...
func (s *Scene) Import(tag string, m *mesh.Mesh, material uint32, mods ...func()) *Renderable {
	obj := &Renderable{tag: tag, Mesh: m, transform: mgl32.Ident4(), Primitive: gl.TRIANGLES}
	obj.material = material
	obj.length = len(obj.Mesh.Indices)
	obj.m = gl.GetUniformLocation(obj.material, "uModelMatrix")
	obj.v = gl.GetUniformLocation(obj.material, "uViewMatrix")
	obj.p = gl.GetUniformLocation(obj.material, "uProjectionMatrix")

	position := gl.GetAttribLocation(obj.material, "aPosition")
	colors := gl.GetAttribLocation(obj.material, "aColor")

	vao := gl.CreateVertexArray()
	gl.BindVertexArray(vao)

	obj.indicesBuffer = gl.NewBuffer(gl.ELEMENT_ARRAY_BUFFER, gl.STATIC_DRAW, obj.Mesh.Indices)
	obj.positionBuffer = gl.NewBuffer(gl.ARRAY_BUFFER, gl.STATIC_DRAW, obj.Mesh.Position)
	gl.VertexAttribPointer(position, 3, gl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(position)
	obj.colorBuffer = gl.NewBuffer(gl.ARRAY_BUFFER, gl.STATIC_DRAW, obj.Mesh.Colors)
	gl.VertexAttribPointer(colors, 4, gl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(colors)

	obj.vao = vao
	gl.BindVertexArray(0)

	s.Objects = append(s.Objects, obj)

	return obj
}

// QueryObject ...
func (s *Scene) QueryObject(tag string) *Renderable {
	for _, obj := range s.Objects {
		if obj.tag == tag {
			return obj
		}
	}
	return nil
}
