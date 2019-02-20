package gravitygl

import (
	"unsafe"

	"github.com/go-gl/gl/v4.3-compatibility/gl"
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
