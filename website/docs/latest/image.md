<!--markdownlint-disable md010-->
# 🖼️ Image module

Features for rendering images

## Types

```go
type ImageEmbed struct {
	Filesystem embed.FS
	Image      string
}
```

## Usage

!!! note
    For this you will need a [renderer](window-render.md) initialized.

First, load an image using `LoadImage()` and specify file path, pos X and Y, width, and height.

```go
image := renderer.LoadImage("path/to/image.png", 10, 10, 50, 50) // returns an Image (internal type)
```

Then, whenever you feel like it, draw it.

```go
image.Draw() // renders your image
```

We recommend using raw paths, however, **heads up as they won't work on the web.**

For supporting web, you'll need to use the same func, but passing an `ImageEmbed` instead of a raw path. See the `ImageEmbed` struct above.

```go
var embeddedFiles embed.FS

imageEmbed := vuelto.ImageEmbed{
    Filesystem: embeddedFiles,
    Image:      "image.png",
}

image := renderer.LoadImage(imageEmbed, 0, 0, 1, 1)

for !window.Close() {
    image.Draw()
    window.Refresh()
}
```
