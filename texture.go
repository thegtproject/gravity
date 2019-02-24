package gravity

import (
	"image"
	"unsafe"

	"github.com/thegtproject/gravity/internal/gravitygl"
)

// TextureParameter ...
type TextureParameter = uint32

// Texture Paramaters
const (
	WrapS     TextureParameter = gravitygl.TEXTURE_WRAP_S
	WrapT     TextureParameter = gravitygl.TEXTURE_WRAP_T
	WrapR     TextureParameter = gravitygl.TEXTURE_WRAP_R
	MinFilter TextureParameter = gravitygl.TEXTURE_MIN_FILTER
	MagFilter TextureParameter = gravitygl.TEXTURE_MAG_FILTER
)

// Texture ...
type Texture struct {
	ImageData      []*ImageData
	Unit           uint32
	ID             uint32
	TextureID      uint32
	Target         uint32
	Mips           int32
	Format         int32
	Originalformat uint32
	MagFilter      int32
	MinFilter      int32
	WrapS          int32
	WrapT          int32
	WrapR          int32
}

// UploadToGPU ...
func UploadToGPU(t *Texture) {

	gravitygl.ActiveTexture(t.Unit)
	gravitygl.BindTexture(t.Target, t.TextureID)
	// glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_BASE_LEVEL, 0)
	// glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAX_LEVEL, 0)

	// TODO: https://www.khronos.org/opengl/wiki/Image_Formats#Required_formats

	if t.Target == gravitygl.TEXTURE_CUBE_MAP {
		CubeMapTargets := []uint32{
			gravitygl.TEXTURE_CUBE_MAP_POSITIVE_X,
			gravitygl.TEXTURE_CUBE_MAP_NEGATIVE_X,
			gravitygl.TEXTURE_CUBE_MAP_POSITIVE_Y,
			gravitygl.TEXTURE_CUBE_MAP_NEGATIVE_Y,
			gravitygl.TEXTURE_CUBE_MAP_POSITIVE_Z,
			gravitygl.TEXTURE_CUBE_MAP_NEGATIVE_Z,
		}
		for target, imgdata := range t.ImageData {
			gravitygl.TexImage2D(
				CubeMapTargets[target], t.Mips, t.Format,
				imgdata.Width, imgdata.Height, 0, t.Originalformat,
				gravitygl.UNSIGNED_BYTE, unsafe.Pointer(&imgdata.Pix[0]),
			)
		}
	} else {
		for _, imgdata := range t.ImageData {
			gravitygl.TexImage2D(
				t.Target, t.Mips, t.Format, imgdata.Width, imgdata.Height, 0,
				t.Originalformat, gravitygl.UNSIGNED_BYTE, unsafe.Pointer(&imgdata.Pix[0]),
			)
		}
	}

	gravitygl.TexParameteri(t.TextureID, WrapS, t.WrapS)
	gravitygl.TexParameteri(t.TextureID, WrapT, t.WrapT)
	gravitygl.TexParameteri(t.TextureID, WrapR, t.WrapR)
	gravitygl.TexParameteri(t.TextureID, MinFilter, t.MinFilter)
	gravitygl.TexParameteri(t.TextureID, MagFilter, t.MagFilter)
}

// NewTextureFromImage ...
func NewTextureFromImage(target uint32, img ...image.Image) *Texture {
	t := &Texture{
		Unit:           0,
		ID:             0,
		Target:         target,
		TextureID:      gravitygl.CreateTexture(),
		Mips:           0,
		Format:         gravitygl.RGBA8,
		Originalformat: gravitygl.RGBA,
		WrapS:          gravitygl.CLAMP_TO_EDGE,
		WrapT:          gravitygl.CLAMP_TO_EDGE,
		WrapR:          gravitygl.CLAMP_TO_EDGE,
		MinFilter:      gravitygl.LINEAR,
		MagFilter:      gravitygl.LINEAR,
	}
	for _, im := range img {
		t.ImageData = append(t.ImageData, NewImageDataFromImage(im))
	}
	return t
}

// NewCubeMap ...
func NewCubeMap(PosX, NegX, PosY, NegY, PosZ, NegZ string) *Texture {
	tex := NewTextureFromFile(
		gravitygl.TEXTURE_CUBE_MAP, PosX, NegX, PosY, NegY, PosZ, NegZ,
	)
	tex.MinFilter = gravitygl.LINEAR
	tex.MagFilter = gravitygl.LINEAR
	tex.WrapS = gravitygl.CLAMP_TO_EDGE
	tex.WrapT = gravitygl.CLAMP_TO_EDGE
	tex.WrapR = gravitygl.CLAMP_TO_EDGE
	// tex.Format = RGB
	// tex.Originalformat = RGB

	UploadToGPU(tex)

	return tex
}
