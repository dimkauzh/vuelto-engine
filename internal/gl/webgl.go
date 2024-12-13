//go:build js && wasm
// +build js,wasm

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
	"syscall/js"

	"vuelto.pp.ua/internal/gl/webgl"
	"vuelto.pp.ua/internal/trita"
)

type EnableArg struct {
	Arg js.Value
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
	Vao        uint32
	Vbo        js.Value
	Vertices   []float32
	BufferType js.Value
}

type Location struct {
	UniformLocation js.Value
}

var TEXTURE_2D = &EnableArg{webgl.TEXTURE_2D}

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

func NewProgram(vertexshader, fragmentshader Shader) *Program {
	return &Program{
		VertexShader:   vertexshader.Shader,
		FragmentShader: fragmentshader.Shader,
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

func GenBuffers(vertices []float32) *Buffer {
	buffer := webgl.CreateBuffer()
	webgl.BindBuffer(webgl.ARRAY_BUFFER, buffer)
	webgl.BufferData(webgl.ARRAY_BUFFER, vertices, webgl.STATIC_DRAW)
	return &Buffer{
		Vertices: vertices,
		Vbo:      buffer,
	}
}

func (b *Buffer) BindVA() {}

func (b *Buffer) BindVBO() {
	webgl.BindBuffer(webgl.ARRAY_BUFFER, b.Vbo)
}

func (b *Buffer) Data() {
	webgl.BufferData(webgl.ARRAY_BUFFER, b.Vertices, webgl.STATIC_DRAW)
}

func (b *Buffer) Delete() {
	webgl.DeleteBuffer(b.Vbo)
}

func InitVertexAttrib() {
	webgl.VertexAttribPointer(0, 3, webgl.FLOAT, false, 0, 0)
	webgl.EnableVertexAttribArray(0)
}

func DrawElements(corners int) {}

func DrawArrays(verticesCount int) {
	webgl.DrawArrays(webgl.TRIANGLE_FAN, 0, verticesCount)
}

func Clear() {
	webgl.Clear(webgl.COLOR_BUFFER_BIT)
	webgl.Clear(webgl.DEPTH_BUFFER_BIT)
}

func Enable(args ...EnableArg) {
	for _, capability := range args {
		webgl.Enable(capability.Arg)
	}
}

func Init() error {
	webgl.InitWebGL()
	return nil
}

func Viewport(x, y, width, height int) {
	webgl.Viewport(x, y, width, height)
}
