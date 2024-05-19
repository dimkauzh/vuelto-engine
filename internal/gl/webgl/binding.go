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
	VERTEX_SHADER   = gl.Get("VERTEX_SHADER")
	FRAGMENT_SHADER = gl.Get("FRAGMENT_SHADER")

	ARRAY_BUFFER = gl.Get("ARRAY_BUFFER")

	FLOAT = gl.Get("FLOAT")
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

func BindBuffer(target js.Value, buffer js.Value) {
	gl.Call("bindBuffer", target, buffer)
}

func BufferData(target js.Value, data []float32, usage js.Value) {
	gl.Call("bufferData", target, js.TypedArrayOf(data), usage)
}

func UseProgram(program js.Value) {
	gl.Call("useProgram", program)
}

func GetAttribLocation(program js.Value, name string) js.Value {
	return gl.Call("getAttribLocation", program, name)
}

func EnableVertexAttribArray(index js.Value) {
	gl.Call("enableVertexAttribArray", index)
}

func VertexAttribPointer(index js.Value, size int, typ js.Value, normalized bool, stride int, offset int) {
	gl.Call("vertexAttribPointer", index, size, typ, normalized, stride, offset)
}

func GetUniformLocation(program js.Value, name string) js.Value {
	return gl.Call("getUniformLocation", program, name)
}

func DrawArrays(mode js.Value, first int, count int) {
	gl.Call("drawArrays", mode, first, count)
}
