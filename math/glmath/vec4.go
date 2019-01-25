package glmath

// Vec4 ...
type Vec4 [4]float32

// V4 ...
func V4() Vec4 {
	return Vec4{0, 0, 0}
}

// Add ...
func (vec *Vec4) Add(v *Vec4) *Vec4 {
	vec[0] += v[0]
	vec[1] += v[1]
	vec[2] += v[2]
	vec[3] += v[3]
	return vec
}

// Multiply ...
func (vec *Vec4) Multiply(f float32) *Vec4 {
	vec[0] *= f
	vec[1] *= f
	vec[2] *= f
	vec[3] *= f
	return vec
}
