package texture

import "image"

// ImageData ...
type ImageData struct {
	*image.RGBA
	W, H int32
}

// Texture ...
type Texture struct {
	Unit           uint32
	Imgdata        []*ImageData
	ID             uint32
	Textureid      uint32
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
