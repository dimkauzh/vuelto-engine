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

package web

import "syscall/js"

var Document JSDocument

type JSDocument struct {
	Body            Body
	DocumentElement DocumentElement
}

type Body struct {
	Style StyleBody
}

type Element struct {
	JSElement js.Value
	Style     js.Value
}

type DocumentElement struct {
	Style StyleDocument
}

type StyleDocument struct{}
type StyleBody struct{}

func (d *JSDocument) CreateElement(newElement string) Element {
	docElement := js.Global().Get("document").Call("createElement", newElement)
	style := docElement.Get("style")

	return Element{
		JSElement: docElement,
		Style:     style,
	}
}

func (d *JSDocument) CreateCanvasElement() Canvas {
	newElement := js.Global().Get("document").Call("createElement", "canvas")

	return Canvas{
		JSCanvas: newElement,
	}
}

func (d *JSDocument) AddEventListener(inputEvent string, inputFunc func(js.Value, []js.Value) any) {
	js.Global().Get("document").Call("addEventListener", inputEvent, js.FuncOf(inputFunc))
}

func (e *Element) Set(keyValue string, inputValue any) {
	e.JSElement.Set(keyValue, inputValue)
}

func (e *Element) AddEventListener(inputEvent string, inputFunc func(js.Value, []js.Value) any) {
	e.JSElement.Call("addEventListener", inputEvent, js.FuncOf(inputFunc))
}

func (b *Body) AppendChild(inputElement Element) {
	appendChild := inputElement.JSElement
	js.Global().Get("document").Get("body").Call("appendChild", appendChild)
}

func (b *Body) AppendCanvasChild(inputCanvas Canvas) {
	appendChild := inputCanvas.JSCanvas
	js.Global().Get("document").Get("body").Call("appendChild", appendChild)
}

func (d *JSDocument) GetElementById(inputID string) Canvas {
	return Canvas{
		JSCanvas: js.Global().Get("document").Call("getElementById", inputID),
	}
}

func (d *DocumentElement) ClientWidth() int {
	return js.Global().Get("document").Get("documentElement").Get("clientWidth").Int()
}

func (d *DocumentElement) ClientHeight() int {
	return js.Global().Get("document").Get("documentElement").Get("clientHeight").Int()
}

func (d *StyleDocument) Set(inputKey string, inputValue any) {
	js.Global().Get("document").Get("documentElement").Get("style").Set(inputKey, inputValue)
}

func (b *StyleBody) Set(inputKey string, inputValue any) {
	js.Global().Get("document").Get("body").Get("style").Set(inputKey, inputValue)
}
