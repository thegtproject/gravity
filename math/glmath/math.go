package glmath

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
