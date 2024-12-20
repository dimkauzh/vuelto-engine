//go:build windows || linux || darwin
// +build windows linux darwin

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

package event

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Event struct {
	Window *glfw.Window
}

type State struct {
	State glfw.Action
}

type Key struct {
	Key glfw.Key
}

var PRESSED = State{glfw.Press}
var RELEASED = State{glfw.Release}

func Init(window *glfw.Window) *Event {
	return &Event{
		Window: window,
	}
}

func (e *Event) SetStickyKeys(value bool) {
	glfwValue := glfw.False
	if value {
		glfwValue = glfw.True
	}
	e.Window.SetInputMode(glfw.StickyKeysMode, glfwValue)
}

func (e *Event) DisableCursor(value bool) {
	glfwValue := glfw.CursorDisabled
	if value {
		glfwValue = glfw.CursorNormal
	}
	e.Window.SetInputMode(glfw.CursorMode, glfwValue)
}

func (e *Event) Key(key Key) State {
	return State{e.Window.GetKey(key.Key)}
}

func (e *Event) MousePos() (float64, float64) {
	return e.Window.GetCursorPos()
}
