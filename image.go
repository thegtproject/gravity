package gravity

import (
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"

	"github.com/ftrvxmtrx/tga"
)

// ImageData ...
type ImageData struct {
	Pix    []uint8
	Width  int32
	Height int32
	Stride int
}

// NewImageDataFromImage ...
func NewImageDataFromImage(img image.Image) *ImageData {
	return &ImageData{
		Pix:    TextureDataFromImage(img).Pix,
		Width:  int32(img.Bounds().Dy()),
		Height: int32(img.Bounds().Dx()),
	}
}

// NewImageDataFromFile ...
func NewImageDataFromFile(filename string) *ImageData {
	img := TextureRGBAFromFile(filename)
	return &ImageData{
		Pix:    img.Pix,
		Width:  int32(img.Bounds().Dy()),
		Height: int32(img.Bounds().Dx()),
	}
}

// NewTextureFromFile ...
func NewTextureFromFile(target uint32, filename ...string) *Texture {
	var images []image.Image
	for _, file := range filename {
		images = append(images, TextureImageFromFile(file))
	}
	return NewTextureFromImage(target, images...)
}

// TextureDataFromImage ...
// GL texture data really should aways be 4 wide (rgba). Fix this
// function to convert it if necessary.
func TextureDataFromImage(img image.Image) *image.RGBA {
	rgba := image.NewRGBA(img.Bounds())

	// TODO: Make sure image is power of 2

	draw.Draw(rgba, rgba.Bounds(), img, img.Bounds().Min, draw.Src)
	verticalFlip(rgba)
	return rgba
}

// TextureRGBAFromFile ...
func TextureRGBAFromFile(filename string) *image.RGBA {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ext := filepath.Ext(filename)
	if ext == ".jpeg" || ext == ".jpg" || ext == ".png" {
		img, _, err := image.Decode(f)
		if err != nil {
			panic(err)
		}
		return TextureDataFromImage(img)
	}

	if ext == ".tga" {
		img, err := tga.Decode(f)
		if err != nil {
			panic(err)
		}
		return TextureDataFromImage(img)
	}
	panic("unsupported texture type")
}

// TextureImageFromFile ...
func TextureImageFromFile(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ext := filepath.Ext(filename)
	if ext == ".jpeg" || ext == ".jpg" || ext == ".png" {
		img, _, err := image.Decode(f)
		if err != nil {
			panic(err)
		}
		return img
	}

	if ext == ".tga" {
		img, err := tga.Decode(f)
		if err != nil {
			panic(err)
		}
		return img
	}
	panic("unsupported texture type")
}

func verticalFlip(rgba *image.RGBA) {
	bounds := rgba.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	tmpRow := make([]uint8, width*4)
	for i, j := 0, height-1; i < j; i, j = i+1, j-1 {
		iRow := rgba.Pix[i*rgba.Stride : i*rgba.Stride+width*4]
		jRow := rgba.Pix[j*rgba.Stride : j*rgba.Stride+width*4]

		copy(tmpRow, iRow)
		copy(iRow, jRow)
		copy(jRow, tmpRow)
	}
}
