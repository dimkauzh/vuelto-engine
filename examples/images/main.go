package main

import (
	"embed"

	vuelto "vuelto.pp.ua/pkg"
)

//go:embed tree.png galaxy.png
var embeddedFiles embed.FS

func main() {
	w := vuelto.NewWindow("Image Example - Vuelto", 800, 600, false)
	ren := w.NewRenderer2D()

	iembed := vuelto.ImageEmbed{
		Filesystem: embeddedFiles,
		Image:      "tree.png",
	}

	i2embed := vuelto.ImageEmbed{
		Filesystem: embeddedFiles,
		Image:      "galaxy.png",
	}

	image1 := ren.LoadImage(iembed, 0.5, 0.5, -0.5, 0.5)
	image := ren.LoadImage(i2embed, 0, 0, 1, 1)

	for !w.Close() {
		image1.Draw()
		image.Draw()
		w.Refresh()

	}
}
