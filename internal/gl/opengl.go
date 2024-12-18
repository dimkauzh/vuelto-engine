//go:build windows || linux || darwin
// +build windows linux darwin

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

package gl

import (
	"fmt"

	gl "vuelto.pp.ua/internal/gl/opengl"
	"vuelto.pp.ua/internal/image"
	"vuelto.pp.ua/internal/trita"
)

type Arguments struct {
	Arg any
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
	Ebo uint32

	Vertices []float32
	Indices  []int32
}

type Location struct {
	UniformLocation int32
}

type Texture struct {
	Texture uint32
}

var TEXTURE_2D = &Arguments{gl.TEXTURE_2D}
var LINEAR = &Arguments{gl.LINEAR}
var NEAREST = &Arguments{gl.NEAREST}

func NewShader(shadertype any) *Shader {
	switch trita.YourType(shadertype) {
	case trita.YourType(FragmentShader{}):
		return &Shader{
			Type:          shadertype,
			WebShader:     shadertype.(FragmentShader).WebShader,
			DesktopShader: shadertype.(FragmentShader).DesktopShader,
		}
	case trita.YourType(VertexShader{}):
		return &Shader{
			Type:          shadertype,
			WebShader:     shadertype.(VertexShader).WebShader,
			DesktopShader: shadertype.(VertexShader).DesktopShader,
		}
	default:
		panic("Unknown shader type")
	}
}

func (s *Shader) Compile() {
	var shaderType uint32

	switch trita.YourType(s.Type) {
	case trita.YourType(VertexShader{}):
		shaderType = gl.VERTEX_SHADER
	case trita.YourType(FragmentShader{}):
		shaderType = gl.FRAGMENT_SHADER
	default:
		panic("Invalid shader type")
	}

	shader := gl.CreateShader(shaderType)
	src, free := gl.Strs(s.DesktopShader + "\x00")
	gl.ShaderSource(shader, 1, src, nil)
	free()
	gl.CompileShader(shader)

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

func GenBuffers(vertices []float32, indices []uint16) *Buffer {
	var vao, vbo, ebo uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)
	gl.GenBuffers(1, &ebo)

	return &Buffer{
		Vao:      vao,
		Vbo:      vbo,
		Ebo:      ebo,
		Vertices: vertices,
		Indices:  gl.Uint16ToInt32(indices),
	}
}

func (b *Buffer) BindVA() {
	gl.BindVertexArray(b.Vao)
}

func (b *Buffer) BindVBO() {
	gl.BindBuffer(gl.ARRAY_BUFFER, b.Vbo)
}

func (b *Buffer) BindEBO() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, b.Ebo)
}

func (b *Buffer) UnBindVA() {
	gl.BindVertexArray(0)
}

func (b *Buffer) UnBindVBO() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func (b *Buffer) UnBindEBO() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}

func (b *Buffer) Data() {
	gl.BufferData(gl.ARRAY_BUFFER, len(b.Vertices)*4, gl.Ptr(b.Vertices), gl.STATIC_DRAW)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(b.Indices)*4, gl.Ptr(b.Indices), gl.STATIC_DRAW)
}

func (b *Buffer) Delete() {
	gl.DeleteVertexArrays(1, &b.Vao)
	gl.DeleteBuffers(1, &b.Vbo)
	gl.DeleteBuffers(1, &b.Ebo)
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

func (t *Texture) Configure(image *image.Image, filter *Arguments) {
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(image.Width), int32(image.Height), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(image.Texture))

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, int32(filter.Arg.(int)))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, int32(filter.Arg.(int)))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
}

func (t *Texture) Delete() {
	gl.DeleteTextures(1, &t.Texture)
}

func SetupVertexAttrib(program *Program) {
	var useTexture int32
	gl.GetUniformiv(program.Program, gl.GetUniformLocation(program.Program, gl.Str("useTexture"+"\x00")), &useTexture)
	if useTexture == 1 {
		gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 5*4, 0)
		gl.EnableVertexAttribArray(0)

		gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 5*4, 3*4)
		gl.EnableVertexAttribArray(1)
	} else {
		gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 3*4, 0)
		gl.EnableVertexAttribArray(0)
	}
}

func DrawElements(indices []uint16) {
	gl.DrawElements(gl.TRIANGLES, int32(len(indices)), gl.UNSIGNED_INT, nil)
}

func DrawArrays(corners int32) {
	gl.DrawArrays(gl.TRIANGLE_FAN, 0, corners)
}

func Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func Enable(args ...*Arguments) {
	for _, arg := range args {
		gl.Enable(uint32(arg.Arg.(int)))
	}
}

func Viewport(x, y, width, height int) {
	gl.Viewport(int32(x), int32(y), int32(width), int32(height))
}

func Init() error {
	if err := gl.Init(); err != nil {
		return fmt.Errorf("failed to initialize OpenGL: %w", err)
	}
	return nil
}
