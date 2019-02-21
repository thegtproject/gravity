package gravitygl

import (
	"github.com/thegtproject/gravity/pkg/core/texture"
)

// NewCubeMap ...
func NewCubeMap(PosX, NegX, PosY, NegY, PosZ, NegZ string) *texture.Texture {
	tex := NewTextureFromFile(
		TEXTURE_CUBE_MAP, PosX, NegX, PosY, NegY, PosZ, NegZ,
	)
	tex.MinFilter = LINEAR
	tex.MagFilter = LINEAR
	tex.WrapS = CLAMP_TO_EDGE
	tex.WrapT = CLAMP_TO_EDGE
	tex.WrapR = CLAMP_TO_EDGE
	// tex.Format = RGB
	// tex.Originalformat = RGB

	UploadToGPU(tex)

	return tex
}

// // GenerateCubeMap ...
// func GenerateCubeMap(PosX, NegX, PosY, NegY, PosZ, NegZ *image.RGBA) uint32 {
// 	tex := CreateTexture()
// 	BindTexture(TEXTURE_CUBE_MAP, tex)
// 	gl.TexImage2D(
// 		gl.TEXTURE_CUBE_MAP_POSITIVE_X, 0, gl.RGBA8,
// 		int32(PosX.Bounds().Dx()), int32(PosX.Bounds().Dy()),
// 		0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&PosX.Pix[0]),
// 	)
// 	gl.TexImage2D(
// 		gl.TEXTURE_CUBE_MAP_NEGATIVE_X, 0, gl.RGBA8,
// 		int32(NegX.Bounds().Dx()), int32(NegX.Bounds().Dy()),
// 		0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&NegX.Pix[0]),
// 	)
// 	gl.TexImage2D(
// 		gl.TEXTURE_CUBE_MAP_POSITIVE_Y, 0, gl.RGBA8,
// 		int32(PosY.Bounds().Dx()), int32(PosY.Bounds().Dy()),
// 		0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&PosY.Pix[0]),
// 	)
// 	gl.TexImage2D(
// 		gl.TEXTURE_CUBE_MAP_NEGATIVE_Y, 0, gl.RGBA8,
// 		int32(NegY.Bounds().Dx()), int32(NegY.Bounds().Dy()),
// 		0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&NegY.Pix[0]),
// 	)
// 	gl.TexImage2D(
// 		gl.TEXTURE_CUBE_MAP_POSITIVE_Z, 0, gl.RGBA8,
// 		int32(PosZ.Bounds().Dx()), int32(PosZ.Bounds().Dy()),
// 		0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&PosZ.Pix[0]),
// 	)
// 	gl.TexImage2D(
// 		gl.TEXTURE_CUBE_MAP_NEGATIVE_Z, 0, gl.RGBA8,
// 		int32(NegZ.Bounds().Dx()), int32(NegZ.Bounds().Dy()),
// 		0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&NegZ.Pix[0]),
// 	)

// 	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
// 	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
// 	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
// 	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
// 	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE)

// 	return tex
// }

// // NewCubeMapMesh ...
// func NewCubeMapMesh(PosX, NegX, PosY, NegY, PosZ, NegZ image.Image) (cubeMesh *mesh.Mesh, texID uint32) {
// 	m := &mesh.Mesh{
// 		Indices:   mesh.CubemapIndices,
// 		Positions: mesh.CubemapPositions,
// 		Colors:    mesh.CubemapColors,
// 		Coords:    mesh.CubemapCoords,
// 	}
// 	tex := GenerateCubeMap(
// 		data.TextureDataFromImage(PosX),
// 		data.TextureDataFromImage(NegX),
// 		data.TextureDataFromImage(PosY),
// 		data.TextureDataFromImage(NegX),
// 		data.TextureDataFromImage(PosZ),
// 		data.TextureDataFromImage(NegZ),
// 	)
// 	return m, tex
// }

// // NewCubeMapFromFiles ...
// func NewCubeMapFromFiles(PosX, NegX, PosY, NegY, PosZ, NegZ string) (cubeMesh *mesh.Mesh, texID uint32) {
// 	TexPosX := data.TextureDataFromFile(PosX)
// 	TexNegX := data.TextureDataFromFile(NegX)
// 	TexPosY := data.TextureDataFromFile(PosY)
// 	TexNegY := data.TextureDataFromFile(NegY)
// 	TexPosZ := data.TextureDataFromFile(PosZ)
// 	TexNegZ := data.TextureDataFromFile(NegZ)
// 	return NewCubeMapMesh(TexPosX, TexNegX, TexPosY, TexNegY, TexPosZ, TexNegZ)
// }
