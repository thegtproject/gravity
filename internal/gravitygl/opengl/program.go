package opengl

import (
	"fmt"

	"github.com/go-gl/gl/v4.3-core/gl"
)

// Program ...
type Program struct {
	ID       ProgramLoc
	Compiled bool
	Attr     *Attributes
	Unif     *Uniforms
}

func (prg *Program) String() string {
	return fmt.Sprintf(
		`Program
 ID:      %d
 Compiled: %v
 Attributes: %v
 Uniforms:   %v
 `,
		prg.ID, prg.Compiled, prg.Attr.data, prg.Unif.data,
	)
}

// NewProgram ...
func NewProgram(vertexShaderSrc, fragmentShaderSrc string) *Program {
	prg := &Program{
		ID:   gl.CreateProgram(),
		Attr: NewAttributesMap(),
		Unif: NewUniformsMap(),
	}

	vertexShader := NewVertexShader(vertexShaderSrc)
	fragmentShader := NewFragmentShader(fragmentShaderSrc)

	gl.AttachShader(prg.ID, vertexShader.loc)
	gl.AttachShader(prg.ID, fragmentShader.loc)
	gl.LinkProgram(prg.ID)

	var success int32
	gl.GetProgramiv(prg.ID, LINK_STATUS, &success)
	if success == FALSE {
		var logLen int32
		gl.GetProgramiv(prg.ID, INFO_LOG_LENGTH, &logLen)
		infoLog := make([]byte, logLen)
		gl.GetProgramInfoLog(prg.ID, logLen, nil, &infoLog[0])
		panic("error linking shader program: " + string(infoLog))
	}

	{
		count := prg.GetProgramIV(ACTIVE_ATTRIBUTES)
		var idx uint32
		for idx = 0; idx < uint32(count); idx++ {
			name, size, ty := prg.GetActiveAttrib(idx)
			prg.Attr.Add(name, idx, ty, size)
		}
	}

	{
		count := prg.GetProgramIV(ACTIVE_UNIFORMS)
		var idx uint32
		for idx = 0; idx < uint32(count); idx++ {
			name, size, ty := prg.GetActiveUniform(idx)
			prg.Unif.Add(name, idx, ty, size)
		}
	}

	prg.Compiled = true
	gl.DeleteShader(vertexShader.loc)
	gl.DeleteShader(fragmentShader.loc)

	prg.Use()
	return prg
}

// GetProgramIV ...
func (prg *Program) GetProgramIV(pname uint32) int32 {
	var value int32
	gl.GetProgramiv(prg.ID, pname, &value)
	return value
}

// GetActiveAttrib ...
func (prg *Program) GetActiveAttrib(index uint32) (name string, size int, ty Enum) {
	// ACTIVE_ATTRIBUTE_MAX_LENGTH returns the length of the largest attribute name
	// so we can prepare a buffer large enough to contain all possible attribute names
	nameBufferSize := prg.GetProgramIV(ACTIVE_ATTRIBUTE_MAX_LENGTH)
	nameBuffer := make([]byte, nameBufferSize)

	// _size   will return the size of the attribute, in units of the type returned in type.
	// _length will be the length of the attribute's name
	var _size, _length int32
	var _type uint32
	gl.GetActiveAttrib(prg.ID, index, nameBufferSize, &_length, &_size, &_type, &nameBuffer[0])

	return string(nameBuffer[0:_length]), int(_size), Enum(_type)
}

// GetActiveUniform ...
func (prg *Program) GetActiveUniform(index uint32) (name string, size int, ty Enum) {
	// ACTIVE_UNIFORM_MAX_LENGTH returns the length of the largest uniform name
	// so we can prepare a buffer large enough to contain all possible uniform names
	nameBufferSize := prg.GetProgramIV(ACTIVE_UNIFORM_MAX_LENGTH)
	nameBuffer := make([]byte, nameBufferSize)

	// _size   will return the size of the uniform, in units of the type returned in type.
	// _length will be the length of the uniform's name
	var _size, _length int32
	var _type uint32
	gl.GetActiveUniform(prg.ID, index, nameBufferSize, &_length, &_size, &_type, &nameBuffer[0])

	return string(nameBuffer[0:_length]), int(_size), Enum(_type)
}

// Use ...
func (prg *Program) Use() {
	gl.UseProgram(prg.ID)
}
