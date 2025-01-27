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

package vuelto

import (
	"embed"

	"vuelto.pp.ua/internal/gl"
	"vuelto.pp.ua/internal/gl/ushaders"
	"vuelto.pp.ua/internal/image"
	"vuelto.pp.ua/internal/trita"
)

type Image struct {
	Pos           *Vector2D
	Width, Height float32

	Buffer  *gl.Buffer
	Texture *gl.Texture
	Indices []uint16
	Program *gl.Program

	Renderer *Renderer2D
}

type ImageEmbed struct {
	Filesystem embed.FS
	Image      string
}

type ImageHTTP struct {
	Url string
}

var ImageArray []uint32

// Loads a new image and returns a Image struct. Can be later drawn using the Draw() method
func (r *Renderer2D) LoadImage(imageFile any, x, y, width, height float32) *Image {
	r.Window.SetCurrent()

	vertexShader := gl.NewShader(gl.VertexShader{
		WebShader:     ushaders.WebVShader,
		DesktopShader: ushaders.DesktopVShader,
	})
	fragmentShader := gl.NewShader(gl.FragmentShader{
		WebShader:     ushaders.WebFShader,
		DesktopShader: ushaders.DesktopFShader,
	})

	vertexShader.Compile()
	defer vertexShader.Delete()

	fragmentShader.Compile()
	defer fragmentShader.Delete()

	program := gl.NewProgram(*vertexShader, *fragmentShader)
	program.Link()

	program.Use()

	vertices := []float32{
		x, y, 0.0, 0.0, 0.0,
		x, y - height, 0.0, 0.0, 1.0,
		x + width, y - height, 0.0, 1.0, 1.0,
		x + width, y, 0.0, 1.0, 0.0,
	}

	program.UniformLocation("uniformColor").Set(0, 0, 0, 1.0)
	program.UniformLocation("useTexture").Set(1)

	indices := []uint16{
		0, 1, 3,
		1, 2, 3,
	}

	var file *image.Image
	switch trita.YourType(imageFile) {
	case trita.YourType(""):
		file = image.Load(imageFile.(string))
	case trita.YourType(ImageEmbed{}):
		embed := imageFile.(ImageEmbed)
		file = image.LoadAsEmbed(embed.Filesystem, embed.Image)
	case trita.YourType(ImageHTTP{}):
		file = image.LoadAsHTTP(imageFile.(ImageHTTP).Url)
	}

	texture := gl.GenTexture()
	texture.Bind()
	texture.Configure(file, gl.NEAREST)
	texture.UnBind()

	buffer := gl.GenBuffers(vertices, indices)
	buffer.Bind(gl.VA, gl.VBO, gl.EBO)

	buffer.Data()
	gl.SetupVertexAttrib(program)

	r.Window.UnsetCurrent()

	return &Image{
		Pos:    NewVector2D(x, y),
		Width:  width,
		Height: height,

		Buffer:  buffer,
		Texture: texture,
		Indices: indices,
		Program: program,

		Renderer: r,
	}
}

// Draws the image that's loaded before.
func (img *Image) Draw() {
	img.Renderer.Window.SetCurrent()

	vertices := []float32{
		img.Pos.X, img.Pos.Y, 0.0, 0.0, 0.0,
		img.Pos.X, img.Pos.Y - img.Height, 0.0, 0.0, 1.0,
		img.Pos.X + img.Width, img.Pos.Y - img.Height, 0.0, 1.0, 1.0,
		img.Pos.X + img.Width, img.Pos.Y, 0.0, 1.0, 0.0,
	}

	img.Program.Use()
	img.Buffer.Bind(gl.VA, gl.VBO, gl.EBO)
	img.Buffer.Update(vertices)

	img.Texture.Bind()
	gl.DrawElements(img.Indices)
	img.Texture.UnBind()

	img.Buffer.UnBind(gl.VA, gl.VBO, gl.EBO)
	img.Program.UnUse()

	img.Renderer.Window.UnsetCurrent()
}
