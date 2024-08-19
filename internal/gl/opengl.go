//go:build windows || linux || darwin
// +build windows linux darwin

package gl

import gl "vuelto.me/internal/gl/opengl"

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

func (b *Buffer) Data() {}

func InitVertexAttrib() {}

func DrawElements() {}
