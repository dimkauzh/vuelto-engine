package main

import (
	"log"

	"vuelto.me/internal/windowing"
)

func main() {
	win, err := windowing.InitWindow()
	if err != nil {
		log.Fatalf("Failed to initialise: %s", err)
	}

	win.Resizable = true
	win.Title = "Test"

	win.Width = 500
	win.Height = 500

	win.Create()

	for {
		win.DrawingTest()

		win.HandleEvents()
		win.UpdateBuffers()
	}
	win.Close()
}
