//go:build windows || linux || darwin
// +build windows linux darwin

/*
 * Copyright (C) 2024 vuelto-org
 *
 * This file is part of the Vuelto project, licensed under the Vuelto License V1.
 * Primary License: GNU GPLv3 or later (see <https://www.gnu.org/licenses/>).
 * If unmaintained, this software defaults to the MIT License as per Vuelto License V1,
 * at which point the copyright no longer applies.
 *
 * Distributed WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 */

package windowing

import (
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type Window struct {
	Resizable bool

	GlfwGLMajor int
	GlfwGLMinor int
	GlfwWindow  *glfw.Window

	Title  string
	Width  int
	Height int
}

func InitWindow() (*Window, error) {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		return nil, err
	}
	return &Window{}, nil
}

func (w *Window) Create() error {
	if w.Resizable {
		glfw.WindowHint(glfw.Resizable, glfw.True)
	} else {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}

	if w.GlfwGLMajor != 0 {
		glfw.WindowHint(glfw.ContextVersionMajor, w.GlfwGLMajor)
	}
	if w.GlfwGLMinor != 0 {
		glfw.WindowHint(glfw.ContextVersionMinor, w.GlfwGLMinor)
	}

	if w.GlfwGLMajor >= 3 {
		glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

		if runtime.GOOS == "darwin" {
			glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
		}
	}

	var err error
	w.GlfwWindow, err = glfw.CreateWindow(w.Width, w.Height, w.Title, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (w *Window) ResizingCallback(inputFunc func(*Window, int, int)) {
	w.GlfwWindow.SetFramebufferSizeCallback(func(win *glfw.Window, newWidth, newHeight int) {
		inputFunc(w, newWidth, newHeight)
	})
}

func (w *Window) SetResizable(resizable bool) {
	if resizable {
		w.GlfwWindow.SetAttrib(glfw.Resizable, glfw.True)
	} else {
		w.GlfwWindow.SetAttrib(glfw.Resizable, glfw.False)
	}
}

func (w *Window) Close() bool {
	for !w.GlfwWindow.ShouldClose() {
		return false
	}
	return true
}

func (w *Window) HandleEvents() {
	glfw.PollEvents()
}

func (w *Window) UpdateBuffers() {
	w.GlfwWindow.SwapBuffers()
}

func (w *Window) ContextCurrent() {
	w.GlfwWindow.MakeContextCurrent()
}

func (w *Window) Destroy() {
	w.GlfwWindow.Destroy()
}

func (w *Window) DrawingTest() {
}
