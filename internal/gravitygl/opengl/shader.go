package opengl

import (
	"github.com/go-gl/gl/v4.3-core/gl"
)

// VertexShader ...
type VertexShader struct {
	loc ShaderLoc
}

// NewVertexShader ...
func NewVertexShader(source string) *VertexShader {
	shader := &VertexShader{
		loc: gl.CreateShader(VERTEX_SHADER),
	}
	src, free := gl.Strs(source)
	defer free()
	length := int32(len(source))
	gl.ShaderSource(shader.loc, 1, src, &length)
	gl.CompileShader(shader.loc)
	var success int32
	gl.GetShaderiv(shader.loc, COMPILE_STATUS, &success)
	if success == FALSE {
		var logLen int32
		gl.GetShaderiv(shader.loc, INFO_LOG_LENGTH, &logLen)

		infoLog := make([]byte, logLen)
		gl.GetShaderInfoLog(shader.loc, logLen, nil, &infoLog[0])
		panic("error compiling vertex shader: " + string(infoLog))
	}

	return shader
}

// FragmentShader ...
type FragmentShader struct {
	loc ShaderLoc
}

// NewFragmentShader ...
func NewFragmentShader(source string) *FragmentShader {
	shader := &FragmentShader{
		loc: gl.CreateShader(FRAGMENT_SHADER),
	}
	src, free := gl.Strs(source)
	defer free()
	length := int32(len(source))
	gl.ShaderSource(shader.loc, 1, src, &length)
	gl.CompileShader(shader.loc)
	var success int32
	gl.GetShaderiv(shader.loc, gl.COMPILE_STATUS, &success)
	if success == gl.FALSE {
		var logLen int32
		gl.GetShaderiv(shader.loc, gl.INFO_LOG_LENGTH, &logLen)

		infoLog := make([]byte, logLen)
		gl.GetShaderInfoLog(shader.loc, logLen, nil, &infoLog[0])
		panic("error compiling fragment shader: " + string(infoLog))
	}

	return shader
}
