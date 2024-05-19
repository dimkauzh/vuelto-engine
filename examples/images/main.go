package main

import (
	vuelto "vuelto.me/pkg"
)

func main() {
	w := vuelto.NewWindow("Image Example - Vuelto", 800, 600, false)
	ren := w.NewRenderer2D()

	image := ren.LoadImage("test/image.png", 300, 300, 250, 250)
	image1 := ren.LoadImage("test/image.png", 100, 100, 150, 150)

	for !w.Close() {
		image.Draw()
		image1.Draw()
		w.Refresh()

	}
}
