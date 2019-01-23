package gravitygl

// MakeProgram ...
func MakeProgram(vertexSrc, fragmentSrc string) (GLProgram, error) {
	program := CreateProgram()

	vs, err := MakeVertexShader(vertexSrc)
	if err != nil {
		return 0, err
	}
	fs, err := MakeFragmentShader(fragmentSrc)
	if err != nil {
		return 0, err
	}
	AttachShader(program, vs)
	AttachShader(program, fs)
	if err := LinkProgram(program); err != nil {
		return 0, err
	}
	return program, nil
}

// MakeVertexShader ...
func MakeVertexShader(source string) (GLShader, error) {
	return MakeShader(VERTEX_SHADER, source)
}

// MakeFragmentShader ...
func MakeFragmentShader(source string) (GLShader, error) {
	return MakeShader(FRAGMENT_SHADER, source)
}

// MakeShader ...
func MakeShader(xtype GLUint32, source string) (GLShader, error) {
	shader := CreateShader(xtype)
	ShaderSource(shader, source)
	if err := CompileShader(shader); err != nil {
		return 0, err
	}
	return shader, nil
}
