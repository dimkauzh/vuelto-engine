//go:build windows || linux || darwin
// +build windows linux darwin

package opengl

/*
#cgo linux LDFLAGS: -lGL
#cgo darwin LDFLAGS: -framework OpenGL
#cgo windows LDFLAGS: -lopengl32

#include "gl.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

var (
	VERTEX_SHADER   uint32 = uint32(C.GL_VERTEX_SHADER)
	FRAGMENT_SHADER uint32 = uint32(C.GL_FRAGMENT_SHADER)

	ARRAY_BUFFER uint32 = uint32(C.GL_ARRAY_BUFFER)
	STATIC_DRAW  uint32 = uint32(C.GL_STATIC_DRAW)
	TRIANGLES    uint32 = uint32(C.GL_TRIANGLES)

	FLOAT uint32 = uint32(C.GL_FLOAT)
	FALSE uint32 = uint32(C.GL_FALSE)
	TRUE  uint32 = uint32(C.GL_TRUE)

	COLOR_BUFFER_BIT uint32 = uint32(C.GL_COLOR_BUFFER_BIT)

	TEXTURE_2D         uint32 = uint32(C.GL_TEXTURE_2D)
	TEXTURE_WRAP_S     uint32 = uint32(C.GL_TEXTURE_WRAP_S)
	TEXTURE_WRAP_T     uint32 = uint32(C.GL_TEXTURE_WRAP_T)
	TEXTURE_MIN_FILTER uint32 = uint32(C.GL_TEXTURE_MIN_FILTER)
	TEXTURE_MAG_FILTER uint32 = uint32(C.GL_TEXTURE_MAG_FILTER)
	CLAMP_TO_EDGE      uint32 = uint32(C.GL_CLAMP_TO_EDGE)

	LINEAR              uint32 = uint32(C.GL_LINEAR)
	RGBA                uint32 = uint32(C.GL_RGBA)
	UNSIGNED_BYTE       uint32 = uint32(C.GL_UNSIGNED_BYTE)
	SRC_ALPHA           uint32 = uint32(C.GL_SRC_ALPHA)
	ONE_MINUS_SRC_ALPHA uint32 = uint32(C.GL_ONE_MINUS_SRC_ALPHA)
	BLEND               uint32 = uint32(C.GL_BLEND)
	DEPTH_BUFFER_BIT    uint32 = uint32(C.GL_DEPTH_BUFFER_BIT)
)

func CreateShader(shaderType uint32) uint32 {
	return uint32(C.glCreateShader(C.uint(shaderType)))
}

func ShaderSource(shader uint32, source string) {
	csource := C.CString(source)
	defer C.free(unsafe.Pointer(csource))
	C.glShaderSource(C.uint(shader), 1, &csource, nil)
}

func CompileShader(shader uint32) {
	C.glCompileShader(C.uint(shader))
}

func AttachShader(program, shader uint32) {
	C.glAttachShader(C.uint(program), C.uint(shader))
}

func DeleteShader(shader uint32) {
	C.glDeleteShader(C.uint(shader))
}

func CreateProgram() uint32 {
	return uint32(C.glCreateProgram())
}

func LinkProgram(program uint32) {
	C.glLinkProgram(C.uint(program))
}

func UseProgram(program uint32) {
	C.glUseProgram(C.uint(program))
}

func DeleteProgram(program uint32) {
	C.glDeleteProgram(C.uint(program))
}

func GenBuffers(n int32, buffers *uint32) {
	C.glGenBuffers(C.int(n), (*C.uint)(unsafe.Pointer(buffers)))
}

func DeleteBuffers(n int32, buffers *uint32) {
	C.glDeleteBuffers(C.int(n), (*C.uint)(unsafe.Pointer(buffers)))
}

func BindVertexArray(array uint32) {
	C.glBindVertexArray(C.uint(array))
}

func DeleteVertexArrays(n int32, arrays *uint32) {
	C.glDeleteVertexArrays(C.int(n), (*C.uint)(unsafe.Pointer(arrays)))
}

func BindBuffer(target, buffer uint32) {
	C.glBindBuffer(C.uint(target), C.uint(buffer))
}

func BufferData(target uint32, size int, data unsafe.Pointer, usage uint32) {
	C.glBufferData(C.uint(target), C.sizeiptr(size), data, C.uint(usage))
}

func VertexAttribPointer(index, size, stride int32, pointer unsafe.Pointer) {
	C.glVertexAttribPointer(C.uint(index), C.int(size), C.uint(GL_FLOAT), C.GLboolean(GL_FALSE), C.int(stride), pointer)
}

func EnableVertexAttribArray(index uint32) {
	C.glEnableVertexAttribArray(C.uint(index))
}

func ClearColor(red, green, blue, alpha float32) {
	C.glClearColor(C.float(red), C.float(green), C.float(blue), C.float(alpha))
}

func DrawArrays(mode uint32, first, count int32) {
	C.glDrawArrays(C.uint(mode), C.int(first), C.int(count))
}

func GetUniformLocation(program uint32, name string) int32 {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return int32(C.glGetUniformLocation(C.uint(program), cname))
}

func Uniform4f(location int32, v0, v1, v2, v3 float32) {
	C.glUniform4f(C.int(location), C.float(v0), C.float(v1), C.float(v2), C.float(v3))
}

func Clear(mask uint32) {
	C.glClear(C.uint(mask))
}

func Enable(cap uint32) {
	C.glEnable(C.uint(cap))
}

func Viewport(x, y, width, height int32) {
	C.glViewport(C.int(x), C.int(y), C.int(width), C.int(height))
}

func Ortho(left, right, bottom, top, near, far float64) {
	C.glOrtho(C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(near), C.GLdouble(far))
}

func BlendFunc(sfactor, dfactor uint32) {
	C.glBlendFunc(C.uint(sfactor), C.uint(dfactor))
}

func DrawElements(mode uint32, count int32, typ uint32, indices unsafe.Pointer) {
	C.glDrawElements(C.uint(mode), C.int(count), C.uint(typ), indices)
}

func GenTextures(n int32, textures *uint32) {
	C.glGenTextures(C.int(n), (*C.uint)(unsafe.Pointer(textures)))
}

func DeleteTextures(n int32, textures *uint32) {
	C.glDeleteTextures(C.int(n), (*C.uint)(unsafe.Pointer(textures)))
}

func BindTexture(target, texture uint32) {
	C.glBindTexture(C.uint(target), C.uint(texture))
}

func TexParameteri(target, pname, param uint32) {
	C.glTexParameteri(C.uint(target), C.uint(pname), C.int(param))
}

func TexImage2D(target, level, internalformat, width, height, border, format, typ uint32, pixels unsafe.Pointer) {
	C.glTexImage2D(C.uint(target), C.int(level), C.int(internalformat), C.int(width), C.int(height), C.int(border), C.uint(format), C.uint(typ), pixels)
}

func GenerateMipmap(target uint32) {
	C.glGenerateMipmap(C.uint(target))
}
