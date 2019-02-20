package gravitygl

import (
	"github.com/go-gl/gl/v4.3-compatibility/gl"
)

// Program ...
type Program struct {
	id         uint32
	Attributes []AttributesInfo
	Uniforms   []UniformInfo
}

// AttributesInfo ...
type AttributesInfo struct {
	Name string
	Loc  uint32
	Size int
	Type Enum
}

// UniformInfo ...
type UniformInfo struct {
	Name string
	Loc  int32
	Size int
	Type Enum
}

// UpdateAttributesUniformsLayout ...
func (prg *Program) UpdateAttributesUniformsLayout() {
	var numAttribs int32
	gl.GetProgramiv(prg.id, ACTIVE_ATTRIBUTES, &numAttribs)
	for i := int32(0); i < numAttribs; i++ {
		n, s, t := GetActiveAttrib(prg.id, uint32(i))
		loc := GetAttribLocation(prg.id, n)
		prg.Attributes = append(prg.Attributes, AttributesInfo{Loc: loc, Name: n, Size: s, Type: t})
	}

	var numUniforms int32
	gl.GetProgramiv(prg.id, ACTIVE_UNIFORMS, &numUniforms)
	for i := int32(0); i < numUniforms; i++ {
		n, s, t := GetActiveUniform(prg.id, i)
		loc := GetUniformLocation(prg.id, n)
		prg.Uniforms = append(prg.Uniforms, UniformInfo{Loc: loc, Name: n, Size: s, Type: t})
	}
}

// NewProgram ...
func NewProgram(vertexSource, fragmentSource string) *Program {
	id, err := MakeProgram(vertexSource, fragmentSource)
	if err != nil {
		panic(err)
	}
	prg := &Program{
		id: id,
	}
	prg.UpdateAttributesUniformsLayout()
	return prg
}

// Use ...
func (prg *Program) Use() {
	UseProgram(prg.id)
}
