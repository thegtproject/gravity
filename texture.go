package gravity

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"unsafe"

	"github.com/thegtproject/gravity/internal/data"

	"github.com/thegtproject/gravity/internal/gravitygl"
)

var _, _ = png.BestSpeed, jpeg.DefaultQuality

// ImageData ...
type ImageData struct {
	*image.RGBA

	w, h int32
}

// Texture ...
type Texture struct {
	Unit           uint32
	imgdata        ImageData
	id             uint32
	textureid      uint32
	target         uint32
	mips           int32
	format         int32
	originalformat uint32
}

// UploadToGPU ...
func (t *Texture) UploadToGPU() {

	gravitygl.ActiveTexture(t.Unit)
	gravitygl.BindTexture(t.target, t.textureid)
	// glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_BASE_LEVEL, 0)
	// glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAX_LEVEL, 0)

	gravitygl.TexParameteri(t.textureid, gravitygl.TEXTURE_WRAP_S, gravitygl.REPEAT)
	gravitygl.TexParameteri(t.textureid, gravitygl.TEXTURE_WRAP_T, gravitygl.REPEAT)
	gravitygl.TexParameteri(t.textureid, gravitygl.TEXTURE_MIN_FILTER, gravitygl.LINEAR)
	gravitygl.TexParameteri(t.textureid, gravitygl.TEXTURE_MAG_FILTER, gravitygl.LINEAR)

	// TODO: https://www.khronos.org/opengl/wiki/Image_Formats#Required_formats
	gravitygl.TexImage2D(t.target, t.mips, t.format, t.imgdata.w, t.imgdata.h, 0, t.originalformat, gravitygl.UNSIGNED_BYTE, unsafe.Pointer(&t.imgdata.Pix[0]))
}

// NewTextureFromFile ...
func NewTextureFromFile(path string) *Texture {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	return NewTextureFromImage(img)
}

// NewTextureFromImage ...
func NewTextureFromImage(img image.Image) *Texture {
	t := &Texture{
		Unit: 0,

		id: 0,
		imgdata: ImageData{
			RGBA: data.TextureDataFromImage(img),
			h:    int32(img.Bounds().Dx()),
			w:    int32(img.Bounds().Dy()),
		},
		target:         gravitygl.TEXTURE_2D,
		textureid:      gravitygl.CreateTexture(),
		mips:           0,
		format:         gravitygl.RGBA8,
		originalformat: gravitygl.RGBA,
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
