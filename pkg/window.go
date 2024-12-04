package vuelto

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
	"vuelto.pp.ua/internal/gl"
)

type Window struct {
	Window        *glfw.Window
	Title         string
	Width, Height int
}

func framebuffersizecallback(window *glfw.Window, newWidth, newHeight int) {
	gl.Viewport(0, 0, newWidth, newHeight)
}

// Creates a new window and returns a Window struct.
func NewWindow(title string, width, height int, resizable bool) *Window {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	if resizable {
		glfw.WindowHint(glfw.Resizable, glfw.True)
	} else {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		log.Fatalln("Error create window:", err)
	}

	window.SetFramebufferSizeCallback(framebuffersizecallback)

	window.MakeContextCurrent()

	gl.Ortho(0, float64(width), float64(height), 0, -1, 1)

	gl.Enable(gl.BLEND)
	gl.Enable(gl.TEXTURE_2D)

	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	return &Window{
		Window: window,
		Title:  title,
		Width:  width,
		Height: height,
	}
}

// Sets the resizable attribute of the window.
func (w *Window) SetResizable(resizable bool) {
	if resizable {
		w.Window.SetAttrib(glfw.Resizable, glfw.True)
	} else {
		w.Window.SetAttrib(glfw.Resizable, glfw.False)
	}
}

// Function created for a loop. Returns true when being closed, and returns false when being active.
func (w *Window) Close() bool {
	for !w.Window.ShouldClose() {
		glfw.PollEvents()
		return false
	}
	cleanTex()
	return true
}

// Refreshes te window. Run this at the end of your loop (except if you're having multiple windows)
func (w *Window) Refresh() {
	w.Window.SwapBuffers()
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

// Sets the context of the window to the current context. (Only use when having multiple windows)
func (w *Window) SetContextCurrent() {
	w.Window.MakeContextCurrent()
}

// Destroys the window and cleans up the memory.
func (w *Window) Destroy() {
	w.Window.Destroy()
	cleanTex()
}
