package data

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"github.com/ftrvxmtrx/tga"
)

var _ = png.BestCompression
var _ = jpeg.DefaultQuality

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
