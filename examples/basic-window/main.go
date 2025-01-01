package main

import vuelto "vuelto.pp.ua/pkg"

func main() {
	// This works in the web too! Only it would be so fun :(
	w := vuelto.NewWindow("hi", 800, 600, false)

	for !w.Close() {
		w.Refresh()
	}
}
