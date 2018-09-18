package opengl

import (
	"github.com/go-gl/gl/v4.3-core/gl"
)

// Program ...
type Program struct {
	loc      ProgramLoc
	Compiled bool
	attr     *Attributes
	unif     *Uniforms
}

// NewProgram ...
func NewProgram(vertexShaderSrc, fragmentShaderSrc string) *Program {
	program := &Program{
		loc:  gl.CreateProgram(),
		attr: NewAttributesMap(),
		unif: NewUniformsMap(),
	}

	vertexShader := NewVertexShader(vertexShaderSrc)
	fragmentShader := NewFragmentShader(fragmentShaderSrc)

	gl.AttachShader(program.loc, vertexShader.loc)
	gl.AttachShader(program.loc, fragmentShader.loc)
	gl.LinkProgram(program.loc)

	var success int32
	gl.GetProgramiv(program.loc, LINK_STATUS, &success)
	if success == FALSE {
		var logLen int32
		gl.GetProgramiv(program.loc, INFO_LOG_LENGTH, &logLen)
		infoLog := make([]byte, logLen)
		gl.GetProgramInfoLog(program.loc, logLen, nil, &infoLog[0])
		panic("error linking shader program: " + string(infoLog))
	}

	{
		count := program.GetProgramIV(ACTIVE_ATTRIBUTES)
		var idx uint32
		for idx = 0; idx < uint32(count); idx++ {
			name, size, ty := program.GetActiveAttrib(idx)
			program.attr.Add(name, idx, ty, size)
		}
	}

	{
		count := program.GetProgramIV(ACTIVE_UNIFORMS)
		var idx uint32
		for idx = 0; idx < uint32(count); idx++ {
			name, size, ty := program.GetActiveUniform(idx)
			program.unif.Add(name, idx, ty, size)
		}
	}

	program.Compiled = true
	gl.DeleteShader(vertexShader.loc)
	gl.DeleteShader(fragmentShader.loc)
	return program
}

// GetProgramIV ...
func (p *Program) GetProgramIV(pname uint32) int32 {
	var value int32
	gl.GetProgramiv(p.loc, pname, &value)
	return value
}

// GetActiveAttrib ...
func (p *Program) GetActiveAttrib(index uint32) (name string, size int, ty Enum) {
	// ACTIVE_ATTRIBUTE_MAX_LENGTH returns the length of the largest attribute name
	// so we can prepare a buffer large enough to contain all possible attribute names
	nameBufferSize := p.GetProgramIV(ACTIVE_ATTRIBUTE_MAX_LENGTH)
	nameBuffer := make([]byte, nameBufferSize)

	// _size   will return the size of the attribute, in units of the type returned in type.
	// _length will be the length of the attribute's name
	var _size, _length int32
	var _type uint32
	gl.GetActiveAttrib(p.loc, index, nameBufferSize, &_length, &_size, &_type, &nameBuffer[0])

	return string(nameBuffer[0:_length]), int(_size), Enum(_type)
}

// GetActiveUniform ...
func (p *Program) GetActiveUniform(index uint32) (name string, size int, ty Enum) {
	// ACTIVE_UNIFORM_MAX_LENGTH returns the length of the largest uniform name
	// so we can prepare a buffer large enough to contain all possible uniform names
	nameBufferSize := p.GetProgramIV(ACTIVE_UNIFORM_MAX_LENGTH)
	nameBuffer := make([]byte, nameBufferSize)

	// _size   will return the size of the uniform, in units of the type returned in type.
	// _length will be the length of the uniform's name
	var _size, _length int32
	var _type uint32
	gl.GetActiveUniform(p.loc, index, nameBufferSize, &_length, &_size, &_type, &nameBuffer[0])

	return string(nameBuffer[0:_length]), int(_size), Enum(_type)
}
