package main

import (
	_ "embed"
	"log"

	"vuelto.pp.ua/internal/gl"
	"vuelto.pp.ua/internal/gl/ushaders"

	windowing "vuelto.pp.ua/internal/window"
)

func framebuffersizecallback(window *windowing.Window, newWidth, newHeight int) {
	gl.Viewport(0, 0, newWidth, newHeight)
}

func main() {
	win, err := windowing.InitWindow()
	if err != nil {
		log.Fatalf("Failed to initialise: %s", err)
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
		-0.9, -0.9, 0.0,
		0.9, -0.9, 0.0,
	}

	program.UniformLocation("uniformColor").Set(0.2, 0.6, 0.3, 1.0)
	program.UniformLocation("useTexture").Set(0)

	indices := []uint16{
		0, 1,
	}

	buffer := gl.GenBuffers(vertices, indices)
	buffer.Bind(gl.VA, gl.VBO, gl.EBO)
	defer buffer.Delete(gl.VA, gl.VBO, gl.EBO)

	buffer.Data()
	gl.SetupVertexAttrib(program)

	for !win.Close() {
		gl.Clear()
		gl.ClearColor(0.2, 0.2, 0.2, 1)

		buffer.Bind(gl.VA, gl.VBO, gl.EBO)
		gl.DrawElements(indices)
		buffer.UnBind(gl.VA, gl.VBO, gl.EBO)

		win.HandleEvents()
		win.UpdateBuffers()
	}
}
