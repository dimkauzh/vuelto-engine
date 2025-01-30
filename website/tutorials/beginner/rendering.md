<!-- markdownlint-disable md010 -->
# 📂 Basic Rendering with vuelto

Hello and welcome to the second tutorial in the beginner series! In this tutorial, we will learn how to render a window with Vuelto. We expect you to at least have the basics from the [first tutorial](new-project.md) setup and ready to go. Let's get started!

!!! note
    This tutorial is designed for the desktop platforms (Windows, macOS, and Linux). The changes for the web platform will be covered in the last paragraph of the tutorial.

## 🖌️ Settings up the renderer

So before the tutorial starts, we need to make sure we have a working renderer. The renderer is a struct that is responsible for rendering everything on the screen. We will use the `vuelto.NewRenderer` function to create a new renderer. Here is how you can do it:

```go
ren := w.NewRenderer2D()
```

And thats it for the renderer! Now we can move on to rendering different things!.

## 🚀 Rendering different shapes onto the window

So lets get started with some graphics. A colored background will be our start for now. So lets run the `ClearColor()` function from the `vuelto` package. This function takes in 4 arguments, the red, green, blue, and alpha values. The values follow RGBA color model, where each value is between 0 and 255. Here is an example of how to set the background color to a gray-ish color. Make sure to run this code inside of the game loop.

```go
// ...
for !w.Close() {
	// ...
	ren.ClearColor([4]int{100, 100, 100, 255}) // R, B, G, A
	// ...
}
// ...
```

And now if you run the code, you should see a gray window. You can change the values to get different colors.

Next up, we can get some geometry on the screen. Lets start off with a basic line. For this we first will need to initialize a new line outside of the game loop (`vuelto.NewLine`), and after that draw it inside of the game loop (`line.Draw`). The line arguments work as follows:

- `x1` and `y1` are the starting point of the line.
- `x2` and `y2` are the ending point of the line.
- `color` is the color of the line, following the RGBA color model (just like in `ClearColor`).

```go
// ...
line := ren.NewLine(0.5, 0.5, -0.5, -0.5, [4]int{10, 145, 245, 255})
// ...
```

And now we can render the line inside of the game loop.

```go
// ...
for !w.Close() {
	// ...
	line.Draw()
	// ...
}
// ...
```

And now you should see a line on the screen. You can change the values to get different lines!

Now that we got basic lines, lets move on to rectangles. It works really similar to lines. The rectangle arguments work as follows:

- `x` and `y` are the top left corner of the rectangle.
- `width` and `height` are the width and height of the rectangle.
- `color` is the color of the rectangle, following the RGBA color model.

First initialize a new rectangle outside of the game loop (`vuelto.NewRect`).

```go
rect := ren.NewRect(0, 0, 1, 1, [4]int{245, 145, 10, 255})
```

And now we can render the rectangle inside of the game loop.

```go
// ...
for !w.Close() {
  // ...
  rect.Draw()
  // ...
}
// ...
```

And now you should see a rectangle on the screen. You can change the values to get different rectangles!

## 🎨 Rendering images

Now that we got some basic shapes on the screen, lets move on to images. The will be a little more complex, but still pretty simple. First we will need to load an image. For this we will use the `ren.LoadImage` function. This function takes in the path to the image and returns a new image. Here is an example of how to load an image:

```go
img := ren.LoadImage("path/to/image.png", 0.5, 0.5, -0.5, 0.5)
```

The arguments work as follows:

- The first argument is the path to the image.
- The next 4 arguments are the x and y coordinates of the top left corner of the image, and the width and height of the image.

And now we can render the image inside of the game loop.

```go
// ...
for !w.Close() {
	// ...
	img.Draw()
	// ...
}
// ...
```

## 🌐 Rendering for the web

If you want to render your game for the web, you will need to make a few changes.

- You will need to adapt your image paths to the web. This means that you will need to change the path to the image to an embedded image.
Thats it!

### 🔄 Adapting the image paths

To adapt the image paths to the web, you will need to embed the images into the binary. This can be done with the `go:embed` directive. Here is an example of how to embed an image:

First import the `embed` package.

```go
// ...
import (
	// ...
	"embed"
	// ...
)
// ...
```

Then, create an embed filesystem with your image/images.

```go
// ...
//go:embed tree.png /path/to/image.png
var embeddedFiles embed.FS
// ...
```

And now create a new `ImageEmbed` object with the path to the image and the filesystem

```go
imageEmbed := vuelto.ImageEmbed{
	Filesystem: embeddedFiles,
	Image:      "<image-name>.png",
}
```

And now you can use the `LoadImage` function with the `imageEmbed` object.

```go
image := ren.LoadImage(imageEmbed, 0.5, 0.5, -0.5, 0.5)
```

And now you should be able to render the image on the web!
