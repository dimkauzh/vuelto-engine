//go:build windows || linux || darwin
// +build windows linux darwin

/*
 * Copyright (C) 2024 vuelto-org
 *
 * This file is part of the Vuelto project, licensed under the Vuelto License V1.
 * Primary License: GNU GPLv3 or later (see <https://www.gnu.org/licenses/>).
 * If unmaintained, this software defaults to the MIT License as per Vuelto License V1,
 * at which point the copyright no longer applies.
 *
 * Distributed WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 */

package gl

type VertexShader struct{}
type FragmentShader struct{}

type Shader struct {
	WebShader     string
	DesktopShader string

	Type any
}

type Program struct {
	Program uint32

	VertexShader   Shader
	FragmentShader Shader
}

type Buffer struct {
	Vao uint32
	Vbo uint32

	Vertices []float32
}

type Location struct {
	UniformLocation int32
}

type Texture struct {
	Texture uint32
}

var VERTEX_SHADER = &VertexShader{}
var FRAGMENT_SHADER = &FragmentShader{}

func NewShader(shadertype any, webshader, desktopshader string) *Shader {
	return &Shader{
		Type: shadertype,

		WebShader:     webshader,
		DesktopShader: desktopshader,
	}
}

func (s *Shader) Compile() {}

func (s *Shader) Delete() {}

func NewProgram(vertexshader, fragmentshader Shader) *Program {
	return &Program{
		VertexShader:   vertexshader,
		FragmentShader: fragmentshader,
	}
}

func (p *Program) Link() {}

func (p *Program) Use() {}

func (p *Program) Delete() {}

func (p *Program) UniformLocation(location string) *Location {
	return &Location{}
}

func (l *Location) Set(arg ...float32) {}

func GenBuffers(vertices []float32) *Buffer {
	return &Buffer{
		Vertices: vertices,
	}
}

func (b *Buffer) BindVA() {}

func (b *Buffer) BindVBO() {}

func (b *Buffer) UnBindVA() {}

func (b *Buffer) UnBindVBO() {}

func (b *Buffer) Data() {}

func (b *Buffer) Delete() {}

func GenTexture() *Texture {
	return &Texture{}
}

func (t *Texture) Bind() {}

func (t *Texture) UnBind() {}

func InitVertexAttrib() {}

func DrawElements(corners int) {}

func Clear() {}

func Enable(args ...uint32) {}

func Viewport(width, height int32) {}

func Ortho() {}

func Texture2D() {}
