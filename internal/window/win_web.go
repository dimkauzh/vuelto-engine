//go:build js || wasm
// +build js wasm

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

package windowing

import (
	"fmt"
	"os"
	"syscall/js"
	"time"

	"vuelto.pp.ua/internal/gl"
	"vuelto.pp.ua/internal/window/web"
)

type Window struct {
	Resizable bool

	JSCanvas web.Canvas

	GlfwGLMajor int
	GlfwGLMinor int

	Title        string
	Width        int
	Height       int
	Transparency bool

	lastTime      time.Time
	deltaTime     float32
	desiredFPS    int
	frameDuration time.Duration
}

var initialized bool

func InitWindow() (*Window, error) {
	if initialized && os.Getenv("VUELTO_DISABLE_BUILD_ERRORS") == "" {
		panic("Web doesnt support having multiple windows!")
	} else if !initialized {
		initialized = true
	}

	w := &Window{
		desiredFPS:    60,
		frameDuration: time.Second / 60,
		lastTime:      time.Now(),
	}

	w.JSCanvas = web.Document.GetElementById("vuelto")
	if w.JSCanvas.IsNull() {
		w.JSCanvas = web.Document.CreateCanvasElement()
		w.JSCanvas.Set("id", "vuelto")
		web.Document.Body.AppendCanvasChild(w.JSCanvas)
	}

	if w.JSCanvas.IsNull() {
		return nil, fmt.Errorf("failed to create or fetch canvas")
	}

	return w, nil
}

func (w *Window) Create() error {
	web.Document.DocumentElement.Style.Set("overflow", "hidden")
	web.Document.Body.Style.Set("overflow", "hidden")

	if w.Resizable {
		w.JSCanvas.Set("width", web.Document.DocumentElement.ClientWidth())
		w.JSCanvas.Set("height", web.Document.DocumentElement.ClientHeight())

		web.Window.AddEventListener("resize", func(this js.Value, p []js.Value) any {
			w.JSCanvas.Set("width", web.Document.DocumentElement.ClientWidth())
			w.JSCanvas.Set("height", web.Document.DocumentElement.ClientHeight())

			w.Width = web.Document.DocumentElement.ClientWidth()
			w.Height = web.Document.DocumentElement.ClientHeight()
			return nil
		})
	} else {
		w.JSCanvas.Set("width", w.Width)
		w.JSCanvas.Set("height", w.Height)
	}

	if w.Transparency && os.Getenv("VUELTO_DISABLE_BUILD_ERRORS") == "" {
		panic("Web doesnt support having multiple windows!")
	}

	return nil
}

func (w *Window) ResizingCallback(inputFunc func(*Window, int, int)) {
	web.Window.AddEventListener("resize", func(this js.Value, p []js.Value) any {
		newWidth := web.Document.DocumentElement.ClientWidth()
		newHeight := web.Document.DocumentElement.ClientHeight()

		w.JSCanvas.Set("width", newWidth)
		w.JSCanvas.Set("height", newHeight)

		gl.Viewport(0, 0, newWidth, newHeight)

		inputFunc(w, newWidth, newHeight)
		return nil
	})
}

func (w *Window) UpdateBuffers() {}

func (w *Window) SetResizable(resizable bool) {
	if resizable {
		w.JSCanvas.Set("width", web.Document.DocumentElement.ClientWidth())
		w.JSCanvas.Set("height", web.Document.DocumentElement.ClientHeight())

		web.Window.AddEventListener("resize", func(this js.Value, p []js.Value) any {
			w.JSCanvas.Set("width", web.Document.DocumentElement.ClientWidth())
			w.JSCanvas.Set("height", web.Document.DocumentElement.ClientHeight())
			return nil
		})
	} else {
		web.Window.RemoveEventListener("resize")
		w.JSCanvas.Set("width", w.Width)
		w.JSCanvas.Set("height", w.Height)
	}
}

func (w *Window) SetTransparency(alpha float32) {
	if os.Getenv("VUELTO_DISABLE_BUILD_ERRORS") == "" {
		panic("SetTransparency() is not supported on web!")
	}
}

func (w *Window) SetTitle(title string) {
	web.Document.Set("title", title)
}

func (w *Window) SetSize(width, height int) {
	w.JSCanvas.Set("width", width)
	w.JSCanvas.Set("height", height)
}

func (w *Window) GetSize() (int, int) {
	return w.JSCanvas.Get("width").Int(), w.JSCanvas.Get("height").Int()
}

func (w *Window) Close() bool {
	return false
}

func (w *Window) ContextCurrent() {}

func (w *Window) UnsetContext() {}

func (w *Window) Destroy() {}

func (w *Window) HandleEvents() {
	now := time.Now()
	w.deltaTime = float32(now.Sub(w.lastTime).Seconds())
	w.lastTime = now

	duration := time.Since(w.lastTime)
	if duration < w.frameDuration {
		time.Sleep(w.frameDuration - duration)
	}
}

func (w *Window) GetDeltaTime() float32 {
	return w.deltaTime
}

func (w *Window) SetFPS(fps int) {
	if fps > 0 {
		w.desiredFPS = fps
		w.frameDuration = time.Second / time.Duration(fps)
	}
}

func (w *Window) GetFPS() int {
	return w.desiredFPS
}
