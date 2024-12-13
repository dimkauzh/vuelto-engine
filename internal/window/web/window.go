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

package web

import "syscall/js"

var Window JSWindow

type JSWindow struct{}

var funcMaps = make(map[string]func(js.Value, []js.Value) any)

func (w *JSWindow) AddEventListener(inputEvent string, inputFunc func(js.Value, []js.Value) any) {
	js.Global().Get("window").Call("addEventListener", inputEvent, js.FuncOf(inputFunc))
	funcMaps[inputEvent] = inputFunc
}

func (w *JSWindow) RemoveEventListener(inputEvent string) {
	js.Global().Get("window").Call("removeEventListener", inputEvent, js.FuncOf(funcMaps[inputEvent]))
	delete(funcMaps, inputEvent)
}
