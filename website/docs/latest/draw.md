<!--markdownlint-disable md010-->
# 🎨 Draw module

A module for drawing UI elements, such as rects.

## Types

```go
[4]int{R, G, B, A} // for colors
```

## Usage

!!! note
    For this you will need a [renderer](renderer.md) initialized.

For now this module is pretty basic. You can draw a rectangle using the `DrawRect()` function, passing X position, Y position, width & height, and an RGBA color (int array as described above).

```go
myRect := renderer.DrawRect(100, 100, 100, 100, [4]int{255, 255, 255, 255}) // draws a white square
```

You can also use `DrawLine()` to draw line that goes from `x1`, `y1` to `x2`, `y2`, and again, with a color.

```go
myLine := renderer.DrawLine(0.5, 0.5, -0.5, -0.5, [4]int{255, 255, 255, 255})
```

You can also create them and render them later, using "New" instead of "Draw", with the same arguments.

```go
myRect := renderer.NewRect(100, 100, 100, 100, [4]int{255, 255, 255, 255})
myLine := renderer.NewLine(0.5, 0.5, -0.5, -0.5, [4]int{255, 255, 255, 255})

// then, when you feel like it, draw them
myRect.Draw();
myLine.Draw();
```

If you want to "clear" the window, you can use `ClearColor()`, passing a color that'll be used as the background of the window.

```go
renderer.ClearColor([4]int{255, 255, 255, 255}) // flashbang!
```
