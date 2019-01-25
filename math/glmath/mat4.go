package glmath

import (
	"bytes"
	"fmt"
	"text/tabwriter"
)

// Mat4 represents a 4x4 matrix
type Mat4 [16]float32

// M4 ...
func M4() Mat4 {
	return Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// TargetTo ...
func (m *Mat4) TargetTo(target Mat4) *Mat4 {

	return m
}

// /**
//  * Generates a matrix that makes something look at something else.
//  *
//  * @param {mat4} out mat4 frustum matrix will be written into
//  * @param {vec3} eye Position of the viewer
//  * @param {vec3} center Point the viewer is looking at
//  * @param {vec3} up vec3 pointing up
//  * @returns {mat4} out
//  */
//  export function targetTo(out, eye, target, up) {
// 	let eyex = eye[0],
// 		eyey = eye[1],
// 		eyez = eye[2],
// 		upx = up[0],
// 		upy = up[1],
// 		upz = up[2];

// 	let z0 = eyex - target[0],
// 		z1 = eyey - target[1],
// 		z2 = eyez - target[2];

// 	let len = z0*z0 + z1*z1 + z2*z2;
// 	if (len > 0) {
// 	  len = 1 / Math.sqrt(len);
// 	  z0 *= len;
// 	  z1 *= len;
// 	  z2 *= len;
// 	}

// 	let x0 = upy * z2 - upz * z1,
// 		x1 = upz * z0 - upx * z2,
// 		x2 = upx * z1 - upy * z0;

// 	len = x0*x0 + x1*x1 + x2*x2;
// 	if (len > 0) {
// 	  len = 1 / Math.sqrt(len);
// 	  x0 *= len;
// 	  x1 *= len;
// 	  x2 *= len;
// 	}

// 	out[0] = x0;
// 	out[1] = x1;
// 	out[2] = x2;
// 	out[3] = 0;
// 	out[4] = z1 * x2 - z2 * x1;
// 	out[5] = z2 * x0 - z0 * x2;
// 	out[6] = z0 * x1 - z1 * x0;
// 	out[7] = 0;
// 	out[8] = z0;
// 	out[9] = z1;
// 	out[10] = z2;
// 	out[11] = 0;
// 	out[12] = eyex;
// 	out[13] = eyey;
// 	out[14] = eyez;
// 	out[15] = 1;
// 	return out;
//   };

// CreateInverse ...
func (m *Mat4) CreateInverse() *Mat4 {
	det := m.Determinant()
	if Equal(det, float32(0.0)) {
		return m
	}
	result := &Mat4{
		-m[7]*m[10]*m[13] + m[6]*m[11]*m[13] + m[7]*m[9]*m[14] - m[5]*m[11]*m[14] - m[6]*m[9]*m[15] + m[5]*m[10]*m[15],
		m[3]*m[10]*m[13] - m[2]*m[11]*m[13] - m[3]*m[9]*m[14] + m[1]*m[11]*m[14] + m[2]*m[9]*m[15] - m[1]*m[10]*m[15],
		-m[3]*m[6]*m[13] + m[2]*m[7]*m[13] + m[3]*m[5]*m[14] - m[1]*m[7]*m[14] - m[2]*m[5]*m[15] + m[1]*m[6]*m[15],
		m[3]*m[6]*m[9] - m[2]*m[7]*m[9] - m[3]*m[5]*m[10] + m[1]*m[7]*m[10] + m[2]*m[5]*m[11] - m[1]*m[6]*m[11],
		m[7]*m[10]*m[12] - m[6]*m[11]*m[12] - m[7]*m[8]*m[14] + m[4]*m[11]*m[14] + m[6]*m[8]*m[15] - m[4]*m[10]*m[15],
		-m[3]*m[10]*m[12] + m[2]*m[11]*m[12] + m[3]*m[8]*m[14] - m[0]*m[11]*m[14] - m[2]*m[8]*m[15] + m[0]*m[10]*m[15],
		m[3]*m[6]*m[12] - m[2]*m[7]*m[12] - m[3]*m[4]*m[14] + m[0]*m[7]*m[14] + m[2]*m[4]*m[15] - m[0]*m[6]*m[15],
		-m[3]*m[6]*m[8] + m[2]*m[7]*m[8] + m[3]*m[4]*m[10] - m[0]*m[7]*m[10] - m[2]*m[4]*m[11] + m[0]*m[6]*m[11],
		-m[7]*m[9]*m[12] + m[5]*m[11]*m[12] + m[7]*m[8]*m[13] - m[4]*m[11]*m[13] - m[5]*m[8]*m[15] + m[4]*m[9]*m[15],
		m[3]*m[9]*m[12] - m[1]*m[11]*m[12] - m[3]*m[8]*m[13] + m[0]*m[11]*m[13] + m[1]*m[8]*m[15] - m[0]*m[9]*m[15],
		-m[3]*m[5]*m[12] + m[1]*m[7]*m[12] + m[3]*m[4]*m[13] - m[0]*m[7]*m[13] - m[1]*m[4]*m[15] + m[0]*m[5]*m[15],
		m[3]*m[5]*m[8] - m[1]*m[7]*m[8] - m[3]*m[4]*m[9] + m[0]*m[7]*m[9] + m[1]*m[4]*m[11] - m[0]*m[5]*m[11],
		m[6]*m[9]*m[12] - m[5]*m[10]*m[12] - m[6]*m[8]*m[13] + m[4]*m[10]*m[13] + m[5]*m[8]*m[14] - m[4]*m[9]*m[14],
		-m[2]*m[9]*m[12] + m[1]*m[10]*m[12] + m[2]*m[8]*m[13] - m[0]*m[10]*m[13] - m[1]*m[8]*m[14] + m[0]*m[9]*m[14],
		m[2]*m[5]*m[12] - m[1]*m[6]*m[12] - m[2]*m[4]*m[13] + m[0]*m[6]*m[13] + m[1]*m[4]*m[14] - m[0]*m[5]*m[14],
		-m[2]*m[5]*m[8] + m[1]*m[6]*m[8] + m[2]*m[4]*m[9] - m[0]*m[6]*m[9] - m[1]*m[4]*m[10] + m[0]*m[5]*m[10],
	}
	return result.Multiply(1 / det)
}

// Invert ...
func (m *Mat4) Invert() *Mat4 {
	det := m.Determinant()
	if Equal(det, float32(0.0)) {
		return m
	}
	m[0] = -m[7]*m[10]*m[13] + m[6]*m[11]*m[13] + m[7]*m[9]*m[14] - m[5]*m[11]*m[14] - m[6]*m[9]*m[15] + m[5]*m[10]*m[15]
	m[1] = m[3]*m[10]*m[13] - m[2]*m[11]*m[13] - m[3]*m[9]*m[14] + m[1]*m[11]*m[14] + m[2]*m[9]*m[15] - m[1]*m[10]*m[15]
	m[2] = -m[3]*m[6]*m[13] + m[2]*m[7]*m[13] + m[3]*m[5]*m[14] - m[1]*m[7]*m[14] - m[2]*m[5]*m[15] + m[1]*m[6]*m[15]
	m[3] = m[3]*m[6]*m[9] - m[2]*m[7]*m[9] - m[3]*m[5]*m[10] + m[1]*m[7]*m[10] + m[2]*m[5]*m[11] - m[1]*m[6]*m[11]
	m[4] = m[7]*m[10]*m[12] - m[6]*m[11]*m[12] - m[7]*m[8]*m[14] + m[4]*m[11]*m[14] + m[6]*m[8]*m[15] - m[4]*m[10]*m[15]
	m[5] = -m[3]*m[10]*m[12] + m[2]*m[11]*m[12] + m[3]*m[8]*m[14] - m[0]*m[11]*m[14] - m[2]*m[8]*m[15] + m[0]*m[10]*m[15]
	m[6] = m[3]*m[6]*m[12] - m[2]*m[7]*m[12] - m[3]*m[4]*m[14] + m[0]*m[7]*m[14] + m[2]*m[4]*m[15] - m[0]*m[6]*m[15]
	m[7] = -m[3]*m[6]*m[8] + m[2]*m[7]*m[8] + m[3]*m[4]*m[10] - m[0]*m[7]*m[10] - m[2]*m[4]*m[11] + m[0]*m[6]*m[11]
	m[8] = -m[7]*m[9]*m[12] + m[5]*m[11]*m[12] + m[7]*m[8]*m[13] - m[4]*m[11]*m[13] - m[5]*m[8]*m[15] + m[4]*m[9]*m[15]
	m[9] = m[3]*m[9]*m[12] - m[1]*m[11]*m[12] - m[3]*m[8]*m[13] + m[0]*m[11]*m[13] + m[1]*m[8]*m[15] - m[0]*m[9]*m[15]
	m[10] = -m[3]*m[5]*m[12] + m[1]*m[7]*m[12] + m[3]*m[4]*m[13] - m[0]*m[7]*m[13] - m[1]*m[4]*m[15] + m[0]*m[5]*m[15]
	m[11] = m[3]*m[5]*m[8] - m[1]*m[7]*m[8] - m[3]*m[4]*m[9] + m[0]*m[7]*m[9] + m[1]*m[4]*m[11] - m[0]*m[5]*m[11]
	m[12] = m[6]*m[9]*m[12] - m[5]*m[10]*m[12] - m[6]*m[8]*m[13] + m[4]*m[10]*m[13] + m[5]*m[8]*m[14] - m[4]*m[9]*m[14]
	m[13] = -m[2]*m[9]*m[12] + m[1]*m[10]*m[12] + m[2]*m[8]*m[13] - m[0]*m[10]*m[13] - m[1]*m[8]*m[14] + m[0]*m[9]*m[14]
	m[14] = m[2]*m[5]*m[12] - m[1]*m[6]*m[12] - m[2]*m[4]*m[13] + m[0]*m[6]*m[13] + m[1]*m[4]*m[14] - m[0]*m[5]*m[14]
	m[15] = -m[2]*m[5]*m[8] + m[1]*m[6]*m[8] + m[2]*m[4]*m[9] - m[0]*m[6]*m[9] - m[1]*m[4]*m[10] + m[0]*m[5]*m[10]
	return m.Multiply(1 / det)
}

// Determinant ...
func (m *Mat4) Determinant() float32 {
	return m[0]*m[5]*m[10]*m[15] -
		m[0]*m[5]*m[11]*m[14] -
		m[0]*m[6]*m[9]*m[15] +
		m[0]*m[6]*m[11]*m[13] +
		m[0]*m[7]*m[9]*m[14] -
		m[0]*m[7]*m[10]*m[13] -
		m[1]*m[4]*m[10]*m[15] +
		m[1]*m[4]*m[11]*m[14] +
		m[1]*m[6]*m[8]*m[15] -
		m[1]*m[6]*m[11]*m[12] -
		m[1]*m[7]*m[8]*m[14] +
		m[1]*m[7]*m[10]*m[12] +
		m[2]*m[4]*m[9]*m[15] -
		m[2]*m[4]*m[11]*m[13] -
		m[2]*m[5]*m[8]*m[15] +
		m[2]*m[5]*m[11]*m[12] +
		m[2]*m[7]*m[8]*m[13] -
		m[2]*m[7]*m[9]*m[12] -
		m[3]*m[4]*m[9]*m[14] +
		m[3]*m[4]*m[10]*m[13] +
		m[3]*m[5]*m[8]*m[14] -
		m[3]*m[5]*m[10]*m[12] -
		m[3]*m[6]*m[8]*m[13] +
		m[3]*m[6]*m[9]*m[12]
}

// Multiply ...
func (m *Mat4) Multiply(f float32) *Mat4 {
	m[0] = m[0] * f
	m[1] = m[1] * f
	m[2] = m[2] * f
	m[3] = m[3] * f
	m[4] = m[4] * f
	m[5] = m[5] * f
	m[6] = m[6] * f
	m[7] = m[7] * f
	m[8] = m[8] * f
	m[9] = m[9] * f
	m[10] = m[10] * f
	m[11] = m[11] * f
	m[12] = m[12] * f
	m[13] = m[13] * f
	m[14] = m[14] * f
	m[15] = m[15] * f
	return m
}

func (m *Mat4) String() string {
	buf := new(bytes.Buffer)
	w := tabwriter.NewWriter(buf, 4, 4, 1, ' ', tabwriter.AlignRight)
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			fmt.Fprintf(w, "%f\t", m[row+col*4])
		}
		fmt.Fprintln(w, "")
	}
	w.Flush()
	return buf.String()
}
