package main

import (
	"image"
	"image/png"
	"os"

	"github.com/thegtproject/gravity"
)

var img image.Image

func main() {
	pix := getpix()
	_ = pix
	//  for i := 0; i < len(pix); i += 4 {
	// r := pix[i+0]
	// g := pix[i+1]
	// b := pix[i+2]
	// a := pix[i+3]

	// tot := r + g + b + a

	// }

	nimg := image.NewRGBA(img.Bounds())
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			nimg.Set(x, y, img.At(x, y))
		}
	}

	for i := 0; i < len(nimg.Pix); i += 4 {
		r := nimg.Pix[i+0]
		g := nimg.Pix[i+1]
		b := nimg.Pix[i+2]
		a := nimg.Pix[i+3]
		_ = a
		tot := r + g + b
		if tot == 0 {
			tot = 1
		}

		// nimg.Pix[i+0] = r / 3 * 255
		// nimg.Pix[i+1] = g / 3 * 255
		// nimg.Pix[i+2] = b / 3 * 255

		nimg.Pix[i] = r
		nimg.Pix[i+1] = 0
		nimg.Pix[i+2] = 0

		nimg.Pix[i+3] = 255
	}

	f, err := os.Create("C:/dev/go/src/github.com/thegtproject/gravity/developing/assets/terr/highR.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, nimg)

}

func savepix(data []uint32) {
	newimg := image.NewRGBA(img.Bounds())
	for i := range data {
		newimg.Pix[i] = (img.(*image.NRGBA64)).Pix[i]
	}
	f, err := os.Create("C:/dev/go/src/github.com/thegtproject/gravity/developing/assets/terr/asdasd.png")
	if err != nil {
		panic(err)
	}
	png.Encode(f, newimg)
}

func getpix() []byte {
	f, err := os.Open("C:/dev/go/src/github.com/thegtproject/gravity/developing/assets/terr/height.png")
	if err != nil {
		panic(err)
	}
	img, _, err = image.Decode(f)
	if err != nil {
		panic(err)
	}

	switch img.(type) {
	case *image.RGBA:
		return img.(*image.RGBA).Pix
	case *image.NRGBA:
		return img.(*image.NRGBA).Pix
	case *image.NRGBA64:
		return img.(*image.NRGBA64).Pix
	case *image.Gray16:
		img := img.(*image.Gray16)
		for y := 0; y < img.Bounds().Dx(); y++ {
			for x := 0; x < img.Bounds().Dy(); x++ {
				gravity.Log.Print(img.Gray16At(x, y))
			}
		}
		return nil
	default:
		x := img.(*image.RGBA)
		_ = x
		panic("unknown type")
	}

}
