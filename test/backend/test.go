package main

import (
	"log"

	"vuelto.me/internal/gl"
	windowing "vuelto.me/internal/window"
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

	vertexShaderSource := `
		#version 330 core
		layout(location = 0) in vec3 aPos;
		void main() {
			gl_Position = vec4(aPos, 1.0);
		}
	`

	fragmentShaderSource := `
		#version 330 core
		out vec4 FragColor;
		void main() {
			FragColor = vec4(1.0, 0.5, 0.2, 1.0);
		}
	`

	webVertexShaderSource := `#version 300 es
precision mediump float;

layout(location = 0) in vec3 aPos;

void main() {
    gl_Position = vec4(aPos, 1.0);
}`

	webFragmentShaderSource := `#version 300 es
precision mediump float;

out vec4 FragColor;

void main() {
    FragColor = vec4(1.0, 0.5, 0.2, 1.0);
}`

	vertexShader := gl.NewShader(gl.VertexShader{
		WebShader:     webVertexShaderSource,
		DesktopShader: vertexShaderSource,
	})
	fragmentShader := gl.NewShader(gl.FragmentShader{
		WebShader:     webFragmentShaderSource,
		DesktopShader: fragmentShaderSource,
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
		-0.5, -0.5, 0.0, // bottom-left
		0.5, -0.5, 0.0, // bottom-right
		0.5, 0.5, 0.0, // top-right
		-0.5, 0.5, 0.0, // top-left
	}

	buffer := gl.GenBuffers(vertices)
	buffer.BindVA()
	buffer.BindVBO()
	defer buffer.Delete()

	buffer.Data()
	gl.InitVertexAttrib()

	for !win.Close() {
		gl.Clear()

		gl.DrawArrays(4)

		win.HandleEvents()
		win.UpdateBuffers()
	}
}
