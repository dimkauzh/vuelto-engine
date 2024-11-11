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

type Line struct {
	Renderer       *Renderer2D
	X1, Y1, X2, Y2 float32
	Color          [4]int
}

type Rect struct {
	Renderer            *Renderer2D
	X, Y, Width, Height float32
	Color               [4]int
}

// Loads a new line and returns a Line struct. Can be later drawn using Draw() method
func (r *Renderer2D) NewLine(x1, y1, x2, y2 float32, color [4]int) *Line {
	return &Line{
		Renderer: r,
		X1:       x1,
		Y1:       y1,
		X2:       x2,
		Y2:       y2,
		Color:    color,
	}
}

// Loads a new rect and returns a Rect struct. Can be later drawn using Draw() method
func (r *Renderer2D) NewRect(x, y, width, height float32, color [4]int) *Rect {
	return &Rect{
		Renderer: r,
		X:        x,
		Y:        y,
		Width:    width,
		Height:   height,
		Color:    color,
	}
}

// Draws the line loaded previously
func (l *Line) Draw() {
	l.Renderer.DrawLine(l.X1, l.Y1, l.X2, l.Y2, l.Color)
}

// Draws the rect loaded previously
func (r *Rect) Draw() {
	r.Renderer.DrawRect(r.X, r.Y, r.Width, r.Height, r.Color)
}
