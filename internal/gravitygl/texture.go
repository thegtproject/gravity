package gravitygl

import (
	"image"

	"github.com/go-gl/gl/v4.3-compatibility/gl"
)

// MakeTextureImg ...
func MakeTextureImg(img image.NRGBA) uint32 {
	tex := CreateTexture()

	ActiveTexture(TEXTURE0)
	BindTexture(tex)

	width := int32(img.Bounds().Dx())
	height := int32(img.Bounds().Dy())

	gl.TexImage2D(
		TEXTURE_2D,
		0,
		RGBA,
		width,
		height,
		0,
		RGBA,
		UNSIGNED_BYTE,
		gl.Ptr(img.Pix),
	)

	TexParameteri(tex, TEXTURE_WRAP_S, CLAMP_TO_EDGE)
	TexParameteri(tex, TEXTURE_WRAP_T, CLAMP_TO_EDGE)
	TexParameteri(tex, TEXTURE_MIN_FILTER, LINEAR)
	TexParameteri(tex, TEXTURE_MAG_FILTER, LINEAR)

	//GenerateMipmap(TEXTURE_2D)
	return tex
}

// MakeTexture ...
func MakeTexture(w, h int32) uint32 {

	tex := CreateTexture()

	BindTexture(tex)

	gl.TexImage2D(
		gl.TEXTURE_2D, 0, gl.RGB, w, h, 0, gl.RGB, gl.UNSIGNED_BYTE, nil,
	)

	TexParameteri(tex, TEXTURE_WRAP_S, REPEAT)
	TexParameteri(tex, TEXTURE_WRAP_T, REPEAT)
	TexParameteri(tex, TEXTURE_MIN_FILTER, LINEAR)
	TexParameteri(tex, TEXTURE_MAG_FILTER, LINEAR)

	GenerateMipmap(TEXTURE_2D)
	return tex
}

// CreateTexture ...
func CreateTexture() uint32 {
	var tex uint32
	gl.GenTextures(1, &tex)
	return tex
}

// ActiveTexture ...
func ActiveTexture(tex uint32) {
	gl.ActiveTexture(tex)
}

// BindTexture ...
func BindTexture(tex uint32) {
	gl.BindTexture(gl.TEXTURE_2D, tex)
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
