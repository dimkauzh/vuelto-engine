package image

import (
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

type Image struct {
	Texture []uint8

	ImageWidth  int
	ImageHeight int
}

func Load(imagePath string) *Image {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatalln("Failed to open image: ", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln("Failed to decode image: ", err)
	}

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Over)

	return &Image{
		Texture:     rgba.Pix,
		ImageWidth:  rgba.Rect.Size().X,
		ImageHeight: rgba.Rect.Size().Y,
	}
}
