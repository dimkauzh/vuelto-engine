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

package vuelto

import (
	"embed"

	"vuelto.pp.ua/internal/font"
	"vuelto.pp.ua/internal/gl"
	"vuelto.pp.ua/internal/gl/ushaders"
	"vuelto.pp.ua/internal/trita"
)

type Font struct {
	X, Y          float32
	Width, Height float32

	Buffer  *gl.Buffer
	Texture *gl.Texture
	Indices []uint16
	Program *gl.Program

	Renderer *UIRenderer
}

type FontEmbed struct {
	Filesystem embed.FS
	Font       string
}

type FontHTTP struct {
	Url string
}

// Loads a new image and returns a Image struct. Can be later drawn using the Draw() method
func (r *UIRenderer) LoadFont(fontFile any, text string, x, y float32, size int) *Font {
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

	program.UniformLocation("uniformColor").Set(1, 1, 1, 1)
	program.UniformLocation("useTexture").Set(1)

	indices := []uint16{
		0, 1, 3,
		1, 2, 3,
	}

	var file *font.Font
	switch trita.YourType(fontFile) {
	case trita.YourType(""):
		file = font.Load(r.Window.Width, r.Window.Height, fontFile.(string), text, size, x, y)
	case trita.YourType(FontEmbed{}):
		embed := fontFile.(FontEmbed)
		file = font.LoadAsEmbed(r.Window.Width, r.Window.Height, embed.Filesystem, embed.Font, text, size, x, y)
	case trita.YourType(FontHTTP{}):
		http := fontFile.(FontHTTP)
		file = font.LoadAsHTTP(r.Window.Width, r.Window.Height, http.Url, text, size, x, y)
	}

	texture := gl.GenTexture()
	texture.Bind()
	texture.Configure(file, gl.NEAREST)
	texture.UnBind()

	vertices := []float32{
		x, y, 0.0, 0.0, 1.0,
		x, y - file.Height, 0.0, 0.0, 0.0,
		x + file.Width, y - file.Height, 0.0, 1.0, 0.0,
		x + file.Width, y, 0.0, 1.0, 1.0,
	}

	buffer := gl.GenBuffers(vertices, indices)
	buffer.Bind(gl.VA, gl.VBO, gl.EBO)

	buffer.Data()
	gl.SetupVertexAttrib(program)

	r.Window.UnsetCurrent()

	return &Font{
		X: x,
		Y: y,

		Width:  file.Width,
		Height: file.Height,

		Buffer:  buffer,
		Texture: texture,
		Indices: indices,
		Program: program,

		Renderer: r,
	}
}

// Draws the image that's loaded before.
func (f *Font) Draw() {
	f.Renderer.Window.SetCurrent()

	vertices := []float32{
		f.X, f.Y, 0.0, 0.0, 1.0,
		f.X, f.Y - f.Height, 0.0, 0.0, 0.0,
		f.X + f.Width, f.Y - f.Height, 0.0, 1.0, 0.0,
		f.X + f.Width, f.Y, 0.0, 1.0, 1.0,
	}

	f.Program.Use()
	f.Buffer.Bind(gl.VA, gl.VBO, gl.EBO)
	f.Buffer.Update(vertices)

	f.Texture.Bind()
	gl.DrawElements(f.Indices)
	f.Texture.UnBind()

	f.Buffer.UnBind(gl.VA, gl.VBO, gl.EBO)
	f.Program.UnUse()

	f.Renderer.Window.UnsetCurrent()
}
