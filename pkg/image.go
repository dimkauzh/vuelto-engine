/*
 * Copyright (C) 2024 vuelto-org
 *
 * This file is part of the Vuelto project, licensed under the VL-Cv1.1 License.
 * Primary License: GNU GPLv3 or later (see <https://www.gnu.org/licenses/>).
 * If unmaintained, this software defaults to the MIT License as per Vuelto License V1,
 * at which point the copyright no longer applies.
 *
 * Distributed WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 */

package vuelto

import (
	gl "vuelto.pp.ua/internal/gl/legacy"
	"vuelto.pp.ua/internal/image"
)

type Image struct {
	Texture       uint32
	X, Y          float32
	Width, Height float32
}

var ImageArray []uint32

// Loads a new image and returns a Image struct. Can be later drawn using the Draw() method
func (r *Renderer2D) LoadImage(imagePath string, x, y, width, height float32) *Image {
	file := image.Load(imagePath)

	var textureID uint32
	gl.GenTextures(1, &textureID)

	gl.BindTexture(gl.TEXTURE_2D, textureID)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, file.Width, file.Height, 0, gl.RGBA, gl.UNSIGNED_BYTE, file.Texture)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	ImageArray = append(ImageArray, textureID)

	return &Image{
		Texture: textureID,
		X:       x,
		Y:       y,
		Width:   width,
		Height:  height,
	}
}

// Draws the image that's loaded before.
func (img *Image) Draw() {
	gl.BindTexture(gl.TEXTURE_2D, img.Texture)
	defer gl.BindTexture(gl.TEXTURE_2D, 0)

	gl.Begin(gl.QUADS)
	defer gl.End()

	gl.TexCoord2f(0.0, 0.0)
	gl.Vertex2f(img.X, img.Y)
	gl.TexCoord2f(1.0, 0.0)
	gl.Vertex2f(img.X+img.Width, img.Y)
	gl.TexCoord2f(1.0, 1.0)
	gl.Vertex2f(img.X+img.Width, img.Y+img.Height)
	gl.TexCoord2f(0.0, 1.0)
	gl.Vertex2f(img.X, img.Y+img.Height)
}

func cleanTex() {
	for _, i := range ImageArray {
		gl.DeleteTextures(1, &i)
	}
}
