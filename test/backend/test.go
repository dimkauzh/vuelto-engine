package main

import (
	"embed"
	_ "embed"
	"log"

	"vuelto.pp.ua/internal/event"
	"vuelto.pp.ua/internal/gl"
	"vuelto.pp.ua/internal/gl/ushaders"
	"vuelto.pp.ua/internal/image"

	windowing "vuelto.pp.ua/internal/window"
)

//go:embed tree.png
var embeddedFiles embed.FS

func frameBufferSizeCallback(window *windowing.Window, newWidth, newHeight int) {
	gl.Viewport(0, 0, newWidth, newHeight)
}

func main() {
	win, err := windowing.InitWindow()
	if err != nil {
		log.Fatalf("Failed to initialize: %s", err)
	}
	defer win.Close()

	win.Resizable = false
	win.Title = "Test"

	win.Width = 500
	win.Height = 500

	win.GlfwGLMajor = 3
	win.GlfwGLMinor = 3
	win.SetFPS(30)

	err = win.Create()
	if err != nil {
		log.Fatalln("Error create window:", err)
	}

	win.ResizingCallback(frameBufferSizeCallback)

	events := event.Init(win)

	err = gl.Init()
	if err != nil {
		log.Fatalf("Failed to initialize: %s", err)
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
	texture.Configure(image.LoadAsHTTP("https://dev-tester.com/content/images/2021/12/blog_cover_further_api_testing_with_http_toolkit.png"), gl.NEAREST)
	texture.UnBind()
	defer texture.Delete()

	buffer := gl.GenBuffers(vertices, indices)
	buffer.Bind(gl.VA, gl.VBO, gl.EBO)
	defer buffer.Delete(gl.VA, gl.VBO, gl.EBO)

	buffer.Data()
	gl.SetupVertexAttrib(program)

	dx := float32(0.01)
	dy := float32(0.0)

	for !win.Close() {
		gl.Clear()
		gl.ClearColor(0.2, 0.2, 0.2, 1)

		if events.Key(event.KeyMap["Left"]) == event.PRESSED {
			dx = -0.01
		} else if events.Key(event.KeyMap["Right"]) == event.PRESSED {
			dx = 0.01
		} else {
			dx = 0.0
		}

		for i := 0; i < len(vertices); i += 5 {
			vertices[i] += dx
			vertices[i+1] += dy
		}

		buffer.Bind(gl.VBO)
		buffer.Update(vertices)

		texture.Bind()
		gl.DrawElements(indices)
		texture.UnBind()

		win.HandleEvents()
		win.UpdateBuffers()
	}
}
