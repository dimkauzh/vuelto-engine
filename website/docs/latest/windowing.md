# 🪟 Window docs

The main window, and the main dependency of your project.

## Usage

Simply initialize a new window and store it in a variable. Pass a title, width, height, and whether it's resizable or not as arguments.

```go
window := NewWindow("title", 800, 600, true)
```

With your new window you can use the renderer, from which most functions come out.
