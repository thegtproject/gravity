package mesh

import (
	"github.com/thegtproject/gravity/pkg/core/texture"
	"github.com/thegtproject/gravity/pkg/math/mgl32"
)

// Mesh ...
type Mesh struct {
	Indices   []uint32
	Positions []float32
	Colors    []float32
	Coords    []float32
	Normals   []float32
	Textures  []*texture.Texture
	Target    []int32
}

// NewMesh ...
func NewMesh() *Mesh {
	return &Mesh{}
}

// FromGob ...
func FromGob(path string) (m *Mesh) {
	m = NewMesh()
	err := readGob(path, m)
	if err != nil {
		panic(err)
	}
	return
}

// ToGob ...
func ToGob(m *Mesh, path string) {
	err := writeGob(path, *m)
	if err != nil {
		panic(err)
	}
}

// Scale ...
func (m *Mesh) Scale(f float32) *Mesh {
	for i := range m.Positions {
		m.Positions[i] *= f
	}
	return m
}

// ScaleXYZ ...
func (m *Mesh) ScaleXYZ(x, y, z float32) *Mesh {
	for i := 0; i < len(m.Positions); i += 3 {
		m.Positions[i+0] *= x
		m.Positions[i+1] *= y
		m.Positions[i+2] *= z
	}
	return m
}

// Rotate ...
func (m *Mesh) Rotate() *Mesh {
	return m
}

// GenerateBoundingBoxMeshSolid ...
func (m *Mesh) GenerateBoundingBoxMeshSolid() *Mesh {
	var (
		minx, maxx float32 = 0, 0
		miny, maxy float32 = 0, 0
		minz, maxz float32 = 0, 0
	)

	for i := 0; i < len(m.Positions); i += 3 {
		x := m.Positions[i+0]
		y := m.Positions[i+1]
		z := m.Positions[i+2]

		if x < minx {
			minx = x
		}
		if x > maxx {
			maxx = x
		}

		if y < miny {
			miny = y
		}
		if y > maxy {
			maxy = y
		}

		if z < minz {
			minz = z
		}
		if z > maxz {
			maxz = z
		}
	}

	return &Mesh{
		Positions: []float32{
			minx, miny, minz,
			minx, miny, maxz,
			minx, maxy, minz,
			minx, maxy, maxz,

			maxx, miny, minz,
			maxx, miny, maxz,
			maxx, maxy, minz,
			maxx, maxy, maxz,

			minx, miny, minz,
			minx, miny, maxz,
			maxx, miny, minz,
			maxx, miny, maxz,

			minx, miny, minz,
			minx, maxy, minz,
			maxx, miny, minz,
			maxx, maxy, minz,

			minx, miny, maxz,
			maxx, miny, maxz,
			minx, maxy, maxz,
			maxx, maxy, maxz,

			minx, maxy, minz,
			maxx, maxy, minz,
			minx, maxy, maxz,
			maxx, maxy, maxz,
		},

		Indices: []uint32{
			0, 1, 2, 2, 1, 3,
			4, 5, 6, 6, 5, 7,
			8, 9, 10, 10, 9, 11,
			12, 13, 14, 14, 13, 15,
			16, 17, 18, 18, 17, 19,
			20, 21, 22, 22, 21, 23,
		},

		Colors: []float32{
			0.1, 0.1, 0.7, 1.0,
			0.1, 0.1, 0.7, 1.0,
			0.1, 0.1, 0.7, 1.0,
			0.1, 0.1, 0.7, 1.0,

			0.5, 0.1, 0.1, 1.0,
			0.5, 0.1, 0.1, 1.0,
			0.5, 0.1, 0.1, 1.0,
			0.5, 0.1, 0.1, 1.0,

			0.1, 0.5, 0.1, 1.0,
			0.1, 0.5, 0.1, 1.0,
			0.1, 0.5, 0.1, 1.0,
			0.1, 0.5, 0.1, 1.0,

			0.7, 0.1, 0.7, 1.0,
			0.7, 0.1, 0.7, 1.0,
			0.7, 0.1, 0.7, 1.0,
			0.7, 0.1, 0.7, 1.0,

			0.5, 0.5, 0.1, 1.0,
			0.5, 0.5, 0.1, 1.0,
			0.5, 0.5, 0.1, 1.0,
			0.5, 0.5, 0.1, 1.0,

			0.1, 0.5, 0.5, 1.0,
			0.1, 0.5, 0.5, 1.0,
			0.1, 0.5, 0.5, 1.0,
			0.1, 0.5, 0.5, 1.0,
		},
	}
}

// GenerateBoundingBoxMeshWireframe ...
func (m *Mesh) GenerateBoundingBoxMeshWireframe() *Mesh {
	var (
		minx, maxx float32 = 0, 0
		miny, maxy float32 = 0, 0
		minz, maxz float32 = 0, 0
	)

	for i := 0; i < len(m.Positions); i += 3 {
		x := m.Positions[i+0]
		y := m.Positions[i+1]
		z := m.Positions[i+2]

		if x < minx {
			minx = x
		}
		if x > maxx {
			maxx = x
		}

		if y < miny {
			miny = y
		}
		if y > maxy {
			maxy = y
		}

		if z < minz {
			minz = z
		}
		if z > maxz {
			maxz = z
		}
	}

	return &Mesh{
		Positions: []float32{
			minx, miny, minz,
			maxx, miny, minz,
			maxx, maxy, minz,
			minx, maxy, minz,

			minx, miny, maxz,
			maxx, miny, maxz,
			maxx, maxy, maxz,
			minx, maxy, maxz,
		},

		Indices: []uint32{
			0, 1, 1, 2, 2, 3, 3, 0,
			4, 5, 5, 6, 6, 7, 7, 4,
			0, 4, 1, 5, 2, 6, 3, 7,
		},
		Colors: []float32{
			0.1, 0.1, 0.7, 1.0,
			0.1, 0.1, 0.7, 1.0,
			0.1, 0.1, 0.7, 1.0,
			0.1, 0.1, 0.7, 1.0,
			0.1, 0.1, 0.7, 1.0,
			0.1, 0.1, 0.7, 1.0,
			0.1, 0.1, 0.7, 1.0,
			0.1, 0.1, 0.7, 1.0,
		},
	}
}

// MinMaxPositions ...
func (m *Mesh) MinMaxPositions() (min, max mgl32.Vec3) {
	var (
		minx, maxx float32 = 0, 0
		miny, maxy float32 = 0, 0
		minz, maxz float32 = 0, 0
	)

	for i := 0; i < len(m.Positions); i += 3 {
		x := m.Positions[i+0]
		y := m.Positions[i+1]
		z := m.Positions[i+2]

		if x < minx {
			minx = x
		}
		if x > maxx {
			maxx = x
		}

		if y < miny {
			miny = y
		}
		if y > maxy {
			maxy = y
		}

		if z < minz {
			minz = z
		}
		if z > maxz {
			maxz = z
		}
	}
	return mgl32.Vec3{minx, miny, minz}, mgl32.Vec3{maxx, maxy, maxz}
}
