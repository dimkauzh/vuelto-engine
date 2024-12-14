/*
 * Copyright (C) 2024 vuelto-org
 *
 * This file is part of the Vuelto project, licensed under the VL-Cv1.1 License.
 * Primary License: GNU GPLv3 or later (see <https://www.gnu.org/licenses/>).
 * If unmaintained, this software defaults to the MIT License as per Vuelto License V1,
 * at which point the copyright no longer applies.
 *
 * Distributed WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 */

package vuelto

import (
	"log"

	gl "vuelto.pp.ua/internal/gl/legacy"
	windowing "vuelto.pp.ua/internal/window"
)

type Window struct {
	Window        *windowing.Window
	Title         string
	Width, Height int
}

func framebuffersizecallback(window *windowing.Window, newWidth, newHeight int) {
	gl.Viewport(0, 0, newWidth, newHeight)
}

// Creates a new window and returns a Window struct.
func NewWindow(title string, width, height int, resizable bool) *Window {
	window, err := windowing.InitWindow()
	if err != nil {
		log.Fatalln("Could not initialise a new window: ", err)
		return nil
	}

	window.GlfwGLMajor = 2
	window.GlfwGLMinor = 1

	window.Title = title
	window.Width = width
	window.Height = height

	window.Resizable = resizable

	err = window.Create()
	if err != nil {
		log.Fatalln("Error create window:", err)
	}

	window.ResizingCallback(framebuffersizecallback)

	window.ContextCurrent()

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
	w.Window.SetResizable(resizable)
}

// Function created for a loop. Returns true when being closed, and returns false when being active.
func (w *Window) Close() bool {
	for !w.Window.Close() {
		w.Window.HandleEvents()
		return false
	}
	cleanTex()
	return true
}

// Refreshes te window. Run this at the end of your loop (except if you're having multiple windows)
func (w *Window) Refresh() {
	w.Window.UpdateBuffers()
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

// Sets the context of the window to the current context. (Only use when having multiple windows)
func (w *Window) SetContextCurrent() {
	w.Window.ContextCurrent()
}

// Destroys the window and cleans up the memory.
func (w *Window) Destroy() {
	w.Window.Destroy()
	cleanTex()
}
