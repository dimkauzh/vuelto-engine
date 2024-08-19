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
	STATIC_DRAW  = gl.Get("STATIC_DRAW")
	TRIANGLES    = gl.Get("TRIANGLES")

	FLOAT = gl.Get("FLOAT")
	FALSE = gl.Get("FALSE")
	TRUE  = gl.Get("TRUE")

	COLOR_BUFFER_BIT = gl.Get("COLOR_BUFFER_BIT")

	TEXTURE_2D         = gl.Get("TEXTURE_2D")
	TEXTURE_WRAP_S     = gl.Get("TEXTURE_WRAP_S")
	TEXTURE_WRAP_T     = gl.Get("TEXTURE_WRAP_T")
	TEXTURE_MIN_FILTER = gl.Get("TEXTURE_MIN_FILTER")
	TEXTURE_MAG_FILTER = gl.Get("TEXTURE_MAG_FILTER")
	CLAMP_TO_EDGE      = gl.Get("CLAMP_TO_EDGE")

	LINEAR              = gl.Get("LINEAR")
	RGBA                = gl.Get("RGBA")
	UNSIGNED_BYTE       = gl.Get("UNSIGNED_BYTE")
	SRC_ALPHA           = gl.Get("SRC_ALPHA")
	ONE_MINUS_SRC_ALPHA = gl.Get("ONE_MINUS_SRC_ALPHA")
	BLEND               = gl.Get("BLEND")
	DEPTH_BUFFER_BIT    = gl.Get("DEPTH_BUFFER_BIT")
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
