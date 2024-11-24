package main

import (
	"log"

	"vuelto.me/internal/gl"
	windowing "vuelto.me/internal/window"
)

func framebuffersizecallback(window *windowing.Window, newWidth, newHeight int) {
	gl.Viewport(int32(newWidth), int32(newHeight))
}

func main() {
	win, err := windowing.InitWindow()
	if err != nil {
		log.Fatalf("Failed to initialise: %s", err)
	}
	defer win.Close()

	err = gl.Init()
	if err != nil {
		log.Fatalf("Failed to initialise: %s", err)
	}

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

	win.ContextCurrent()

	for !win.Close() {
		win.HandleEvents()
		win.UpdateBuffers()
		gl.Clear()
	}
}
