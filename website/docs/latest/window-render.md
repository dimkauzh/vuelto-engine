# 🎥 Windowing and Rendering Module

The main modules of Vuelto: the main window and the main renderer. Your entire project depends on these.

## 🌆 Overview

This module allows you to manage the main window and rendering operations, essential for any Vuelto project.

## 🏠 Window

To create a window, initialize a new instance and store it in a variable. Pass the following arguments:

- **Title**: The title of your window.
- **Width**: The width of the window in pixels.
- **Height**: The height of the window in pixels.
- **Resizable**: A boolean indicating whether the window can be resized.

### Example 1

```go
window := vuelto.NewWindow("My Game", 800, 600, true)
```

The renderer depends on this window to function.

## 🌈 Renderer

The renderer is responsible for drawing content onto the screen. To use it, initialize it using the window instance and store it in a variable.

### Example 2

```go
renderer := window.NewRenderer2D()
```

All drawing and rendering operations rely on this renderer.

## ♻️ Game Loop

Rendering operations should be performed within the game loop. This ensures that the application continues to respond to user input and updates the display correctly.

### Structure of the Game Loop

```go
for !window.Close() {
    // Perform rendering and update logic here
}
```

- **`window.Close()`**: This function returns `true` when the window is closed and `false` when it is active. Use it to control the loop's execution.

### Example Usage

```go
for !window.Close() {
 image.Draw()

  // ...

 window.Refresh()
}
```

Ensure all rendering tasks are enclosed within this loop to maintain smooth operation.
