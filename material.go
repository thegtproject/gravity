package gravity

import (
	"fmt"
	"io/ioutil"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/thegtproject/gravity/internal/gravitygl"
)

// Material ...
type Material interface {
	GetBaseMaterial() *BaseMaterial
	PreRender()
	Render()
}

// BaseMaterial ...
type BaseMaterial struct {
	ID        uint32
	Primitive PrimitiveType
	Program   *gravitygl.Program
}

// NewBaseMaterial ...
func NewBaseMaterial(programName string) *BaseMaterial {
	return &BaseMaterial{
		ID:        0,
		Primitive: Triangles,
		Program:   GetMaterialProgram(programName),
	}
}

// PrimitiveType ...
type PrimitiveType = uint32

// Primitive types
const (
	Triangles = PrimitiveType(gravitygl.TRIANGLES)
	Lines     = PrimitiveType(gravitygl.LINES)
	Points    = PrimitiveType(gravitygl.POINTS)
)

var materialPrograms = map[string]*gravitygl.Program{}

// GetMaterialProgram ...
func GetMaterialProgram(name string) *gravitygl.Program {
	return materialPrograms[name]
}

func addMaterialProgram(name, vertexSource, fragmentSource string) {
	materialPrograms[name] = gravitygl.NewProgram(vertexSource, fragmentSource)
}

func loadDefaultMaterialPrograms() {
	defaultMaterialNames := []string{
		"singlecolor",
		"none",
		"textest",
		"terrain",
		"skybox",
	}

	Log.Print("loading default material programs: ")
	for i, n := range defaultMaterialNames {
		if i < len(defaultMaterialNames)-1 {
			Log.Print(n, ", ")
		} else {
			Log.Print(n)
		}
		loadMaterialProgram(n)
	}
	Log.Println("")
}

func loadMaterialProgram(name string) {
	vertFilename := fmt.Sprintf("assets/shaders/%s.vert.glsl", name)
	fragFilename := fmt.Sprintf("assets/shaders/%s.frag.glsl", name)
	v, err := ioutil.ReadFile(vertFilename)
	if err != nil {
		panic(err)
	}
	f, err := ioutil.ReadFile(fragFilename)
	if err != nil {
		panic(err)
	}
	addMaterialProgram(name, string(v), string(f))
}

// UniformSubmission ...
type UniformSubmission struct {
	Type gravitygl.Enum
	Loc  int32
	Data interface{}
}

// SubmitUniforms ...
func (mat *BaseMaterial) SubmitUniforms(ulist []UniformSubmission) {
	for _, u := range ulist {
		switch u.Type {
		case gravitygl.FLOAT_MAT4:
			data := *u.Data.(*mgl32.Mat4)
			gravitygl.UniformMatrix4fv(u.Loc, data)
		case gravitygl.SAMPLER_2D:
			data := u.Data.(*Texture)
			gravitygl.ActiveTexture(data.Unit)
			gravitygl.BindTexture(data.Target, data.TextureID)
			gravitygl.Uniform1i(u.Loc, int32(data.Unit))
		case gravitygl.SAMPLER_CUBE:
			data := u.Data.(*Texture)
			gravitygl.ActiveTexture(data.Unit)
			gravitygl.BindTexture(data.Target, data.TextureID)
			gravitygl.Uniform1i(u.Loc, int32(data.Unit))
		}
	}
}
