package main

import (
	vuelto "vuelto.pp.ua/pkg"
)

func main() {
	// This does not work in the web, as it doest support transparency :(
	w := vuelto.NewWindow("Image Example - Vuelto", 800, 600, false, true) // Change the second argument to true

	for !w.Close() {
		w.Refresh()
	}
}
