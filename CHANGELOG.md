# Vuelto Changelog

## Vuelto 1.1.1 (05/01/2025)

### Fixes

- Fixed web builds not working.

## Vuelto 1.1.0 (05/01/2025)

### Breaking changes

- Updated to OpenGL 3.3 Core, breaking compatibility with unsupported hardware.
- `DrawLine()` now takes arguments in the order `x1 y1 x2 y2`, instead of `x1 x2 y1 y2`.
- `Image` struct now uses a `Pos` parameter with `Vector2D` type instead of `X,Y` `float32` param for position.
- `Line` struct now uses `Pos1` and `Pos2` params with `Vector2D` type instead of `X1, Y1, X2, Y2`.
- `Rect` struct now uses `Pos` param with `Vector2D` type instead of `X, Y`.

### Additions

- Support for targeting web (WASM + WebGL 2.0).
- Event system
  - KeyPressed, KeyPressedOnce, and KeyReleased
  - Mouse position
- `GetDeltaTime()`
- Framerate managing
  - `GetFPS()`
  - `SetFPS()`

### Changes

- The above mentioned breaking changes.
- Params `X` and `Y` (and `Z`) for `Vector2D` and `Vector3D` are now of type `float32` instead of `float64`.
- Improved performance.

---

Changes prior to 1.1.0 aren't logged. Use GitHub's diff view if you want to see changes from older versions.
