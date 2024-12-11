package main

import (
	"log"

	"vuelto.pp.ua/internal/gl"
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

	win.ContextCurrent()

	vertexShader := gl.NewShader(gl.VertexShader{
		WebShader:     WebVertexShaderSource,
		DesktopShader: VertexShaderSource,
	})
	fragmentShader := gl.NewShader(gl.FragmentShader{
		WebShader:     WebFragmentShaderSource,
		DesktopShader: FragmentShaderSource,
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
		0.5, 0.5, 0.0, // bottom-left
		0.5, -0.5, 0.0, // bottom-right
		-0.5, -0.5, 0.0, // top-right
		-0.5, 0.5, 0.0, // top-left
	}

	indices := []uint32{
		0, 1, 3, // bottom-left, bottom-right, top-right
		1, 2, 3, // bottom-left, top-right, top-left
	}

	buffer := gl.GenBuffers(vertices, indices)
	buffer.BindVA()
	buffer.BindVBO()
	buffer.BindEBO()
	defer buffer.Delete()

	buffer.Data()
	gl.InitVertexAttrib()

	for !win.Close() {
		gl.Clear()

		gl.DrawElements(indices)

		win.HandleEvents()
		win.UpdateBuffers()
	}

}
