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

import "syscall/js"

var Console JSConsole

type JSConsole struct{}

func (c *JSConsole) Log(message ...any) {
	js.Global().Get("console").Call("log", message...)
}

func (c *JSConsole) Warn(message ...any) {
	js.Global().Get("console").Call("warn", message...)
}

func (c *JSConsole) Error(message ...any) {
	js.Global().Get("console").Call("error", message...)
}
