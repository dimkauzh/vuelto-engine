package main

import (
	vuelto "vuelto.pp.ua/pkg"
)

func main() {
	// Warning: This example is not working in the web! This is due to the web not supporting multiple windows.
	w1 := vuelto.NewWindow("hi", 800, 600, false)
	w2 := vuelto.NewWindow("hi2", 800, 600, false)

	ren1 := w1.NewRenderer2D()
	ren2 := w2.NewRenderer2D()

	image := ren2.LoadImage("test/image.png", 300, 300, 250, 250)
	image1 := ren2.LoadImage("test/tree.png", 100, 100, 150, 150)

	for !w1.Close() && !w2.Close() {
		w1.SetCurrent()
		ren1.ClearColor([4]int{100, 100, 100, 255})

		ren1.DrawRect(0, 0, 500, 500, [4]int{10, 145, 245, 255})

		w1.Refresh()
		w2.SetCurrent()

		image.Draw()
		image1.Draw()

		w2.Refresh()

	}
}
