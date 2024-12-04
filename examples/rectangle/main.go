package main

import (
	vuelto "vuelto.pp.ua/pkg"
)

func main() {
	w := vuelto.NewWindow("hi", 800, 600, false)

	ren := w.NewRenderer2D()

	for !w.Close() {
		ren.ClearColor([4]int{100, 100, 100, 255})

		ren.DrawRect(0, 0, 500, 500, [4]int{10, 145, 245, 255})

		w.Refresh()

	}
}
