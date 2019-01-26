package gravity

import (
	"fmt"
	"io/ioutil"

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
	}

	fmt.Print("loading default material programs: ")
	for _, n := range defaultMaterialNames {
		fmt.Print(n, ", ")
		loadMaterialProgram(n)
	}
	fmt.Println("")
}

func loadMaterialProgram(name string) {
	vertFilename := fmt.Sprintf("../assets/shaders/%s.vert.glsl", name)
	fragFilename := fmt.Sprintf("../assets/shaders/%s.frag.glsl", name)
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
