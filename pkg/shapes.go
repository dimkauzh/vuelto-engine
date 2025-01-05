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
	"vuelto.pp.ua/internal/gl"
	"vuelto.pp.ua/internal/gl/ushaders"
)

type Line struct {
	Renderer *Renderer2D
	Pos1     *Vector2D
	Pos2     *Vector2D
	Color    [4]int

	Buffer  *gl.Buffer
	Program *gl.Program
	Indices []uint16
}

type Rect struct {
	Renderer *Renderer2D
	Pos      *Vector2D
	Width    float32
	Height   float32
	Color    [4]int

	Buffer  *gl.Buffer
	Program *gl.Program
	Indices []uint16
}

// Loads a new line and returns a Line struct. Can be later drawn using Draw() method
func (r *Renderer2D) NewLine(x1, y1, x2, y2 float32, color [4]int) *Line {
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
		x1, y1, 0.0,
		x2, y2, 0.0,
	}

	program.UniformLocation("uniformColor").Set(
		float32(color[0])/255,
		float32(color[1])/255,
		float32(color[2])/255,
		float32(color[3])/255,
	)

	program.UniformLocation("useTexture").Set(0)

	indices := []uint16{
		0, 1,
	}

	buffer := gl.GenBuffers(vertices, indices)
	buffer.Bind(gl.VA, gl.VBO, gl.EBO)

	buffer.Data()
	gl.SetupVertexAttrib(program)

	return &Line{
		Renderer: r,
		Pos1:     NewVector2D(x1, y1),
		Pos2:     NewVector2D(x2, y2),
		Color:    color,

		Buffer:  buffer,
		Program: program,
		Indices: indices,
	}
}

// Loads a new rect and returns a Rect struct. Can be later drawn using Draw() method
func (r *Renderer2D) NewRect(x, y, width, height float32, color [4]int) *Rect {
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
		x, y, 0.0,
		x, y - height, 0.0,
		x + width, y - height, 0.0,
		x + width, y, 0.0,
	}
	program.UniformLocation("useTexture").Set(0)
	program.UniformLocation("uniformColor").Set(
		float32(color[0])/255,
		float32(color[1])/255,
		float32(color[2])/255,
		float32(color[3])/255,
	)

	indices := []uint16{
		0, 1, 3,
		1, 2, 3,
	}

	buffer := gl.GenBuffers(vertices, indices)
	buffer.Bind(gl.VA, gl.VBO, gl.EBO)

	buffer.Data()
	gl.SetupVertexAttrib(program)

	return &Rect{
		Renderer: r,
		Width:    width,
		Height:   height,
		Pos:      NewVector2D(x, y),
		Color:    color,

		Buffer:  buffer,
		Program: program,
		Indices: indices,
	}
}

// Draws the line loaded previously
func (l *Line) Draw() {
	vertices := []float32{
		l.Pos1.X, l.Pos1.Y, 0.0,
		l.Pos2.X, l.Pos2.Y, 0.0,
	}

	l.Program.Use()

	l.Buffer.Bind(gl.VA, gl.VBO, gl.EBO)
	l.Buffer.Update(vertices)
	gl.DrawElements(l.Indices)
	l.Buffer.UnBind(gl.VA, gl.VBO, gl.EBO)
}

// Draws the rect loaded previously
func (r *Rect) Draw() {

	vertices := []float32{
		r.Pos.X, r.Pos.Y, 0.0,
		r.Pos.X, r.Pos.Y - r.Height, 0.0,
		r.Pos.X + r.Width, r.Pos.Y - r.Height, 0.0,
		r.Pos.X + r.Width, r.Pos.Y, 0.0,
	}

	r.Program.Use()

	r.Buffer.Bind(gl.VA, gl.VBO, gl.EBO)
	r.Buffer.Update(vertices)
	gl.DrawElements(r.Indices)
	r.Buffer.UnBind(gl.VA, gl.VBO, gl.EBO)
}
