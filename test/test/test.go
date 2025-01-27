package main

import (
	"fmt"

	vuelto "vuelto.pp.ua/pkg"
)

func main() {
	w1 := vuelto.NewWindow("hi", 800, 600, false)
	w2 := vuelto.NewWindow("hi2", 800, 600, false)

	ren1 := w1.NewRenderer2D()
	ren2 := w2.NewRenderer2D()

	w1.SetCurrent()
	rect := ren1.NewRect(0, 0, 0.5, 0.5, [4]int{10, 145, 245, 255})
	line := ren1.NewLine(0.1, 0.1, 0.4, 0.4, [4]int{10, 145, 245, 255})

	w2.SetCurrent()
	image := ren2.LoadImage("test/test/tree.png", 0.1, -0.1, 0.4, -0.4)
	image1 := ren2.LoadImage("test/test/galaxy.png", -0.1, -0.1, 0.4, 0.4)
	image2 := ren2.LoadImage(vuelto.ImageHTTP{
		Url: "https://dev-tester.com/content/images/2021/12/blog_cover_further_api_testing_with_http_toolkit.png",
	}, -0.1, 0.1, 0.4, 0.4)

	for !w1.Close() && !w2.Close() {
		w1.SetCurrent()
		ren1.ClearColor([4]int{100, 100, 100, 255})

		if w1.KeyPressed(vuelto.Keys["Left"]) {
			rect.Pos.X = rect.Pos.X - 0.5*w1.GetDeltaTime()
		} else if w1.KeyPressed(vuelto.Keys["Right"]) {
			rect.Pos.X = rect.Pos.X + 0.5*w1.GetDeltaTime()
		}

		fmt.Println(w1.MousePos())

		rect.Draw()

		line.Draw()
		image2.Draw()

		w1.Refresh()
		w2.SetCurrent()

		image.Draw()
		image1.Draw()

		w2.Refresh()

	}
}
