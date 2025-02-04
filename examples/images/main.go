package main

import (
	"embed"

	vuelto "vuelto.pp.ua/pkg"
)

//go:embed tree.png galaxy.png
var embeddedFiles embed.FS

func main() {
	// This works in the web too! This is because of the images being embedded!
	w := vuelto.NewWindow("Image Example - Vuelto", 800, 600, true, false)
	ren := w.NewRenderer2D()

	imageEmbed := vuelto.ImageEmbed{
		Filesystem: embeddedFiles,
		Image:      "tree.png",
	}

	imageEmbedTwo := vuelto.ImageEmbed{
		Filesystem: embeddedFiles,
		Image:      "galaxy.png",
	}

	imageOne := ren.LoadImage(imageEmbed, 0.7, 0.3, -0.5, 0.5)
	imageTwo := ren.LoadImage(imageEmbedTwo, 0, 0, 1, 1)

	for !w.Close() {
		imageTwo.Draw()
		imageOne.Draw()
		w.Refresh()
	}
}
