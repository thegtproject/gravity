package glmath

import "math"

// Epsilon ...
const Epsilon float32 = 1.0E-8

// Variables
var (
	MinNormal = float32(1.1754943508222875e-38) // 1 / 2**(127 - 1)
)

// Abs ...
func Abs(a float32) float32 {
	if a < 0 {
		return -a
	} else if a == 0 {
		return 0
	}

	return a
}

// AlmostEqual ...
func AlmostEqual(a, b, epsilon float32) bool {
	if a == b {
		return true
	}
	diff := Abs(a - b)
	if a*b == 0 || diff < MinNormal {
		return diff < epsilon*epsilon
	}
	return diff/(Abs(a)+Abs(b)) < epsilon
}

// Equal ...
func Equal(a, b float32) bool {
	return AlmostEqual(a, b, Epsilon)
}

// Perspective ...
func Perspective(fovy, aspect, near, far float32) Mat4 {
	nmf, f := near-far, float32(1./math.Tan(float64(fovy)/2.0))
	return Mat4{
		float32(f / aspect), 0, 0, 0,
		0, float32(f), 0, 0,
		0, 0, float32((near + far) / nmf), -1,
		0, 0, float32((2. * far * near) / nmf), 0,
	}
}

// Sqrt ...
func Sqrt(v float32) float32 {
	return float32(math.Sqrt(float64(v)))
}
