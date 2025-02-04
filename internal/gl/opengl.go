//go:build windows || linux || darwin
// +build windows linux darwin

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
	"fmt"
	"log"
	"strings"

	"vuelto.pp.ua/internal/font"
	gl "vuelto.pp.ua/internal/gl/opengl"
	"vuelto.pp.ua/internal/image"
	"vuelto.pp.ua/internal/trita"
)

type Arguments struct {
	Arg uint
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
var BLEND = &Arguments{gl.BLEND}
var LINEAR = &Arguments{gl.LINEAR}
var NEAREST = &Arguments{gl.NEAREST}
var VBO = &Arguments{gl.ARRAY_BUFFER}
var EBO = &Arguments{gl.ELEMENT_ARRAY_BUFFER}
var VA = &Arguments{gl.VERTEX_ARRAY}

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

	var success int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		panic(fmt.Sprintf("Failed to compile shader: %v", log))
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

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		logg := make([]byte, logLength)
		gl.GetProgramInfoLog(program, logLength, nil, &logg[0])
		log.Fatalf("Program linking failed: %s", string(logg))
	}

	p.Program = program
}

func (p *Program) Use() {
	gl.UseProgram(p.Program)
}

func (p *Program) UnUse() {
	gl.UseProgram(0)
}

func (p *Program) Delete() {
	gl.DeleteProgram(p.Program)
}

func (p *Program) UniformLocation(location string) *Location {
	loc := gl.GetUniformLocation(p.Program, gl.Str(location+"\x00"))
	if loc == -1 {
		log.Fatalln("Uniform not found: ", location)
		return nil
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

func (b *Buffer) Bind(args ...*Arguments) {
	for _, arg := range args {
		switch arg {
		case VA:
			gl.BindVertexArray(b.Vao)
		case VBO:
			gl.BindBuffer(gl.ARRAY_BUFFER, b.Vbo)
		case EBO:
			gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, b.Ebo)
		default:
			fmt.Printf("Unknown argument: %v\n", arg)
		}
	}
}

func (b *Buffer) UnBind(args ...*Arguments) {
	for _, arg := range args {
		switch arg {
		case VA:
			gl.BindVertexArray(0)
		case VBO:
			gl.BindBuffer(gl.ARRAY_BUFFER, 0)
		case EBO:
			gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
		default:
			fmt.Printf("Unknown argument: %v\n", arg)
		}
	}
}

func (b *Buffer) Data() {
	gl.BufferData(gl.ARRAY_BUFFER, len(b.Vertices)*4, gl.Ptr(b.Vertices), gl.STATIC_DRAW)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(b.Indices)*4, gl.Ptr(b.Indices), gl.STATIC_DRAW)
}

func (b *Buffer) Update(data []float32) {
	gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, gl.Ptr(data), gl.DYNAMIC_DRAW)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(b.Indices)*4, gl.Ptr(b.Indices), gl.DYNAMIC_DRAW)
}

func (b *Buffer) Delete(args ...*Arguments) {
	for _, arg := range args {
		switch arg {
		case VA:
			gl.DeleteVertexArrays(1, &b.Vao)
		case VBO:
			gl.DeleteBuffers(1, &b.Vbo)
		case EBO:
			gl.DeleteBuffers(1, &b.Ebo)
		default:
			fmt.Printf("Unknown argument: %v\n", arg)
		}
	}
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

func (t *Texture) Configure(inputImage any, filter *Arguments) {
	switch trita.YourType(inputImage) {
	case trita.YourType(&image.Image{}):
		outputImage := inputImage.(*image.Image)
		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(outputImage.Width), int32(outputImage.Height), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(outputImage.Texture))
	case trita.YourType(&font.Font{}):
		outputImage := inputImage.(*font.Font)
		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(outputImage.Widthbound), int32(outputImage.Heightbound), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(outputImage.Texture))
	default:
		panic("Unknown texture type")
	}

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, int32(filter.Arg))
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, int32(filter.Arg))
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
	if len(indices) == 2 {
		gl.DrawElements(gl.LINES, int32(len(indices)), gl.UNSIGNED_INT, nil)
	} else {
		gl.DrawElements(gl.TRIANGLES, int32(len(indices)), gl.UNSIGNED_INT, nil)
	}
}

func DrawArrays(corners int32) {
	gl.DrawArrays(gl.TRIANGLE_FAN, 0, corners)
}

func Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func ClearColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
}

func Enable(args ...*Arguments) {
	for _, arg := range args {
		gl.Enable(uint32(arg.Arg))
	}
}

func EnableBlend() {
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
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
