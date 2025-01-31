package main

import (
	"embed"

	vuelto "vuelto.pp.ua/pkg"
)

//go:embed font.ttf
var embeddedFiles embed.FS

func main() {
	// This works in the web too!
	w := vuelto.NewWindow("hi", 800, 600, false, false)

	ren := w.NewRenderer2D()
	uir := w.NewUIRenderer()

	font := uir.LoadFont(vuelto.FontEmbed{Filesystem: embeddedFiles, Font: "font.ttf"}, "Hello, world!", -0.1, -0.1, 30)

	for !w.Close() {
		ren.ClearColor([4]int{100, 100, 100, 255})

		font.Draw()

		w.Refresh()
	}
}
