
# 🎥 Windowing and rendering module

The main modules of Vuelto - the main window and the main renderer. Your entire project depends on these.

## Usage

### Window

Simply initialize a new window and store it in a variable (your window). Pass a title, width, height, and whether it's resizable or not as arguments.

```go
window := NewWindow("title", 800, 600, true)
```

The renderer depends on this.

### Renderer

Simply initialize it by calling it and storing it into a variable (your renderer).

```go
renderer := window.NewRenderer2D()
```

Drawing, vectors, and others depend on this renderer.
