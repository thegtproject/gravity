//+build !web

package gravitygl

import (
	"errors"

	"github.com/go-gl/gl/v4.3-compatibility/gl"
	"github.com/thegtproject/gravity/math/glmath"
)

// ClearColor ...
func ClearColor(c glmath.Vec4) {
	gl.ClearColor(c[0], c[1], c[2], c[3])
}

// Enable ...
func Enable(option GLUint32) {
	gl.Enable(option)
}

// Disable ...
func Disable(option GLUint32) {
	gl.Disable(option)
}

// DepthFunc ...
func DepthFunc(xfunc GLUint32) {
	gl.DepthFunc(xfunc)
}

// BlendFunc ...
func BlendFunc(sfactor, dfactor GLUint32) {
	gl.BlendFunc(sfactor, dfactor)
}

// DepthMask ...
func DepthMask(flag bool) {
	gl.DepthMask(flag)
}

// ViewPort ...
func ViewPort(x, y, width, height GLInt32) {
	gl.Viewport(x, y, width, height)
}

// Clear ...
func Clear(mask GLUint32) {
	gl.Clear(mask)
}

// GetGLVersion ...
func GetGLVersion() string {
	return GetString(gl.VERSION)
}

// GetString ...
func GetString(name GLUint32) string {
	return gl.GoStr(gl.GetString(name))
}

// CreateVertexArray ...
func CreateVertexArray() GLVertexArray {
	var vao GLVertexArray
	gl.CreateVertexArrays(1, &vao)
	return vao
}

// BindVertexArray ...
func BindVertexArray(vao GLVertexArray) {
	gl.BindVertexArray(vao)
}

// CreateProgram ...
func CreateProgram() GLProgram {
	return gl.CreateProgram()
}

// AttachShader ...
func AttachShader(program GLProgram, shader GLShader) {
	gl.AttachShader(program, shader)
}

// LinkProgram ...
func LinkProgram(program GLProgram) error {
	gl.LinkProgram(program)
	success := GetProgramParameterb(program, LINK_STATUS)
	if !success {
		return errors.New(
			GetProgramInfoLog(program),
		)
	}
	return nil
}

// GetProgramParameterb ...
func GetProgramParameterb(program GLProgram, pname GLUint32) bool {
	var val int32
	gl.GetProgramiv(program, pname, &val)
	return val == TRUE
}

// GetProgramParameteri ...
func GetProgramParameteri(program GLProgram, pname GLUint32) GLInt32 {
	var val int32
	gl.GetProgramiv(program, pname, &val)
	return val
}

// GetProgramInfoLog ...
func GetProgramInfoLog(program GLProgram) string {
	var logLen int32
	gl.GetProgramiv(program, INFO_LOG_LENGTH, &logLen)
	logStr := make([]byte, logLen)
	gl.GetProgramInfoLog(program, logLen, nil, &logStr[0])
	return string(logStr)
}

// DeleteShader ...
func DeleteShader(shader GLShader) {
	gl.DeleteShader(shader)
}

// UseProgram ...
func UseProgram(program GLProgram) {
	gl.UseProgram(program)
}

// GetActiveAttrib ...
func GetActiveAttrib(program GLProgram, attr GLAttribute) (name string, size int, ty Enum) {
	// ACTIVE_ATTRIBUTE_MAX_LENGTH returns the length of the largest attribute name
	// so we can prepare a buffer large enough to contain all possible attribute names
	var nameBufferSize int32
	gl.GetProgramiv(program, ACTIVE_ATTRIBUTE_MAX_LENGTH, &nameBufferSize)
	nameBuffer := make([]byte, nameBufferSize)

	// _size   will return the size of the attribute, in units of the type returned in type.
	// _length will be the length of the attribute's name
	var _size, _length int32
	var _type uint32
	gl.GetActiveAttrib(program, attr, nameBufferSize, &_length, &_size, &_type, &nameBuffer[0])

	return string(nameBuffer[0:_length]), int(_size), Enum(_type)
}

// GetActiveUniform ...
func GetActiveUniform(program GLProgram, unif GLUniform) (name string, size int, ty Enum) {
	if unif < 0 {
		return
	}
	// ACTIVE_UNIFORM_MAX_LENGTH returns the length of the largest uniform name
	// so we can prepare a buffer large enough to contain all possible uniform names
	var nameBufferSize int32
	gl.GetProgramiv(program, ACTIVE_UNIFORM_MAX_LENGTH, &nameBufferSize)
	nameBuffer := make([]byte, nameBufferSize)

	// _size   will return the size of the uniform, in units of the type returned in type.
	// _length will be the length of the uniform's name
	var _size, _length int32
	var _type uint32
	gl.GetActiveUniform(program, uint32(unif), nameBufferSize, &_length, &_size, &_type, &nameBuffer[0])

	return string(nameBuffer[0:_length]), int(_size), Enum(_type)
}

// CreateShader ...
func CreateShader(xtype GLUint32) GLShader {
	return gl.CreateShader(xtype)
}

// ShaderSource ...
func ShaderSource(shader GLShader, source string) {
	src, free := gl.Strs(source)
	defer free()
	length := int32(len(source))
	gl.ShaderSource(shader, 1, src, &length)
}

// CompileShader ...
func CompileShader(shader GLShader) error {
	gl.CompileShader(shader)
	if !GetShaderParameterb(shader, COMPILE_STATUS) {
		return errors.New(GetShaderInfoLog(shader))
	}
	return nil
}

// GetShaderInfoLog ...
func GetShaderInfoLog(shader GLShader) string {
	logLen := GetShaderParameteri(shader, INFO_LOG_LENGTH)
	infoLog := make([]byte, logLen)
	gl.GetShaderInfoLog(shader, logLen, nil, &infoLog[0])
	return string(infoLog)
}

// GetShaderParameteri ...
func GetShaderParameteri(shader GLShader, pname GLUint32) GLInt32 {
	var val int32
	gl.GetShaderiv(shader, pname, &val)
	return val
}

// GetShaderParameterb ...
func GetShaderParameterb(shader GLShader, pname GLUint32) bool {
	var val int32
	gl.GetShaderiv(shader, pname, &val)
	return val == TRUE
}

// CreateBuffer ...
func CreateBuffer() GLBuffer {
	var buffer GLBuffer
	gl.CreateBuffers(1, &buffer)
	return buffer
}

// BufferData ...
func BufferData(target GLUint32, usage GLUint32, data interface{}) {
	ty := resolveDataType(data)
	d := makeBufferSlice(data, ty)
	size := getInterfaceSliceLen(d, ty) * ty.Size()
	gl.BufferData(target, size, gl.Ptr(d), usage)
}

// BindBuffer ...
func BindBuffer(target GLUint32, buffer GLBuffer) {
	gl.BindBuffer(target, buffer)
}

// VertexAttribPointer ...
//   opengl: VertexAttribPointer(index uint32, size int32, xtype uint32, normalized bool, stride int32, pointer unsafe.Pointer)
//    webgl: VertexAttribPointer(index    int, size   int, typ      int, normal     bool, stride   int,  offset int)
func VertexAttribPointer(index EnumType, size GLInt32, xtype EnumType, normalized bool, stride GLInt32, offset int) {
	gl.VertexAttribPointer(index, size, xtype, normalized, stride, gl.PtrOffset(offset))
}

// EnableVertexAttribArray ...
func EnableVertexAttribArray(index EnumType) {
	gl.EnableVertexAttribArray(index)
}

// DrawElements ...
func DrawElements(mode EnumType, count int, xtype EnumType, offset int) {
	gl.DrawElements(mode, int32(count), xtype, gl.PtrOffset(offset))
}

// LineWidth ...
func LineWidth(width float32) {
	gl.LineWidth(width)
}

// Hint ...
func Hint(target, mode uint32) {
	gl.Hint(target, mode)
}

// Scissor ...
func Scissor(x, y, w, h int32) {
	gl.Scissor(x, y, w, h)
}

// BlendEquation ...
func BlendEquation(mode uint32) {
	gl.BlendEquation(mode)
}

// GetAttribLocation ...
func GetAttribLocation(program GLProgram, name string) EnumType {
	b := []byte(name)
	return EnumType(gl.GetAttribLocation(program, &b[0]))
}

// Uniform1i ...
func Uniform1i(loc GLUniform, val int32) {
	gl.Uniform1i(int32(loc), val)
}

// Uniform1f ...
func Uniform1f(loc GLUniform, val float32) {
	gl.Uniform1f(int32(loc), val)
}

// Uniform3fv ...
func Uniform3fv(loc GLUniform, val [3]float32) {
	gl.Uniform3fv(loc, 1, &val[0])
}

// Uniform4f ...
func Uniform4f(loc GLUniform, val [4]float32) {
	gl.Uniform4fv(loc, 1, &val[0])
}

// GetUniformLocation ...
func GetUniformLocation(program GLProgram, name string) GLUniform {
	b := []byte(name)
	return gl.GetUniformLocation(program, &b[0])
}

// UniformMatrix4fv ...
func UniformMatrix4fv(loc GLUniform, val [16]float32) {
	gl.UniformMatrix4fv(loc, 1, false, &val[0])
}

// ClearDepth ...
func ClearDepth(depth float64) {
	gl.ClearDepth(depth)
}

// GetError ...
func GetError() Enum {
	return Enum(gl.GetError())
}
