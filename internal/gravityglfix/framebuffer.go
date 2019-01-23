package gravitygl

import (
	"github.com/go-gl/gl/v4.3-compatibility/gl"
)

// MakeFramebuffer ...
func MakeFramebuffer() uint32 {
	fb := CreateFramebuffer()
	return fb
}

// CreateFramebuffer ...
func CreateFramebuffer() uint32 {
	var fb uint32
	gl.GenFramebuffers(1, &fb)
	return fb
}

// BindFramebuffer ...
func BindFramebuffer(fb uint32) {
	gl.BindFramebuffer(gl.FRAMEBUFFER, fb)
}

// CreateRenderbuffer ...
func CreateRenderbuffer() uint32 {
	var rb uint32
	gl.GenRenderbuffers(1, &rb)
	return rb
}

// BindRenderbuffer ...
func BindRenderbuffer(rb uint32) {
	gl.BindRenderbuffer(gl.RENDERBUFFER, rb)
}

// RenderbufferStorage ...
func RenderbufferStorage(w, h int32) {
	gl.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH24_STENCIL8, w, h)
}

// FramebufferRenderbuffer ...
func FramebufferRenderbuffer(rb uint32) {
	gl.FramebufferRenderbuffer(
		gl.FRAMEBUFFER, gl.DEPTH_STENCIL_ATTACHMENT, gl.RENDERBUFFER, rb,
	)
}

// // TexParameteri ...
// func TexParameteri(tex uint32, pname uint32, param int32) {
// 	gl.TextureParameteri(tex, pname, param)
// }

// // GenerateMipmap ...
// func GenerateMipmap(target uint32) {
// 	gl.GenerateMipmap(target)
// }
