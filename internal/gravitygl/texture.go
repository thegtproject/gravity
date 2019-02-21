package gravitygl

import (
	"image"
	"image/jpeg"
	"image/png"
	"unsafe"

	"github.com/go-gl/gl/v4.3-compatibility/gl"
	"github.com/thegtproject/gravity/internal/data"
	"github.com/thegtproject/gravity/pkg/core/texture"
)

// // MakeTextureImg ...
// func MakeTextureImg(img image.NRGBA) uint32 {
// 	tex := CreateTexture()

// 	ActiveTexture(TEXTURE0)
// 	BindTexture(TEXTURE_2D, tex)

// 	width := int32(img.Bounds().Dx())
// 	height := int32(img.Bounds().Dy())

// 	gl.TexImage2D(
// 		TEXTURE_2D,
// 		0,
// 		RGBA,
// 		width,
// 		height,
// 		0,
// 		RGBA,
// 		UNSIGNED_BYTE,
// 		gl.Ptr(img.Pix),
// 	)

// 	TexParameteri(tex, TEXTURE_WRAP_S, CLAMP_TO_EDGE)
// 	TexParameteri(tex, TEXTURE_WRAP_T, CLAMP_TO_EDGE)
// 	TexParameteri(tex, TEXTURE_MIN_FILTER, LINEAR)
// 	TexParameteri(tex, TEXTURE_MAG_FILTER, LINEAR)

// 	//GenerateMipmap(TEXTURE_2D)
// 	return tex
// }

// TexImage2D ...
func TexImage2D(target uint32, level int32, internalformat int32, width int32, height int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	gl.TexImage2D(target, level, internalformat, width, height, border, format, xtype, pixels)
}

// CreateTexture ...
func CreateTexture() uint32 {
	var tex uint32
	gl.GenTextures(1, &tex)
	return tex
}

// ActiveTexture ...
func ActiveTexture(unit uint32) {
	gl.ActiveTexture(TEXTURE0 + unit)
}

// BindTexture ...
func BindTexture(target, tex uint32) {
	gl.BindTexture(target, tex)
}

// TexParameteri ...
func TexParameteri(tex uint32, pname uint32, param int32) {
	gl.TextureParameteri(tex, pname, param)
}

// GenerateMipmap ...
func GenerateMipmap(target uint32) {
	gl.GenerateMipmap(target)
}

// FramebufferTexture2D ...
func FramebufferTexture2D(tex uint32) {
	gl.FramebufferTexture2D(
		gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, tex, 0,
	)
}

// PixelStorei ...
func PixelStorei(pname uint32, param int32) {
	gl.PixelStorei(pname, param)
}

var _, _ = png.BestSpeed, jpeg.DefaultQuality

// NewImageDataFromImage ...
func NewImageDataFromImage(img image.Image) *texture.ImageData {
	return &texture.ImageData{
		RGBA: data.TextureDataFromImage(img),
		H:    int32(img.Bounds().Dx()),
		W:    int32(img.Bounds().Dy()),
	}
}

// NewImageDataFromFile ...
func NewImageDataFromFile(filename string) *texture.ImageData {
	img := data.TextureRGBAFromFile(filename)
	return &texture.ImageData{
		RGBA: img,
		H:    int32(img.Bounds().Dx()),
		W:    int32(img.Bounds().Dy()),
	}
}

// UploadToGPU ...
func UploadToGPU(t *texture.Texture) {

	ActiveTexture(t.Unit)
	BindTexture(t.Target, t.Textureid)
	// glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_BASE_LEVEL, 0)
	// glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAX_LEVEL, 0)

	// TODO: https://www.khronos.org/opengl/wiki/Image_Formats#Required_formats

	if t.Target == TEXTURE_CUBE_MAP {
		CubeMapTargets := []uint32{
			TEXTURE_CUBE_MAP_POSITIVE_X,
			TEXTURE_CUBE_MAP_NEGATIVE_X,
			TEXTURE_CUBE_MAP_POSITIVE_Y,
			TEXTURE_CUBE_MAP_NEGATIVE_Y,
			TEXTURE_CUBE_MAP_POSITIVE_Z,
			TEXTURE_CUBE_MAP_NEGATIVE_Z,
		}
		for target, imgdata := range t.Imgdata {
			TexImage2D(
				CubeMapTargets[target], t.Mips, t.Format,
				imgdata.W, imgdata.H, 0, t.Originalformat,
				UNSIGNED_BYTE, unsafe.Pointer(&imgdata.Pix[0]),
			)
		}
	} else {
		for _, imgdata := range t.Imgdata {
			TexImage2D(
				t.Target, t.Mips, t.Format, imgdata.W, imgdata.H, 0,
				t.Originalformat, UNSIGNED_BYTE, unsafe.Pointer(&imgdata.Pix[0]),
			)
		}
	}

	TexParameteri(t.Textureid, TEXTURE_WRAP_S, t.WrapS)
	TexParameteri(t.Textureid, TEXTURE_WRAP_T, t.WrapT)
	TexParameteri(t.Textureid, TEXTURE_WRAP_R, t.WrapR)
	TexParameteri(t.Textureid, TEXTURE_MIN_FILTER, t.MinFilter)
	TexParameteri(t.Textureid, TEXTURE_MAG_FILTER, t.MagFilter)
}

// NewTextureFromFile ...
func NewTextureFromFile(target uint32, filename ...string) *texture.Texture {
	var images []image.Image
	for _, file := range filename {
		images = append(images, data.TextureImageFromFile(file))
	}
	return NewTextureFromImage(target, images...)
}

// NewTextureFromImage ...
func NewTextureFromImage(target uint32, img ...image.Image) *texture.Texture {
	t := &texture.Texture{
		Unit:           0,
		ID:             0,
		Target:         target,
		Textureid:      CreateTexture(),
		Mips:           0,
		Format:         RGBA8,
		Originalformat: RGBA,
	}
	for _, im := range img {
		t.Imgdata = append(t.Imgdata, NewImageDataFromImage(im))
	}
	return t
}

// if(depth_buffer_precision == 16)
// {
//   GLushort mypixels[width*height];
//   glReadPixels(0, 0, width, height, GL_DEPTH_COMPONENT, GL_UNSIGNED_SHORT, mypixels);
// }
// else if(depth_buffer_precision == 24)
// {
//   GLuint mypixels[width*height];    //There is no 24 bit variable, so we'll have to settle for 32 bit
//   glReadPixels(0, 0, width, height, GL_DEPTH_COMPONENT, GL_UNSIGNED_INT_24_8, mypixels);  //No upconversion.
// }
// else if(depth_buffer_precision == 32)
// {
//   GLuint mypixels[width*height];
//   glReadPixels(0, 0, width, height, GL_DEPTH_COMPONENT, GL_UNSIGNED_INT, mypixels);
// }
// If you have a depth/stencil format, you can get the depth/stencil data this way:

//  GLuint mypixels[width*height];
//  glReadPixels(0, 0, width, height, GL_DEPTH_STENCIL, GL_UNSIGNED_INT_24_8, mypixels);

//---------------------------

// Better code would be to use texture storage functions (if you have OpenGL 4.2 or ARB_texture_storage) to allocate the texture's storage, then upload with glTexSubImage2D:

// glGenTextures(1, &textureID);
// glBindTexture(GL_TEXTURE_2D, textureID);
// glTexStorage2D(GL_TEXTURE_2D, 1, GL_RGBA8, width, height);
// glTexSubImage2D(GL_TEXTURE_2D, 0​, 0, 0, width​, height​, GL_BGRA, GL_UNSIGNED_BYTE, pixels);

// ------------------------------------------

// Automatic mipmap generation
// Mipmaps of a texture can be automatically generated with the glGenerateMipmap function. OpenGL 3.0 or greater is required for this function (or the extension GL_ARB_framebuffer_object). The function works quite simply; when you call it for a texture, mipmaps are generated for that texture:

// glGenTextures(1, &textureID);
// glBindTexture(GL_TEXTURE_2D, textureID);
// glTexStorage2D(GL_TEXTURE_2D, num_mipmaps, GL_RGBA8, width, height);
// glTexSubImage2D(GL_TEXTURE_2D, 0​, 0, 0, width​, height​, GL_BGRA, GL_UNSIGNED_BYTE, pixels);
// glGenerateMipmap(GL_TEXTURE_2D);  //Generate num_mipmaps number of mipmaps here.
// glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_S, GL_REPEAT);
// glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_T, GL_REPEAT);
// glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_LINEAR);
// glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR_MIPMAP_LINEAR);
// If texture storage is not available, you can use the older API:

// glGenTextures(1, &textureID);
// glBindTexture(GL_TEXTURE_2D, textureID);
// glTexImage2D(GL_TEXTURE_2D, 0, GL_RGBA8, width, height, 0, GL_BGRA, GL_UNSIGNED_BYTE, pixels);
// glGenerateMipmap(GL_TEXTURE_2D);  //Generate mipmaps now!!!
// glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_S, GL_REPEAT);
// glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_T, GL_REPEAT);
// glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_LINEAR);
// glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR_MIPMAP_LINEAR);
