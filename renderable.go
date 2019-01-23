package gravity

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/thegtproject/gravity/internal/gravitygl"
	"github.com/thegtproject/gravity/mesh"
)

// Renderable ...
type Renderable struct {
	Mesh      *mesh.Mesh
	material  uint32
	Primitive uint32
	length    int
	tag       string

	transform Mat4

	vao            uint32
	indicesBuffer  *gravitygl.Buffer
	positionBuffer *gravitygl.Buffer
	colorBuffer    *gravitygl.Buffer
	coordBuffer    *gravitygl.Buffer

	m, v, p int32
}

// Position ...
func (r *Renderable) Position() Vec3 {
	return Vec3{r.transform[12], r.transform[13], r.transform[14]}
}

// Transform ...
func (r *Renderable) Transform(m Mat4) {
	r.transform = r.transform.Mul4(m)
}

// Scale ...
func (r *Renderable) Scale(x, y, z float32) {
	r.transform = r.transform.Mul4(
		mgl32.Scale3D(x, y, z),
	)
}

// Translate ...
func (r *Renderable) Translate(x, y, z float32) {
	r.transform = r.transform.Mul4(
		mgl32.Translate3D(x, y, z),
	)
}

// RotateZ ...
func (r *Renderable) RotateZ(angle float32) {
	r.Transform(mgl32.HomogRotate3DZ(angle))
}
