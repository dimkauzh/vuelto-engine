//go:build js && wasm
// +build js,wasm

/*
 * Copyright (C) 2024 vuelto-org
 *
 * This file is part of the Vuelto project, licensed under the VL-Cv1.1 License.
 * Primary License: GNU GPLv3 or later (see <https://www.gnu.org/licenses/>).
 * If unmaintained, this software defaults to the MIT License as per Vuelto License V1.1,
 * at which point the copyright no longer applies.
 *
 * Distributed WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 */

package web

import (
	"log"
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
	context := c.JSCanvas.Call("getContext", inputContext)
	if context.IsNull() {
		log.Fatalln("failed to get context: ", inputContext)
		return Context{}
	}
	return Context{context}
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
