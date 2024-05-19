//go:build windows || linux || darwin
// +build windows linux darwin

package opengl

/*
#cgo linux LDFLAGS: -lGL
#cgo darwin LDFLAGS: -framework OpenGL
#cgo windows LDFLAGS: -lopengl32

#include "opengl/gl.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

const (
	GL_VERTEX_SHADER   = uint32(C.GL_VERTEX_SHADER)
	GL_FRAGMENT_SHADER = uint32(C.GL_FRAGMENT_SHADER)

	GL_ARRAY_BUFFER = uint32(C.GL_ARRAY_BUFFER)
	GL_STATIC_DRAW  = uint32(C.GL_STATIC_DRAW)
	GL_TRIANGLES    = uint32(C.GL_TRIANGLES)

	GL_FLOAT = uint32(C.GL_FLOAT)
	GL_FALSE = uint32(C.GL_FALSE)
	GL_TRUE  = uint32(C.GL_TRUE)

	GL_COLOR_BUFFER_BIT = uint32(C.GL_COLOR_BUFFER_BIT)

	TEXTURE_2D         = uint32(C.GL_TEXTURE_2D)
	TEXTURE_WRAP_S     = uint32(C.GL_TEXTURE_WRAP_S)
	TEXTURE_WRAP_T     = uint32(C.GL_TEXTURE_WRAP_T)
	TEXTURE_MIN_FILTER = uint32(C.GL_TEXTURE_MIN_FILTER)
	TEXTURE_MAG_FILTER = uint32(C.GL_TEXTURE_MAG_FILTER)
	CLAMP_TO_EDGE      = uint32(C.GL_CLAMP_TO_EDGE)

	LINEAR              = uint32(C.GL_LINEAR)
	RGBA                = uint32(C.GL_RGBA)
	UNSIGNED_BYTE       = uint32(C.GL_UNSIGNED_BYTE)
	SRC_ALPHA           = uint32(C.GL_SRC_ALPHA)
	ONE_MINUS_SRC_ALPHA = uint32(C.GL_ONE_MINUS_SRC_ALPHA)
	BLEND               = uint32(C.GL_BLEND)
	DEPTH_BUFFER_BIT    = uint32(C.GL_DEPTH_BUFFER_BIT)
	COLOR_BUFFER_BIT    = uint32(C.GL_COLOR_BUFFER_BIT)
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

func CreateProgram() uint32 {
	return uint32(C.glCreateProgram())
}

func AttachShader(program, shader uint32) {
	C.glAttachShader(C.uint(program), C.uint(shader))
}

func LinkProgram(program uint32) {
	C.glLinkProgram(C.uint(program))
}

func DeleteShader(shader uint32) {
	C.glDeleteShader(C.uint(shader))
}

func UseProgram(program uint32) {
	C.glUseProgram(C.uint(program))
}

func GenBuffers(n int32, buffers *uint32) {
	C.glGenBuffers(C.int(n), (*C.uint)(unsafe.Pointer(buffers)))
}

func BindVertexArray(array uint32) {
	C.glBindVertexArray(C.uint(array))
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

func Clear(mask uint32) {
	C.glClear(C.uint(mask))
}

func Enable(cap uint32) {
	C.glEnable(C.uint(cap))
}

func UseProgram(program uint32) {
	C.glUseProgram(C.uint(program))
}

func DrawArrays(mode uint32, first, count int32) {
	C.glDrawArrays(C.uint(mode), C.int(first), C.int(count))
}
