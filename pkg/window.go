/*
 * Copyright (C) 2025 vuelto-org
 *
 * This file is part of the Vuelto project, licensed under the VL-Cv1.1 License.
 * Primary License: GNU GPLv3 or later (see <https://www.gnu.org/licenses/>).
 * If unmaintained, this software defaults to the MIT License as per Vuelto License V1.1,
 * at which point the copyright no longer applies.
 *
 * Distributed WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 */

package vuelto

import (
	"log"

	"vuelto.pp.ua/internal/event"
	"vuelto.pp.ua/internal/gl"
	windowing "vuelto.pp.ua/internal/window"
)

type Window struct {
	Window        *windowing.Window
	Title         string
	Width, Height int

	Event *event.Event
}

func frameBufferSizeCallback(window *windowing.Window, newWidth, newHeight int) {
	gl.Viewport(0, 0, newWidth, newHeight)
}

// Creates a new window and returns a Window struct.
func NewWindow(title string, width, height int, resizable bool, transparent bool) *Window {
	window, err := windowing.InitWindow()
	if err != nil {
		log.Fatalln("Could not initialize a new window: ", err)
		return nil
	}
	defer window.Close()

	window.GlfwGLMajor = 3
	window.GlfwGLMinor = 3

	window.Title = title
	window.Width = width
	window.Height = height

	window.Resizable = resizable
	window.Transparency = transparent

	err = window.Create()
	if err != nil {
		log.Fatalln("Error create window:", err)
	}

	window.ResizingCallback(frameBufferSizeCallback)

	window.ContextCurrent()

	events := event.Init(window)

	err = gl.Init()
	if err != nil {
		log.Fatalf("Failed to initialize: %s", err)
	}

	gl.Enable(gl.TEXTURE_2D, gl.BLEND)
	gl.EnableBlend()

	return &Window{
		Window: window,
		Title:  title,
		Width:  width,
		Height: height,
		Event:  events,
	}
}

// Set callback to the resize of the window
func (w *Window) SetResizeCallback(callback func(window *Window, newWidth, newHeight int)) {
	w.Window.ResizingCallback(func(window *windowing.Window, newWidth, newHeight int) {
		gl.Viewport(0, 0, newWidth, newHeight)
		callback(w, newWidth, newHeight)
	})
}

// Sets the resizable attribute of the window.
func (w *Window) SetResizable(resizable bool) {
	w.Window.SetResizable(resizable)
}

// Sets the title of the window.
func (w *Window) SetTitle(title string) {
	w.Window.SetTitle(title)
}

// Sets the size of the window.
func (w *Window) SetSize(width, height int) {
	w.Window.SetSize(width, height)
}

// Returns the size of the window.
func (w *Window) GetSize() (int, int) {
	return w.Window.GetSize()
}

// Function created for a loop. Returns true when being closed, and returns false when being active.
func (w *Window) Close() bool {
	for !w.Window.Close() {
		return false
	}
	return true
}

// Refreshes te window. Run this at the end of your loop (except if you're having multiple windows)
func (w *Window) Refresh() {
	w.SetCurrent()
	w.Window.HandleEvents()
	w.Window.UpdateBuffers()
	gl.Clear()
	w.UnsetCurrent()
}

// Sets the context of the window to the current context. (Only use when having multiple windows)
func (w *Window) SetCurrent() {
	w.Window.ContextCurrent()
}

// Unsets the context of the window. (Only use when having multiple windows)
func (w *Window) UnsetCurrent() {
	w.Window.UnsetContext()
}

// Destroys the window and cleans up the memory.
func (w *Window) Destroy() {
	w.Window.Destroy()
}

func (w *Window) GetDeltaTime() float32 {
	return float32(w.Window.GetDeltaTime())
}

func (w *Window) SetFPS(fps int) {
	w.Window.SetFPS(fps)
}

func (w *Window) GetFPS() int {
	return w.Window.GetFPS()
}
