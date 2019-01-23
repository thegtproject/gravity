package gravitygl

import (
	"fmt"
	"unsafe"

	"github.com/go-gl/gl/v4.3-compatibility/gl"
)

type (
	// GLObject is the common gl type returned for gl objects created
	GLObject = uint32
	// GLBuffer is the gl location of a buffer
	GLBuffer = uint32
	// GLShader is the gl location of a shader
	GLShader = uint32
	// GLProgram is the gl location of a program
	GLProgram = uint32
	// GLVertexArray is the location of a VAO
	GLVertexArray = uint32

	// GLAttribute is the gl location of an attribute
	GLAttribute = uint32
	// GLUniform is the gl location of a uniform
	GLUniform = int32

	// GLUint32 is opengl's uint32
	GLUint32 = uint32
	// GLInt32 is opengl's int32
	GLInt32 = int32

	// EnumType is opengl's constants/enums type
	EnumType = uint32
)

var (
	// Target ...
	Target = "desktop/glfw"

	// Update ...
	Update = func() {}

	// GLNullObj ...
	GLNullObj = EnumType(0)
)

// Init ...
func Init() error {
	err := gl.Init()
	gl.DebugMessageCallback(func(
		source uint32,
		gltype uint32,
		id uint32,
		severity uint32,
		length int32,
		message string,
		userParam unsafe.Pointer) {

		if id != 131185 {
			fmt.Println(source, gltype, id, severity, length, message, userParam)
		}
	}, nil)
	return err
}
