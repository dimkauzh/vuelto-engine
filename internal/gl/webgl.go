//go:build js || wasm
// +build js wasm

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

package gl

import (
	"log"
	"syscall/js"

	"vuelto.pp.ua/internal/font"
	"vuelto.pp.ua/internal/gl/webgl"
	"vuelto.pp.ua/internal/image"
	"vuelto.pp.ua/internal/trita"
)

type Arguments struct {
	Arg *js.Value
}

type Shader struct {
	WebShader     string
	DesktopShader string

	Type   js.Value
	Shader js.Value
}

type Program struct {
	Program js.Value

	VertexShader   js.Value
	FragmentShader js.Value
}

type Buffer struct {
	Vao js.Value
	Vbo js.Value
	Ebo js.Value

	Vertices []float32
	Indices  []uint16
}

type Location struct {
	UniformLocation js.Value
}

type Texture struct {
	Texture js.Value
}

var TEXTURE_2D = &Arguments{&webgl.TEXTURE_2D}
var BLEND = &Arguments{&webgl.BLEND}
var LINEAR = &Arguments{&webgl.LINEAR}
var NEAREST = &Arguments{&webgl.NEAREST}
var VBO = &Arguments{&webgl.ARRAY_BUFFER}
var EBO = &Arguments{&webgl.ELEMENT_ARRAY_BUFFER}
var VA = &Arguments{&webgl.VERTEX_ARRAY}

func NewShader(shadertype any) *Shader {
	shader := &Shader{}

	switch trita.YourType(shadertype) {
	case trita.YourType(FragmentShader{}):
		shader.Type = webgl.FRAGMENT_SHADER
		shader.DesktopShader = shadertype.(FragmentShader).DesktopShader
		shader.WebShader = shadertype.(FragmentShader).WebShader
	case trita.YourType(VertexShader{}):
		shader.Type = webgl.VERTEX_SHADER
		shader.DesktopShader = shadertype.(VertexShader).DesktopShader
		shader.WebShader = shadertype.(VertexShader).WebShader
	default:
		panic("Unknown shader type")
	}

	shader.Shader = webgl.CreateShader(shader.Type)
	webgl.ShaderSource(shader.Shader, shader.WebShader)

	return shader
}

func (s *Shader) Compile() {
	webgl.CompileShader(s.Shader)
}

func (s *Shader) Delete() {
	webgl.DeleteShader(s.Shader)
}

func NewProgram(vertexShader, fragmentShader Shader) *Program {
	return &Program{
		VertexShader:   vertexShader.Shader,
		FragmentShader: fragmentShader.Shader,
		Program:        webgl.CreateProgram(),
	}
}

func (p *Program) Link() {
	webgl.AttachShader(p.Program, p.VertexShader)
	webgl.AttachShader(p.Program, p.FragmentShader)
	webgl.LinkProgram(p.Program)
}

func (p *Program) Use() {
	webgl.UseProgram(p.Program)
}

func (p *Program) UnUse() {
	webgl.UseProgram(js.Null())
}

func (p *Program) Delete() {
	webgl.DeleteProgram(p.Program)
}

func (p *Program) UniformLocation(location string) *Location {
	return &Location{
		UniformLocation: webgl.GetUniformLocation(p.Program, location),
	}
}

func (l *Location) Set(arg ...float32) {
	switch len(arg) {
	case 1:
		webgl.Uniform1f(l.UniformLocation, arg[0])
	case 2:
		webgl.Uniform2f(l.UniformLocation, arg[0], arg[1])
	case 3:
		webgl.Uniform3f(l.UniformLocation, arg[0], arg[1], arg[2])
	case 4:
		webgl.Uniform4f(l.UniformLocation, arg[0], arg[1], arg[2], arg[3])
	default:
		panic("Unsupported uniform length")
	}
}

func GenBuffers(vertices []float32, indices []uint16) *Buffer {
	vao := webgl.CreateVertexArray()
	vbo := webgl.CreateBuffer()
	ebo := webgl.CreateBuffer()

	webgl.BindBuffer(webgl.ARRAY_BUFFER, vbo)
	webgl.BufferData(webgl.ARRAY_BUFFER, webgl.NewFloat32Array(vertices), webgl.STATIC_DRAW)

	webgl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, ebo)
	webgl.BufferData(webgl.ELEMENT_ARRAY_BUFFER, webgl.NewUint16Array(indices), webgl.STATIC_DRAW)

	return &Buffer{
		Vao:      vao,
		Vbo:      vbo,
		Ebo:      ebo,
		Vertices: vertices,
		Indices:  indices,
	}
}

func (b *Buffer) Bind(args ...*Arguments) {
	for _, arg := range args {
		switch arg {
		case VA:
			webgl.BindVertexArray(b.Vao)
		case VBO:
			webgl.BindBuffer(webgl.ARRAY_BUFFER, b.Vbo)
		case EBO:
			webgl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, b.Ebo)
		default:
			log.Fatalln("Unknown argument: ", arg)
		}
	}
}

func (b *Buffer) UnBind(args ...*Arguments) {
	for _, arg := range args {
		switch arg {
		case VA:
			webgl.BindVertexArray(js.Null())
		case VBO:
			webgl.BindBuffer(webgl.ARRAY_BUFFER, js.Null())
		case EBO:
			webgl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, js.Null())
		default:
			log.Fatalln("Unknown argument: ", arg)
		}
	}
}

func (b *Buffer) Data() {
	webgl.BufferData(webgl.ARRAY_BUFFER, webgl.NewFloat32Array(b.Vertices), webgl.STATIC_DRAW)
	webgl.BufferData(webgl.ELEMENT_ARRAY_BUFFER, webgl.NewUint16Array(b.Indices), webgl.STATIC_DRAW)
}

func (b *Buffer) Update(data []float32) {
	webgl.BufferData(webgl.ARRAY_BUFFER, webgl.NewFloat32Array(data), webgl.DYNAMIC_DRAW)
}

func (b *Buffer) Delete(args ...*Arguments) {
	for _, arg := range args {
		switch arg {
		case VA:
			webgl.DeleteVertexArray(b.Vao)
		case VBO:
			webgl.DeleteBuffer(b.Vbo)
		case EBO:
			webgl.DeleteBuffer(b.Ebo)
		default:
			log.Fatalln("Unknown argument: ", arg)
		}
	}
}

func GenTexture() *Texture {
	return &Texture{Texture: webgl.CreateTexture()}
}

func (t *Texture) Bind() {
	webgl.ActiveTexture(webgl.TEXTURE0)
	webgl.BindTexture(webgl.TEXTURE_2D, t.Texture)
}

func (t *Texture) UnBind() {
	webgl.BindTexture(webgl.TEXTURE_2D, js.Null())
}

func (t *Texture) Configure(inputImage any, filter *Arguments) {
	switch trita.YourType(inputImage) {
	case trita.YourType(&image.Image{}):
		outputImage := inputImage.(*image.Image)
		webgl.TexImage2D(webgl.TEXTURE_2D, 0, webgl.RGBA, outputImage.Width, outputImage.Height, 0, webgl.RGBA, webgl.UNSIGNED_BYTE, outputImage.Texture)
	case trita.YourType(&font.Font{}):
		outputImage := inputImage.(*font.Font)
		webgl.TexImage2D(webgl.TEXTURE_2D, 0, webgl.RGBA, outputImage.Widthbound, outputImage.Heightbound, 0, webgl.RGBA, webgl.UNSIGNED_BYTE, outputImage.Texture)
	default:
		panic("Unknown texture type")
	}

	webgl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_MIN_FILTER, *filter.Arg)
	webgl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_MAG_FILTER, *filter.Arg)
	webgl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_WRAP_S, webgl.CLAMP_TO_EDGE)
	webgl.TexParameteri(webgl.TEXTURE_2D, webgl.TEXTURE_WRAP_T, webgl.CLAMP_TO_EDGE)
}

func (t *Texture) Delete() {
	webgl.DeleteTexture(t.Texture)
}

func SetupVertexAttrib(program *Program) {
	useTexture := webgl.GetUniform(program.Program, webgl.GetUniformLocation(program.Program, "useTexture")).Bool()
	if useTexture == true {
		webgl.EnableVertexAttribArray(1)
		webgl.VertexAttribPointer(0, 3, webgl.FLOAT, false, 5*4, 0)
		webgl.EnableVertexAttribArray(0)

		webgl.VertexAttribPointer(1, 2, webgl.FLOAT, false, 5*4, 3*4)
		webgl.EnableVertexAttribArray(0)
	} else {
		webgl.VertexAttribPointer(0, 3, webgl.FLOAT, false, 3*4, 0)
		webgl.EnableVertexAttribArray(0)
	}
}

func DrawElements(indices []uint16) {
	if len(indices) == 2 {
		webgl.DrawElements(webgl.LINES, len(indices), webgl.UNSIGNED_SHORT, 0)
	} else {
		webgl.DrawElements(webgl.TRIANGLES, len(indices), webgl.UNSIGNED_SHORT, 0)
	}
}

func DrawArrays(verticesCount int) {
	webgl.DrawArrays(webgl.TRIANGLE_FAN, 0, verticesCount)
}

func Clear() {
	webgl.Clear(webgl.COLOR_BUFFER_BIT)
	webgl.Clear(webgl.DEPTH_BUFFER_BIT)
}

func ClearColor(r, g, b, a float32) {
	webgl.ClearColor(r, g, b, a)
}

func Enable(args ...*Arguments) {
	for _, capability := range args {
		webgl.Enable(*capability.Arg)
	}
}

func EnableBlend() {
	webgl.BlendFunc(webgl.SRC_ALPHA, webgl.ONE_MINUS_SRC_ALPHA)
}

func Viewport(x, y, width, height int) {
	webgl.Viewport(x, y, width, height)
}

func Init() error {
	webgl.InitWebGL()
	return nil
}
