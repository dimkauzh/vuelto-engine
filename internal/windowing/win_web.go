//go:build js && wasm
// +build js,wasm

package windowing

import (
	"syscall/js"
	"time"

	"vuelto.me/internal/windowing/web"
)

type Window struct {
	Resizable bool

	JSCanvas web.Canvas

  GlfwGLMajor int
	GlfwGLMinor int

	Title  string
	Width  int
	Height int
}

func InitWindow() (*Window, error) {
	w := &Window{}

	w.JSCanvas = web.Document.GetElementById("vuelto")
	if w.JSCanvas.IsNull() {
		w.JSCanvas = web.Document.CreateCanvasElement()
		w.JSCanvas.Set("id", "canvas")
		web.Document.Body.AppendCanvasChild(w.JSCanvas)
	}

	return w, nil
}

func (w *Window) CreateWindow() error {
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

	return nil
}

func (w *Window) ResizingCallback(inputFunc func(*Window, int, int)) {
	web.Window.AddEventListener("resize", func(this js.Value, p []js.Value) any {
		inputFunc(w, web.Document.DocumentElement.ClientWidth(), web.Document.DocumentElement.ClientHeight())
		return nil
	})
}

func (w *Window) UpdateBuffers() {
	time.Sleep(time.Millisecond * 32)
}

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

func (w *Window) ContextCurrent() {}
func (w *Window) Destroy() {}
func (w *Window) HandleEvents() {}
func (w *Window) Close() {}
