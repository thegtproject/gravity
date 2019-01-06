package gravity

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/thegtproject/gravity/mesh"
	"github.com/thegtproject/gravitygl"
)

// Renderable ...
type Renderable struct {
	Mesh      *mesh.Mesh
	material  uint32
	Primitive uint32
	length    int
	tag       string

	transform mgl32.Mat4

	vao            uint32
	indicesBuffer  *gravitygl.Buffer
	positionBuffer *gravitygl.Buffer
	colorBuffer    *gravitygl.Buffer

	m, v, p int32
}

// Transform ...
func (r *Renderable) Transform(m mgl32.Mat4) {
	r.transform = r.transform.Mul4(m)
}

// RotateZ ...
func (r *Renderable) RotateZ(angle float32) {
	r.Transform(mgl32.HomogRotate3DZ(angle))
}
