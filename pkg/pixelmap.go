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
	"vuelto.pp.ua/internal/gl"
	"vuelto.pp.ua/internal/gl/ushaders"
	"vuelto.pp.ua/internal/image"
)

type Pixelmap struct {
	Texture  map[int]map[int][4]float32
	Renderer *Renderer2D

	Width  int
	Height int
}

// Loads a new pixelmap and returns a Pixelmap struct. Can be later drawn using Draw() method
func (r *Renderer2D) NewPixelmap() *Pixelmap {
	return &Pixelmap{
		Renderer: r,
		Width:    r.Window.Width,
		Height:   r.Window.Height,
	}
}

// SetPixel assigns an RGBA color to the (x, y) coordinate.
func (p *Pixelmap) SetPixel(x, y int, color []float32) {
	if x < 0 || x >= p.Width || y < 0 || y >= p.Height || len(color) < 4 {
		return
	}
	if p.Texture[x] == nil {
		p.Texture[x] = make(map[int][4]float32)
	}
	p.Texture[x][y] = [4]float32{color[0], color[1], color[2], color[3]}
}

// Draw renders the pixelmap on the screen.
func (p *Pixelmap) Draw() {
	p.Renderer.Window.SetCurrent()

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
		-1, -1, 0.0, 0.0, 0.0,
		-1, -1 - 1, 0.0, 0.0, 1.0,
		-1 + 1, -1 - 1, 0.0, 1.0, 1.0,
		-1 + 1, -1, 0.0, 1.0, 0.0,
	}

	program.UniformLocation("uniformColor").Set(0, 0, 0, 1.0)
	program.UniformLocation("useTexture").Set(1)

	indices := []uint16{
		0, 1, 3,
		1, 2, 3,
	}

	texture := gl.GenTexture()
	texture.Bind()
	texture.Configure(image.LoadPixelmap(p.Texture, p.Renderer.Window.Width, p.Renderer.Window.Height), gl.NEAREST)
	texture.UnBind()

	buffer := gl.GenBuffers(vertices, indices)
	buffer.Bind(gl.VA, gl.VBO, gl.EBO)

	buffer.Data()
	gl.SetupVertexAttrib(program)

	program.Use()
	buffer.Bind(gl.VA, gl.VBO, gl.EBO)
	buffer.Update(vertices)

	texture.Bind()
	gl.DrawElements(indices)
	texture.UnBind()

	buffer.UnBind(gl.VA, gl.VBO, gl.EBO)
	program.UnUse()

	p.Renderer.Window.UnsetCurrent()
}
