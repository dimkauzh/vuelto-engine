package main

import vuelto "vuelto.pp.ua/pkg"

func main() {
	w := vuelto.NewWindow("hi", 800, 600, false)

	for !w.Close() {
		w.Refresh()
	}
}
