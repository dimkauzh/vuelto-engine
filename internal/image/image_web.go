//go:build js && wasm
// +build js,wasm

/*
 * Copyright (C) 2024 vuelto-org
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
	_ "image/jpeg"
	_ "image/png"
	"log"
	"syscall/js"
)

type Image struct {
	Path    string
	Texture js.Value
	Width   int
	Height  int
}

func Load(imagePath string) *Image {
	panic("Load() is not supported in web assembly")
}

func LoadAsEmbed(fs embed.FS, imagePath string) *Image {
	imgFile, err := fs.ReadFile(imagePath)
	if err != nil {
		log.Fatalf("failed to read embedded image '%s': %v", imagePath, err)
	}

	img, _, err := image.Decode(bytes.NewReader(imgFile))
	if err != nil {
		log.Fatalf("failed to decode image '%s': %v", imagePath, err)
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

	data := js.Global().Get("Uint8Array").New(len(rgbaImg.Pix))
	for i, v := range rgbaImg.Pix {
		data.SetIndex(i, v)
	}

	return &Image{
		Path:    imagePath,
		Texture: data,
		Width:   rgbaImg.Bounds().Dx(),
		Height:  rgbaImg.Bounds().Dy(),
	}
}
