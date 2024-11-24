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

import (
	"fmt"

	gl "vuelto.me/internal/gl/opengl"
)

type VertexShader struct{}
type FragmentShader struct{}

type EnableArg struct {
	Arg uint32
}

type Shader struct {
	WebShader     string
	DesktopShader string

	Type   any
	Shader uint32
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

var TEXTURE_2D = &EnableArg{gl.TEXTURE_2D}

func NewShader(shadertype any, webshader, desktopshader string) *Shader {
	return &Shader{
		Type:          shadertype,
		WebShader:     webshader,
		DesktopShader: desktopshader,
	}
}

func (s *Shader) Compile() {
	var shaderType uint32
	switch s.Type.(type) {
	case *VertexShader:
		shaderType = gl.VERTEX_SHADER
	case *FragmentShader:
		shaderType = gl.FRAGMENT_SHADER
	default:
		panic("invalid shader type")
	}

	shader := gl.CreateShader(shaderType)
	src, free := gl.Strs(s.DesktopShader + "\x00")
	gl.ShaderSource(shader, 1, src, nil)
	free()
	gl.CompileShader(shader)

	var success int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		log := make([]byte, logLength+1)
		gl.GetShaderInfoLog(shader, logLength, nil, &log[0])
		panic(fmt.Sprintf("failed to compile shader: %s", log))
	}

	s.Type = shader
}

func (s *Shader) Delete() {
	if shader, ok := s.Type.(uint32); ok {
		gl.DeleteShader(shader)
	}
}

func NewProgram(vertexshader, fragmentshader Shader) *Program {
	return &Program{
		VertexShader:   vertexshader,
		FragmentShader: fragmentshader,
	}
}

func (p *Program) Link() {
	program := gl.CreateProgram()

	vertexShader, ok := p.VertexShader.Type.(uint32)
	if !ok {
		panic("vertex shader is not compiled")
	}
	fragmentShader, ok := p.FragmentShader.Type.(uint32)
	if !ok {
		panic("fragment shader is not compiled")
	}

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var success int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)
		log := make([]byte, logLength+1)
		gl.GetProgramInfoLog(program, logLength, nil, &log[0])
		panic(fmt.Sprintf("failed to link program: %s", log))
	}

	p.Program = program
}

func (p *Program) Use() {
	gl.UseProgram(p.Program)
}

func (p *Program) Delete() {
	gl.DeleteProgram(p.Program)
}

func (p *Program) UniformLocation(location string) *Location {
	loc := gl.GetUniformLocation(p.Program, gl.Str(location+"\x00"))
	if loc == -1 {
		panic(fmt.Sprintf("uniform %s not found", location))
	}
	return &Location{UniformLocation: loc}
}

func (l *Location) Set(arg ...float32) {
	switch len(arg) {
	case 1:
		gl.Uniform1f(l.UniformLocation, arg[0])
	case 2:
		gl.Uniform2f(l.UniformLocation, arg[0], arg[1])
	case 3:
		gl.Uniform3f(l.UniformLocation, arg[0], arg[1], arg[2])
	case 4:
		gl.Uniform4f(l.UniformLocation, arg[0], arg[1], arg[2], arg[3])
	default:
		panic("unsupported uniform length")
	}
}

func GenBuffers(vertices []float32) *Buffer {
	var vao, vbo uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)

	return &Buffer{
		Vao:      vao,
		Vbo:      vbo,
		Vertices: vertices,
	}
}

func (b *Buffer) BindVA() {
	gl.BindVertexArray(b.Vao)
}

func (b *Buffer) BindVBO() {
	gl.BindBuffer(gl.ARRAY_BUFFER, b.Vbo)
}

func (b *Buffer) UnBindVA() {
	gl.BindVertexArray(0)
}

func (b *Buffer) UnBindVBO() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func (b *Buffer) Data() {
	gl.BufferData(gl.ARRAY_BUFFER, len(b.Vertices)*4, gl.Ptr(b.Vertices), gl.STATIC_DRAW)
}

func (b *Buffer) Delete() {
	gl.DeleteVertexArrays(1, &b.Vao)
	gl.DeleteBuffers(1, &b.Vbo)
}

func GenTexture() *Texture {
	var texture uint32
	gl.GenTextures(1, &texture)
	return &Texture{Texture: texture}
}

func (t *Texture) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, t.Texture)
}

func (t *Texture) UnBind() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func InitVertexAttrib() {
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)
}

func DrawElements(corners int) {
	gl.DrawElements(gl.TRIANGLES, int32(corners), gl.UNSIGNED_INT, gl.PtrOffset(0))
}

func Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func Enable(args ...EnableArg) {
	for _, arg := range args {
		gl.Enable(arg.Arg)
	}
}

func Viewport(width, height int32) {
	gl.Viewport(0, 0, width, height)
}

func Ortho() {
}

func Init() error {
	if err := gl.Init(); err != nil {
		return fmt.Errorf("failed to initialize OpenGL: %w", err)
	}
	return nil
}
