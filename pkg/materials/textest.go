package materials

import (
	"github.com/thegtproject/gravity"
)

// TexTest ...
type TexTest struct {
	*gravity.BaseMaterial
}

var _ gravity.Material = &TexTest{}

// NewTexTest ...
func NewTexTest() *TexTest {
	return &TexTest{
		BaseMaterial: gravity.NewBaseMaterial("textest"),
	}
}

// GetBaseMaterial ...
func (mat *TexTest) GetBaseMaterial() *gravity.BaseMaterial {
	return mat.BaseMaterial
}

// PreRender ...
func (mat *TexTest) PreRender() {
	mat.Program.Use()
}

// Render ...
func (mat *TexTest) Render() {

}
