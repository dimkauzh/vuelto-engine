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
	windowing "vuelto.pp.ua/internal/window"
)

type Event struct {
	Window *windowing.Window
}

type State struct {
	State glfw.Action
}

type Key struct {
	Key glfw.Key
}

type Setting struct {
	Mode  glfw.InputMode
	True  int
	False int
}

var (
	STICKY_KEYS = Setting{
		Mode:  glfw.StickyKeysMode,
		True:  glfw.True,
		False: glfw.False,
	}
	DISABLE_CURSOR = Setting{
		Mode:  glfw.CursorMode,
		True:  glfw.CursorDisabled,
		False: glfw.CursorNormal,
	}

	PRESSED  = State{glfw.Press}
	RELEASED = State{glfw.Release}
)

func Init(window *windowing.Window) *Event {
	return &Event{
		Window: window,
	}
}

func (e *Event) SetSetting(setting Setting, value bool) {
	if value {
		e.Window.GlfwWindow.SetInputMode(setting.Mode, setting.True)
	} else {
		e.Window.GlfwWindow.SetInputMode(setting.Mode, setting.False)
	}
}

func (e *Event) Key(key Key) State {
	return State{e.Window.GlfwWindow.GetKey(key.Key)}
}

func (e *Event) MousePos() (float64, float64) {
	return e.Window.GlfwWindow.GetCursorPos()
}
