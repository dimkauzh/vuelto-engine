//go:build windows || linux || darwin
// +build windows linux darwin

/*
 * Copyright (C) 2025 vuelto-org
 *
 * This file is part of the Vuelto project, licensed under the VL-Cv1.1 License.
 * Primary License: GNU GPLv3 or later (see <https://www.gnu.org/licenses/>).
 * If unmaintained, this software defaults to the MIT License as per Vuelto License V1.1,
 * at which point the copyright no longer applies.
 *
 * Distributed WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 */

package image

import (
	"bytes"
	"embed"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"os"
)

type Image struct {
	Path    string
	Texture []uint8

	Width  int
	Height int
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
		Texture: rgba.Pix,
		Width:   rgba.Rect.Size().X,
		Height:  rgba.Rect.Size().Y,
	}
}

func LoadAsEmbed(fs embed.FS, imagePath string) *Image {
	imgFile, err := fs.ReadFile(imagePath)
	if err != nil {
		log.Fatalf("failed to read embedded image '%s': %v", imagePath, err)
	}

	return loadImage(imgFile, imagePath)
}

func LoadAsHTTP(imageUrl string) *Image {
	if !(len(imageUrl) > 4 && (imageUrl[:7] == "http://" || imageUrl[:8] == "https://")) {
		panic("Load() only supports HTTP and HTTPS paths in web assembly")
	}

	resp, err := http.Get(imageUrl)
	if err != nil {
		log.Fatalf("failed to fetch image '%s': %v", imageUrl, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("failed to fetch image '%s': status code %d", imageUrl, resp.StatusCode)
	}

	imgData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read image data from '%s': %v", imageUrl, err)
	}

	return loadImage(imgData, imageUrl)
}

func loadImage(imgData []byte, imageUrl string) *Image {
	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		log.Fatalf("failed to decode image '%s': %v", imageUrl, err)
	}

	rgbaImg, ok := img.(*image.RGBA)
	if !ok {
		bounds := img.Bounds()
		rgbaImg = image.NewRGBA(bounds)
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				rgbaImg.Set(x, y, img.At(x, y))
			}
		}
	}

	return &Image{
		Path:    imageUrl,
		Texture: rgbaImg.Pix,
		Width:   rgbaImg.Bounds().Dx(),
		Height:  rgbaImg.Bounds().Dy(),
	}
}

func LoadPixelmap(pixels map[int]map[int][4]float32, width, height int) *Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for x, col := range pixels {
		for y, c := range col {
			r := uint8(c[0] * 255)
			g := uint8(c[1] * 255)
			b := uint8(c[2] * 255)
			a := uint8(c[3] * 255)
			img.Set(x, y, color.RGBA{r, g, b, a})
		}
	}

	return &Image{
		Texture: img.Pix,
		Width:   img.Rect.Size().X,
		Height:  img.Rect.Size().Y,
	}
}
