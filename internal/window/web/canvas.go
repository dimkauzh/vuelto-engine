package web

import (
	"syscall/js"

	"strconv"
)

type Canvas struct {
	JSCanvas js.Value
}

type Context struct {
	JSContext js.Value
}

func (c *Canvas) GetContext(inputContext string) Context {
	return Context{c.JSCanvas.Call("getContext", inputContext)}
}

func (c *Canvas) Get(inputKey string) js.Value {
	return c.JSCanvas.Get(inputKey)
}

func (c *Canvas) Set(inputKey string, inputValue any) {
	c.JSCanvas.Set(inputKey, inputValue)
}

func (c *Canvas) IsNull() bool {
	return c.JSCanvas.IsNull()
}

func (c *Canvas) AddEventListener(inputEvent string, inputFunc func(this js.Value, args []js.Value) any) {
	js.Global().Get("window").Call("addEventListener", inputEvent, js.FuncOf(inputFunc))
}

func (c *Context) Set(inputKey string, inputValue any) {
	c.JSContext.Set(inputKey, inputValue)
}

func (c *Context) FillText(inputText string, x, y float64) {
	c.JSContext.Call("fillText", inputText, x, y)
}

func (c *Context) FillRect(x, y, width, height float64) {
	c.JSContext.Call("fillRect", x, y, width, height)
}

func (c *Context) ClearRect(x, y, width, height float64) {
	c.JSContext.Call("clearRect", x, y, width, height)
}

func (c *Context) DrawImage(inputImage js.Value, x, y, width, height float64) {
	c.JSContext.Call("drawImage", inputImage, x, y, width, height)
}

func (c *Context) Call(inputKey string, inputArgs ...any) js.Value {
	return c.JSContext.Call(inputKey, inputArgs)
}

func (c *Context) Get(inputKey string) js.Value {
	return c.JSContext.Get(inputKey)
}

func GetRGBA(color [4]int) string {
	return "rgba(" +
		strconv.Itoa(color[0]) + "," +
		strconv.Itoa(color[1]) + "," +
		strconv.Itoa(color[2]) + "," +
		strconv.FormatFloat(float64(color[3])/255.0, 'f', -1, 64) + ")"
}
