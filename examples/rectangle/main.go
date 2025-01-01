package main

import (
	vuelto "vuelto.pp.ua/pkg"
)

func main() {
	w := vuelto.NewWindow("hi", 800, 600, false)

	ren := w.NewRenderer2D()

	rect := ren.NewRect(0, 0, -1, -1, [4]int{10, 145, 245, 255})
	rect2 := ren.NewRect(0, 0, 1, 1, [4]int{245, 145, 10, 255})
	line := ren.NewLine(0.5, 0.5, -0.5, -0.5, [4]int{10, 145, 245, 255})

	for !w.Close() {
		ren.ClearColor([4]int{100, 100, 100, 255})
		ren.DrawLine(-0.9, -0.9, 0.9, -0.9, [4]int{10, 145, 245, 255})

		rect.Draw()
		rect2.Draw()
		line.Draw()

		w.Refresh()

	}
}
