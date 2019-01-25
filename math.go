package gravity

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/thegtproject/gravity/common"
)

type (
	// Vec3 type alias
	Vec3 = mgl32.Vec3
	// Vec4 type alias
	Vec4 = mgl32.Vec4
	// Mat4 type alias
	Mat4 = mgl32.Mat4
	// Quat type alias
	Quat = mgl32.Quat
)

type (
	// XVec3 = mgl32.Vec3
	// XVec4 = mgl32.Vec4
	XMat4 = common.XMat4
	// XQuat = mgl32.Quat
)

// func fromRotationTranslation(out *Mat4, q Quat, v Vec3) {
// 	// Quaternion math
// 	x, y, z, w := q.V[0], q.V[1], q.V[2], q.W
// 	x2 := x + x
// 	y2 := y + y
// 	z2 := z + z

// 	xx := x * x2
// 	xy := x * y2
// 	xz := x * z2
// 	yy := y * y2
// 	yz := y * z2
// 	zz := z * z2
// 	wx := w * x2
// 	wy := w * y2
// 	wz := w * z2

// 	out[0] = 1 - (yy + zz)
// 	out[1] = xy + wz
// 	out[2] = xz - wy
// 	out[3] = 0
// 	out[4] = xy - wz
// 	out[5] = 1 - (xx + zz)
// 	out[6] = yz + wx
// 	out[7] = 0
// 	out[8] = xz + wy
// 	out[9] = yz - wx
// 	out[10] = 1 - (xx + yy)
// 	out[11] = 0
// 	out[12] = v[0]
// 	out[13] = v[1]
// 	out[14] = v[2]
// 	out[15] = 1

// }

// /**
//  * Creates a matrix from a quaternion rotation and vector translation
//  * This is equivalent to (but much faster than):
//  *
//  *     mat4.identity(dest);
//  *     mat4.translate(dest, vec);
//  *     quatMat := mat4.create();
//  *     quat4.toMat4(quat, quatMat);
//  *     mat4.multiply(dest, quatMat);
//  *
//  * @param {mat4} out mat4 receiving operation result
//  * @param {quat4} q Rotation quaternion
//  * @param {vec3} v Translation vector
//  * @returns {mat4} out
//  */
//  export function fromRotationTranslation(out, q, v) {
// 	// Quaternion math
// 	x := q[0], y = q[1], z = q[2], w = q[3];
// 	x2 := x + x;
// 	y2 := y + y;
// 	z2 := z + z;

// 	xx := x * x2;
// 	xy := x * y2;
// 	xz := x * z2;
// 	yy := y * y2;
// 	yz := y * z2;
// 	zz := z * z2;
// 	wx := w * x2;
// 	wy := w * y2;
// 	wz := w * z2;

// 	out[0] = 1 - (yy + zz);
// 	out[1] = xy + wz;
// 	out[2] = xz - wy;
// 	out[3] = 0;
// 	out[4] = xy - wz;
// 	out[5] = 1 - (xx + zz);
// 	out[6] = yz + wx;
// 	out[7] = 0;
// 	out[8] = xz + wy;
// 	out[9] = yz - wx;
// 	out[10] = 1 - (xx + yy);
// 	out[11] = 0;
// 	out[12] = v[0];
// 	out[13] = v[1];
// 	out[14] = v[2];
// 	out[15] = 1;

// 	return out;
//   }
