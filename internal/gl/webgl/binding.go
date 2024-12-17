//go:build js && wasm
// +build js,wasm

/*
 * Copyright (C) 2024 vuelto-org
 *
 * This file is part of the Vuelto project, licensed under the VL-Cv1.1 License.
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

	"vuelto.pp.ua/internal/window/web"
)

var (
	canvas web.Canvas
	gl     js.Value

	VERTEX_SHADER        js.Value
	FRAGMENT_SHADER      js.Value
	ARRAY_BUFFER         js.Value
	STATIC_DRAW          js.Value
	TRIANGLES            js.Value
	TRIANGLE_FAN         js.Value
	FLOAT                js.Value
	FALSE                js.Value
	TRUE                 js.Value
	COLOR_BUFFER_BIT     js.Value
	TEXTURE_2D           js.Value
	TEXTURE_WRAP_S       js.Value
	TEXTURE_WRAP_T       js.Value
	TEXTURE_MIN_FILTER   js.Value
	TEXTURE_MAG_FILTER   js.Value
	CLAMP_TO_EDGE        js.Value
	LINEAR               js.Value
	RGBA                 js.Value
	UNSIGNED_BYTE        js.Value
	SRC_ALPHA            js.Value
	ONE_MINUS_SRC_ALPHA  js.Value
	BLEND                js.Value
	DEPTH_BUFFER_BIT     js.Value
	NEAREST              js.Value
	UNSIGNED_SHORT       js.Value
	ELEMENT_ARRAY_BUFFER js.Value
	TEXTURE0             js.Value
)

func InitWebGL() {
	canvas = web.Document.GetElementById("vuelto")
	gl = canvas.GetContext("webgl2").JSContext

	VERTEX_SHADER = gl.Get("VERTEX_SHADER")
	FRAGMENT_SHADER = gl.Get("FRAGMENT_SHADER")
	ARRAY_BUFFER = gl.Get("ARRAY_BUFFER")
	STATIC_DRAW = gl.Get("STATIC_DRAW")
	TRIANGLES = gl.Get("TRIANGLES")
	TRIANGLE_FAN = gl.Get("TRIANGLE_FAN")
	FLOAT = gl.Get("FLOAT")
	FALSE = gl.Get("FALSE")
	TRUE = gl.Get("TRUE")
	COLOR_BUFFER_BIT = gl.Get("COLOR_BUFFER_BIT")
	TEXTURE_2D = gl.Get("TEXTURE_2D")
	TEXTURE_WRAP_S = gl.Get("TEXTURE_WRAP_S")
	TEXTURE_WRAP_T = gl.Get("TEXTURE_WRAP_T")
	TEXTURE_MIN_FILTER = gl.Get("TEXTURE_MIN_FILTER")
	TEXTURE_MAG_FILTER = gl.Get("TEXTURE_MAG_FILTER")
	CLAMP_TO_EDGE = gl.Get("CLAMP_TO_EDGE")
	LINEAR = gl.Get("LINEAR")
	RGBA = gl.Get("RGBA")
	UNSIGNED_BYTE = gl.Get("UNSIGNED_BYTE")
	SRC_ALPHA = gl.Get("SRC_ALPHA")
	ONE_MINUS_SRC_ALPHA = gl.Get("ONE_MINUS_SRC_ALPHA")
	BLEND = gl.Get("BLEND")
	DEPTH_BUFFER_BIT = gl.Get("DEPTH_BUFFER_BIT")
	NEAREST = gl.Get("NEAREST")
	UNSIGNED_SHORT = gl.Get("UNSIGNED_SHORT")
	ELEMENT_ARRAY_BUFFER = gl.Get("ELEMENT_ARRAY_BUFFER")
	TEXTURE0 = gl.Get("TEXTURE0")
}

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

func DeleteProgram(program js.Value) {
	gl.Call("deleteProgram", program)
}

func CreateVertexArray() js.Value {
	return gl.Call("createVertexArray")
}

func BindVertexArray(vao js.Value) {
	if vao.IsUndefined() || vao.IsNull() {
		gl.Call("bindVertexArray", nil)
	} else {
		gl.Call("bindVertexArray", vao)
	}
}

func DeleteVertexArray(vao js.Value) {
	gl.Call("deleteVertexArray", vao)
}

func CreateBuffer() js.Value {
	return gl.Call("createBuffer")
}

func BindBuffer(target js.Value, buffer js.Value) {
	if buffer.IsUndefined() || buffer.IsNull() {
		gl.Call("bindBuffer", target, nil)
	} else {
		gl.Call("bindBuffer", target, buffer)
	}
}

func BufferData(target js.Value, data js.Value, usage js.Value) {
	gl.Call("bufferData", target, data, usage)
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

func VertexAttribPointer(index int, size int, typ js.Value, normalized bool, stride int, offset int) {
	gl.Call("vertexAttribPointer", index, size, typ, normalized, stride, offset)
}

func GetUniformLocation(program js.Value, name string) js.Value {
	return gl.Call("getUniformLocation", program, name)
}

func GetUniform(program js.Value, location js.Value) js.Value {
	return gl.Call("getUniform", program, location)
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

func DrawElements(mode js.Value, count int, typ js.Value, offset int) {
	gl.Call("drawElements", mode, count, typ, offset)
}

func DrawElementsInstanced(mode js.Value, count int, typ js.Value, offset int, primcount int) {
	gl.Call("drawElementsInstanced", mode, count, typ, offset, primcount)
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

func DeleteTexture(texture js.Value) {
	gl.Call("deleteTexture", texture)
}

func BindTexture(target js.Value, texture js.Value) {
	if texture.IsUndefined() || texture.IsNull() {
		gl.Call("bindTexture", target, nil)
	} else {
		gl.Call("bindTexture", target, texture)
	}
}

func ActiveTexture(texture js.Value) {
	gl.Call("activeTexture", texture)
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

func Viewport(x, y, width, height int) {
	gl.Call("viewport", x, y, width, height)
}

func NewFloat32Array(values []float32) js.Value {
	array := js.Global().Get("Float32Array").New(len(values))
	for i, v := range values {
		array.SetIndex(i, v)
	}
	return array
}

func NewUint16Array(values []uint16) js.Value {
	array := js.Global().Get("Uint16Array").New(len(values))
	for i, v := range values {
		array.SetIndex(i, v)
	}
	return array
}

func Int32ToFloat32(input []int32) []float32 {
	output := make([]float32, len(input))
	for i, v := range input {
		output[i] = float32(v)
	}
	return output
}

func Int32ToUint16(input []int32) []uint16 {
	output := make([]uint16, len(input))
	for i, v := range input {
		output[i] = uint16(v)
	}
	return output
}

func NewUint16ArrayFromInt32(input []int32) js.Value {
	return NewUint16Array(Int32ToUint16(input))
}
