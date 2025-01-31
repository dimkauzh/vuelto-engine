package main

import (
	vuelto "vuelto.pp.ua/pkg"
)

func main() {
	// This works in the web too! Only it wouldn't be so fun :(
	win := vuelto.NewWindow("hi", 800, 600, false, false)

	for !win.Close() {
		win.Refresh()
	}
}
