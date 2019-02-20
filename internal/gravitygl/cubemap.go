package gravitygl

import (
	"image"
	"unsafe"

	"github.com/go-gl/gl/v4.3-compatibility/gl"
	"github.com/thegtproject/gravity/internal/data"

	"github.com/thegtproject/gravity/pkg/core/texture"
	"github.com/thegtproject/gravity/pkg/mesh"
)

// GL_TEXTURE_CUBE_MAP_POSITIVE_X	Right
// GL_TEXTURE_CUBE_MAP_NEGATIVE_X	Left
// GL_TEXTURE_CUBE_MAP_POSITIVE_Y	Top
// GL_TEXTURE_CUBE_MAP_NEGATIVE_Y	Bottom
// GL_TEXTURE_CUBE_MAP_POSITIVE_Z	Back
// GL_TEXTURE_CUBE_MAP_NEGATIVE_Z	Front

// NewCubeMap ...
func NewCubeMap(PosX, NegX, PosY, NegY, PosZ, NegZ *image.RGBA) uint32 {
	tex := CreateTexture()
	gl.TexImage2D(
		gl.TEXTURE_CUBE_MAP_POSITIVE_X, 0, gl.RGBA8,
		int32(PosX.Bounds().Dx()), int32(PosX.Bounds().Dy()),
		0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&PosX.Pix[0]),
	)
	gl.TexImage2D(
		gl.TEXTURE_CUBE_MAP_NEGATIVE_X, 0, gl.RGBA8,
		int32(NegX.Bounds().Dx()), int32(NegX.Bounds().Dy()),
		0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&NegX.Pix[0]),
	)
	gl.TexImage2D(
		gl.TEXTURE_CUBE_MAP_POSITIVE_Y, 0, gl.RGBA8,
		int32(PosY.Bounds().Dx()), int32(PosY.Bounds().Dy()),
		0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&PosY.Pix[0]),
	)
	gl.TexImage2D(
		gl.TEXTURE_CUBE_MAP_NEGATIVE_Y, 0, gl.RGBA8,
		int32(NegY.Bounds().Dx()), int32(NegY.Bounds().Dy()),
		0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&NegY.Pix[0]),
	)
	gl.TexImage2D(
		gl.TEXTURE_CUBE_MAP_POSITIVE_Z, 0, gl.RGBA8,
		int32(PosZ.Bounds().Dx()), int32(PosZ.Bounds().Dy()),
		0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&PosZ.Pix[0]),
	)
	gl.TexImage2D(
		gl.TEXTURE_CUBE_MAP_NEGATIVE_Z, 0, gl.RGBA8,
		int32(NegZ.Bounds().Dx()), int32(NegZ.Bounds().Dy()),
		0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&NegZ.Pix[0]),
	)

	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE)
	return tex
}

// NewCubeMapMesh ...
func NewCubeMapMesh(PosX, NegX, PosY, NegY, PosZ, NegZ image.Image) *mesh.Mesh {
	return &mesh.Mesh{
		Indices:   mesh.CubemapIndices,
		Positions: mesh.CubemapPositions,
		Colors:    mesh.CubemapColors,
		Coords:    mesh.CubemapCoords,
		Textures: []*texture.Texture{
			NewTextureFromImage(data.TextureDataFromImage(PosX), TEXTURE_CUBE_MAP),
			NewTextureFromImage(data.TextureDataFromImage(NegX), TEXTURE_CUBE_MAP),
			NewTextureFromImage(data.TextureDataFromImage(PosY), TEXTURE_CUBE_MAP),
			NewTextureFromImage(data.TextureDataFromImage(NegX), TEXTURE_CUBE_MAP),
			NewTextureFromImage(data.TextureDataFromImage(PosZ), TEXTURE_CUBE_MAP),
			NewTextureFromImage(data.TextureDataFromImage(NegZ), TEXTURE_CUBE_MAP),
		},
	}
}

// NewCubeMapFromFiles ...
func NewCubeMapFromFiles(PosX, NegX, PosY, NegY, PosZ, NegZ string) (TexPosX, TexNegX, TexPosY, TexNegY, TexPosZ, TexNegZ *image.RGBA) {
	TexPosX = data.TextureDataFromFile(PosX)
	TexNegX = data.TextureDataFromFile(NegX)
	TexPosY = data.TextureDataFromFile(PosY)
	TexNegY = data.TextureDataFromFile(NegY)
	TexPosZ = data.TextureDataFromFile(PosZ)
	TexNegZ = data.TextureDataFromFile(NegZ)
	return
}
