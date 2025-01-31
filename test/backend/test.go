package main

import (
	"embed"
	_ "embed"
	"log"

	"vuelto.pp.ua/internal/event"
	"vuelto.pp.ua/internal/font"
	"vuelto.pp.ua/internal/gl"
	"vuelto.pp.ua/internal/gl/ushaders"

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
	win.ContextCurrent()

	events := event.Init(win)

	err = gl.Init()
	if err != nil {
		log.Fatalf("Failed to initialize: %s", err)
	}

	gl.Enable(gl.TEXTURE_2D, gl.BLEND)
	gl.EnableBlend()

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

	program.UniformLocation("uniformColor").Set(1, 1, 1, 1)
	program.UniformLocation("useTexture").Set(1)

	indices := []uint16{
		0, 1, 3, // bottom-left, bottom-right, top-right
		1, 2, 3, // bottom-left, top-right, top-left
	}

	mytext := font.LoadAsHTTP(500, 500, "https://github.com/fusionengine-org/fusion/raw/refs/heads/main/src/fusionengine/external/font.ttf", "TESTESTEST", 30, 0, 0)

	vertices := []float32{
		0, 0, 0.0, 0.0, 1.0,
		0, 0 - mytext.Height, 0.0, 0.0, 0.0,
		0 + mytext.Width, 0 - mytext.Height, 0.0, 1.0, 0.0,
		0 + mytext.Width, 0, 0.0, 1.0, 1.0,
	}

	texture := gl.GenTexture()
	texture.Bind()
	texture.Configure(mytext, gl.NEAREST)
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

		for i := 0; i < len(mytext.Vertices); i += 5 {
			mytext.Vertices[i] += dx
			mytext.Vertices[i+1] += dy
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
