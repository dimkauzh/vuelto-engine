//go:build js && wasm
// +build js,wasm

/*
 * Copyright (C) 2024 vuelto-org
 *
 * This file is part of the Vuelto project, licensed under the Vuelto License V1.
 * Primary License: GNU GPLv3 or later (see <https://www.gnu.org/licenses/>).
 * If unmaintained, this software defaults to the MIT License as per Vuelto License V1,
 * at which point the copyright no longer applies.
 *
 * Distributed WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 */

package webgl

import (
	"syscall/js"

	"vuelto.me/internal/window/web"
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

func DeleteShader(shader js.Value) {
	gl.Call("deleteShader", shader)
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
	gl.Call("bufferData", target, NewFloat32Array(data), usage)
}

func DeleteBuffer(buffer js.Value) {
	gl.Call("deleteBuffer", buffer)
}

func UseProgram(program js.Value) {
	gl.Call("useProgram", program)
}

func GetAttribLocation(program js.Value, name string) js.Value {
	return gl.Call("getAttribLocation", program, name)
}

func EnableVertexAttribArray(index int) {
	gl.Call("enableVertexAttribArray", index)
}

func VertexAttribPointer(index js.Value, size int, typ js.Value, normalized bool, stride int, offset int) {
	gl.Call("vertexAttribPointer", index, size, typ, normalized, stride, offset)
}

func GetUniformLocation(program js.Value, name string) js.Value {
	return gl.Call("getUniformLocation", program, name)
}

func Uniform1f(location js.Value, x float32) {
	gl.Call("uniform1f", location, x)
}

func Uniform2f(location js.Value, x, y float32) {
	gl.Call("uniform2f", location, x, y)
}

func Uniform3f(location js.Value, x, y, z float32) {
	gl.Call("uniform3f", location, x, y, z)
}

func Uniform4f(location js.Value, x, y, z, w float32) {
	gl.Call("uniform4f", location, x, y, z, w)
}

func DrawArrays(mode js.Value, first int, count int) {
	gl.Call("drawArrays", mode, first, count)
}

func GetShaderParameter(shader js.Value, pname js.Value) js.Value {
	return gl.Call("getShaderParameter", shader, pname)
}

func GetShaderInfoLog(shader js.Value) string {
	return gl.Call("getShaderInfoLog", shader).String()
}

func GetProgramParameter(program js.Value, pname js.Value) js.Value {
	return gl.Call("getProgramParameter", program, pname)
}

func GetProgramInfoLog(program js.Value) string {
	return gl.Call("getProgramInfoLog", program).String()
}

func CreateTexture() js.Value {
	return gl.Call("createTexture")
}

func BindTexture(target js.Value, texture js.Value) {
	gl.Call("bindTexture", target, texture)
}

func TexParameteri(target js.Value, pname js.Value, param js.Value) {
	gl.Call("texParameteri", target, pname, param)
}

func TexImage2D(target js.Value, level int, internalFormat js.Value, width, height, border int, format, typ, pixels js.Value) {
	gl.Call("texImage2D", target, level, internalFormat, width, height, border, format, typ, pixels)
}

func ClearColor(r, g, b, a float32) {
	gl.Call("clearColor", r, g, b, a)
}

func Clear(mask js.Value) {
	gl.Call("clear", mask)
}

func ClearDepth(depth float64) {
	gl.Call("clearDepth", depth)
}

func BlendFunc(sfactor, dfactor js.Value) {
	gl.Call("blendFunc", sfactor, dfactor)
}

func Enable(capability js.Value) {
	gl.Call("enable", capability)
}

func NewFloat32Array(values []float32) js.Value {
	return js.Global().Get("Float32Array").New(len(values)).Call("set", js.ValueOf(values))
}
