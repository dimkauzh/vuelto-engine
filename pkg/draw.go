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

package vuelto

import (
	"vuelto.pp.ua/internal/gl"
)

// Draws a new rect by the given x, y, width, height and color.
func (r *Renderer2D) DrawRect(x, y, width, height float32, color [4]int) {
	rect := r.NewRect(x, y, width, height, color)
	rect.Draw()
}

// Clears the screen with the specific color that is provided
func (r *Renderer2D) ClearColor(color [4]int) {
	gl.ClearColor(
		float32(color[0])/255,
		float32(color[1])/255,
		float32(color[2])/255,
		float32(color[3])/255,
	)
}

// Draws a new line by the given x1, y1, x2, y2 and color
func (r *Renderer2D) DrawLine(x1, y1, x2, y2 float32, color [4]int) {
	line := r.NewLine(x1, y1, x2, y2, color)
	line.Draw()
}
