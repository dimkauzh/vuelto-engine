package main

import (
	"embed"
	_ "embed"
	"log"

	"vuelto.pp.ua/internal/gl"
	"vuelto.pp.ua/internal/gl/ushaders"
	"vuelto.pp.ua/internal/image"

	windowing "vuelto.pp.ua/internal/window"
)

//go:embed tree.png
var embeddedFiles embed.FS

func framebuffersizecallback(window *windowing.Window, newWidth, newHeight int) {
	gl.Viewport(0, 0, newWidth, newHeight)
}

func main() {
	win, err := windowing.InitWindow()
	if err != nil {
		log.Fatalf("Failed to initialise: %s", err)
	}
	defer win.Close()

	win.Resizable = true
	win.Title = "Test"

	win.Width = 500
	win.Height = 500

	win.GlfwGLMajor = 3
	win.GlfwGLMinor = 3

	err = win.Create()
	if err != nil {
		log.Fatalln("Error create window:", err)
	}

	win.ResizingCallback(framebuffersizecallback)

	err = gl.Init()
	if err != nil {
		log.Fatalf("Failed to initialise: %s", err)
	}

	gl.Enable(gl.TEXTURE_2D)

	win.ContextCurrent()

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
	defer program.Delete()

	program.Use()

	vertices := []float32{
		// Positions      // Texture Coords
		-0.5, 0.5, 0.0, 0.0, 0.0, // Top-left
		-0.5, -0.5, 0.0, 0.0, 1.0, // Bottom-left
		0.5, -0.5, 0.0, 1.0, 1.0, // Bottom-right
		0.5, 0.5, 0.0, 1.0, 0.0, // Top-right
	}

	program.UniformLocation("uniformColor").Set(0, 0, 0, 1.0)
	program.UniformLocation("useTexture").Set(1)

	indices := []uint16{
		0, 1, 3, // bottom-left, bottom-right, top-right
		1, 2, 3, // bottom-left, top-right, top-left
	}

	texture := gl.GenTexture()
	texture.Bind()
	texture.Configure(image.LoadAsEmbed(embeddedFiles, "tree.png"), gl.NEAREST)
	texture.UnBind()
	defer texture.Delete()

	buffer := gl.GenBuffers(vertices, indices)
	buffer.BindVA()
	buffer.BindVBO()
	buffer.BindEBO()
	defer buffer.Delete()

	buffer.Data()
	gl.SetupVertexAttrib(program)

	for !win.Close() {
		gl.Clear()

		texture.Bind()
		gl.DrawElements(indices)
		texture.UnBind()

		win.HandleEvents()
		win.UpdateBuffers()
	}

}
