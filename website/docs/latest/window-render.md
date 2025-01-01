
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

Drawing and rendering depend on this renderer.

### Window close loop

Rendering should happen inside of the Window's close loop. This is:

```go
for !window.Close() {
    // do stuff
}
```

window.Close() will become true when the window it closed, and will return false when being active. You should render your stuff in there.
