package mesh

import (
	"math"
	"math/rand"
)

// NewSphere ...
func NewSphere(radius float32, slices, stacks int) *Mesh {
	var (
		indices   []uint16
		positions []float32
		colors    []float32
		coords    []float32
		normals   []float32
	)
	const (
		pi  = float32(3.14159265358979323846264338327950288419716939937510582097494459)
		pi2 = 2.0 * pi
	)

	for i := 0; i <= stacks; i++ {
		coordY := float32(i) / float32(stacks)
		phi := coordY * pi
		for j := 0; j <= slices; j++ {
			coordX := float32(j) / float32(slices)
			theta := coordX * pi2

			X := float32(math.Cos(float64(theta))) * float32(math.Sin(float64(phi)))
			Y := float32(math.Cos(float64(phi)))
			Z := float32(math.Sin(float64(theta))) * float32(math.Sin(float64(phi)))

			positions = append(positions, X*radius, Y*radius, Z*radius)
			normals = append(normals, X, Y, Z)
			coords = append(coords, coordX, coordY)
		}
	}

	for i := 0; i < slices*stacks+slices; i++ {
		indices = append(indices, uint16(i))
		indices = append(indices, uint16(i+slices+1))
		indices = append(indices, uint16(i+slices))

		indices = append(indices, uint16(i+slices+1))
		indices = append(indices, uint16(i))
		indices = append(indices, uint16(i+1))
	}

	for i := 0; i < len(positions)/3; i++ {
		r := rand.Float32()
		g := rand.Float32()
		b := rand.Float32()
		a := float32(1.0)
		colors = append(colors, r, g, b, a)
	}

	return &Mesh{
		Indices:   indices,
		Positions: positions,
		Colors:    colors,
		Coords:    coords,
		Normals:   normals,
	}
}
