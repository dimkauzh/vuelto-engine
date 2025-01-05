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

package event

import (
	"syscall/js"

	windowing "vuelto.pp.ua/internal/window"
	"vuelto.pp.ua/internal/window/web"
)

type Event struct {
	Window         *windowing.Window
	IgnoreWarnings bool
}

type State struct {
	State string
}

type Key struct {
	Key string
}

type Setting struct {
	Mode  int
	True  int
	False int
}

var (
	STICKY_KEYS = Setting{
		Mode:  0,
		True:  0,
		False: 0,
	}
	DISABLE_CURSOR = Setting{
		Mode:  0,
		True:  0,
		False: 0,
	}

	PRESSED  = State{"keydown"}
	RELEASED = State{"keyup"}

	keyPressed = make(map[string]bool)
	mouseCords = [2]float32{0, 0}
)

func Init(window *windowing.Window) *Event {
	web.Document.AddEventListener("keydown", func(this js.Value, p []js.Value) interface{} {
		event := p[0]
		key := event.Get("key").String()

		keyPressed[key] = true

		return nil
	})

	web.Document.AddEventListener("keyup", func(this js.Value, p []js.Value) interface{} {
		event := p[0]
		key := event.Get("key").String()

		keyPressed[key] = false

		return nil
	})

	web.Document.AddEventListener("mousemove", func(this js.Value, p []js.Value) interface{} {
		event := p[0]

		mouseCords[0] = float32(event.Get("clientX").Float())
		mouseCords[1] = float32(event.Get("clientY").Float())

		return nil
	})

	return &Event{
		Window: window,
	}
}

func (e *Event) SetSetting(setting Setting, value bool) {
	web.Console.Warn("WARNING: SETSETTING IS NOT WORKING IN WEB BUILDS!")
}

func (e *Event) Key(key Key) State {
	if keyPressed[key.Key] {
		return PRESSED
	}
	return RELEASED
}

func (e *Event) MousePos() (float32, float32) {
	return mouseCords[0], mouseCords[1]
}
