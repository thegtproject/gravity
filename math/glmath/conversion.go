package glmath

import "math"

// DegToRad ...
func DegToRad(angle float32) float32 {
	return angle * float32(math.Pi) / 180
}

// RadToDeg ...
func RadToDeg(angle float32) float32 {
	return angle * 180 / float32(math.Pi)
}
