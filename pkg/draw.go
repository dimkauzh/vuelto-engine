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

package vuelto

import gl "vuelto.me/internal/gl/legacy"

// Draws a new rect by the given x, y, width, height and color
func (r *Renderer2D) DrawRect(x, y, width, height float32, color [4]int) {
	gl.Begin(gl.QUADS)
	defer gl.End()

	gl.Color4f(
		float32(color[0])/255,
		float32(color[1])/255,
		float32(color[2])/255,
		float32(color[3])/255,
	)
	gl.Vertex2f(x, y)
	gl.Vertex2f(x+width, y)
	gl.Vertex2f(x+width, y+height)
	gl.Vertex2f(x, y+height)
}

// Clears the screen whith the specific color that is provided
func (r *Renderer2D) ClearColor(color [4]int) {
	gl.ClearColor(
		float32(color[0])/255,
		float32(color[1])/255,
		float32(color[2])/255,
		float32(color[3])/255,
	)
}

// Draws a new line by the given x1, x2, y1, y2 and color
func (r *Renderer2D) DrawLine(x1, x2, y1, y2 float32, color [4]int) {
	gl.LineWidth(1)

	gl.Begin(gl.LINES)
	defer gl.End()

	gl.Color4f(
		float32(color[0])/255,
		float32(color[1])/255,
		float32(color[2])/255,
		float32(color[3])/255,
	)
	gl.Vertex2f(x1, y1)
	gl.Vertex2f(x2, y2)
}
