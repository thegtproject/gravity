package gravity

import "github.com/thegtproject/gravity/internal/gravitygl"

// Scene ...
type Scene struct {
	objRenderable []Object
	objOther      []Object
	cam           *Camera
}

// Render ...
func (s *Scene) Render() {
	for i, obj := range s.objRenderable {
		_ = i

		baseObj := obj.Base()
		baseObj.vao.Bind()

		obj.Prepare()

		gravitygl.DrawElements(baseObj.Primitive, baseObj.vao.Length(), gravitygl.UNSIGNED_SHORT, 0)

	}
}

// SetCamera ...
func (s *Scene) SetCamera(c *Camera) {
	s.cam = c
}

// Import ...
func (s *Scene) Import(obj Object) {
	if obj.Renderable() {
		s.objRenderable = append(s.objRenderable, obj)
	} else {
		s.objOther = append(s.objOther, obj)
	}
}
