<!--markdownlint-disable md010-->
# 💡 Events module

Features for detecting user events.

## Types

```go
type Vector2D struct {
	X float64
	Y float64
}
```

Events include detecting keyboard keys, so here's the names Vuelto uses for identifying keys when asked.

```go
{
    "A",
    "B",
    "C",
    "D",
    "E",
    "F",
    "G",
    "H",
    "I",
    "J",
    "K",
    "L",
    "M",
    "N",
    "O",
    "P",
    "Q",
    "R",
    "S",
    "T",
    "U",
    "V",
    "W",
    "X",
    "Y",
    "Z",
    "Up",
    "Down",
    "Left",
    "Right",
    "Num0",
    "Num1",
    "Num2",
    "Num3",
    "Num4",
    "Num5",
    "Num6",
    "Num7",
    "Num8",
    "Num9",
    "Space",
    "Enter",
    "Escape",
    "Tab",
    "Shift",
    "Control",
    "Alt",
    "F1",
    "F2",
    "F3",
    "F4",
    "F5",
    "F6",
    "F7",
    "F8",
    "F9",
    "F10",
    "F11",
    "F12"
}
```

## Usage

For now you have the ability to detect key press and key release, with `KeyPressed()` and `KeyReleased()`. Both functions take a key name (as shown above) as a unique argument, and return a boolean indicating the status.

```go
if KeyPressed("E") {
    // this runs when i press E
}

if KeyReleased("F") {
    // this runs when i STOP pressing F (of course requires a previous press)
}
```

The return value of these functions is _reactive_, meaning their values changes immediately to reflect what the user's doing. With `if KeyPressed("E")`, the code inside the `if` statement will fire as many times as I press `E`.

For mouse position, you can use `MousePos()` with no args. It returns a `Vector2D`.

```go
// assume my mouse is at X 20 and Y 30
pos := win.MousePos() // { X: 20, Y: 30 }
pos.X // 20 (X)
pos.Y // 30 (Y)
pos.Pos() // { X: 20, Y: 30 }
```

Unlike key press events, this value does not constantly change, it gives the exact value for the moment of the call. Use loops if you constantly need the value.
