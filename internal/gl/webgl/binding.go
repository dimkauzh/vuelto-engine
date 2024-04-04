//go:build js && wasm
// +build js,wasm

package webgl

import (
  "syscall/js"
  "vuelto.me/internal/windowing/web"
)


var canvas web.Canvas = web.Document.GetElementById("vuelto")
var gl web.Context = canvas.GetContext("webgl")

var (
  VERTEX_SHADER = gl.Get("VERTEX_SHADER")
  FRAGMENT_SHADER = gl.Get("FRAGMENT_SHADER")
)

func CreateShader(inputType js.Value) js.Value {
  return gl.Call("createShader", inputType)
}

func ShaderSource(shader js.Value, source string) {
  gl.Call("shaderSource", shader, source)
}

func CompileShader(shader js.Value) {
  gl.Call("compileShader", shader)
}

func CreateProgram() js.Value {
  return gl.Call("createProgram")
}

func AttachShader(program, shaderType js.Value) {
  gl.Call("attachShader", program, shaderType)
}

func LinkProgram(program js.Value) {
  gl.Call("linkProgram", program)
}

func CreateBuffer() js.Value {
  return gl.Call("createBuffer")
}

