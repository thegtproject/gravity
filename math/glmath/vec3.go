package glmath

// Vec3 ...
type Vec3 [3]float32

// V3 ...
func V3() *Vec3 {
	return &Vec3{0, 0, 0}
}

// Add ...
func (vec *Vec3) Add(v *Vec3) *Vec3 {
	vec[0] += v[0]
	vec[1] += v[1]
	vec[2] += v[2]
	return vec
}

// Sub ...
func (vec *Vec3) Sub(v *Vec3) *Vec3 {
	vec[0] -= v[0]
	vec[1] -= v[1]
	vec[2] -= v[2]
	return vec
}

// Multiply ...
func (vec *Vec3) Multiply(f float32) *Vec3 {
	vec[0] *= f
	vec[1] *= f
	vec[2] *= f
	return vec
}

// CreateCross ...
func (vec *Vec3) CreateCross(v Vec3) Vec3 {
	return Vec3{vec[1]*v[2] - vec[2]*v[1], vec[2]*v[0] - vec[0]*v[2], vec[0]*v[1] - vec[1]*v[0]}
}

// Clone ...
func (vec *Vec3) Clone() Vec3 {
	return Vec3{vec[0], vec[1], vec[2]}
}
