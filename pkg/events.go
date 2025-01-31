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

package vuelto

import "vuelto.pp.ua/internal/event"

var (
	Keys     = event.KeyMap
	keyState = make(map[event.Key]bool)
)

func (w *Window) KeyPressed(key event.Key) bool {
	return w.Event.Key(key) == event.PRESSED
}

func (w *Window) KeyReleased(key event.Key) bool {
	return w.Event.Key(key) == event.RELEASED
}

func (w *Window) KeyPressedOnce(key event.Key) bool {
	if pressed := w.KeyPressed(key); pressed && !keyState[key] {
		keyState[key] = true
		return true
	} else if !pressed {
		keyState[key] = false
	}
	return false
}

func (w *Window) MousePos() *Vector2D {
	return NewVector2D(w.Event.MousePos())
}
