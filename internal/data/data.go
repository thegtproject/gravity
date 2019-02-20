package data

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
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

// TextureDataFromFile ...
func TextureDataFromFile(filename string) *image.RGBA {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	return TextureDataFromImage(img)
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
