package materials

import (
	"github.com/thegtproject/gravity"
)

// None ...
type None struct {
	*gravity.BaseMaterial
}

var _ gravity.Material = &None{}

// NewNone ...
func NewNone() *None {
	return &None{
		BaseMaterial: gravity.NewBaseMaterial("none"),
	}
}

// GetBaseMaterial ...
func (mat *None) GetBaseMaterial() *gravity.BaseMaterial {
	return mat.BaseMaterial
}

// PreRender ...
func (mat *None) PreRender() {
	mat.Program.Use()
}

// Render ...
func (mat *None) Render() {

}
