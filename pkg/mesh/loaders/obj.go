package loaders

import (
	"io/ioutil"

	"github.com/thegtproject/gravity/pkg/mesh"
	"github.com/udhos/gwob"
)

var m gwob.MaterialLib
var opts = &gwob.ObjParserOptions{
	IgnoreNormals: false,
}

// LoadObj ...
func LoadObj(filename string) *mesh.Mesh {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	obj, err := gwob.NewObjFromBuf("objName", buf, opts)
	if err != nil {
		panic(err)
	}

	return gwobMeshToMesh(obj)
}

func gwobMeshToMesh(obj *gwob.Obj) *mesh.Mesh {
	m := mesh.NewMesh()

	m.Indices = make([]uint32, len(obj.Indices))
	m.Positions = make([]float32, obj.NumberOfElements()*3)
	m.Coords = make([]float32, obj.NumberOfElements()*2)
	m.Colors = make([]float32, obj.NumberOfElements()*4)

	var i int
	for i = 0; i < obj.NumberOfElements(); i++ {
		vertex := obj.Coord[i*5 : i*5+5]
		m.Positions[i*3+0] = vertex[0]
		m.Positions[i*3+1] = vertex[1]
		m.Positions[i*3+2] = vertex[2]
		m.Coords[i*2+0] = vertex[3]
		m.Coords[i*2+1] = vertex[4]
	}
	for i := range obj.Indices {
		if obj.Indices[i] < 0 {
			panic("wtf")
		}
		m.Indices[i] = uint32(obj.Indices[i])
	}

	min, max := m.MinMaxPositions()
	minZ, maxZ := min.Z(), max.Z()
	z := maxZ - minZ
	for i := 0; i < len(m.Positions)/3; i++ {
		n := m.Positions[i*3+2] / z
		m.Colors[i*4+0] = 0.2
		m.Colors[i*4+1] = n
		m.Colors[i*4+2] = 0.2
		m.Colors[i*4+3] = 1.0
	}

	return m
}
