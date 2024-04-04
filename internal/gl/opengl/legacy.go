//go:build windows || linux || darwin
// +build windows linux darwin

package gl

/*
#cgo linux LDFLAGS: -lGL
#cgo darwin LDFLAGS: -framework OpenGL
#cgo windows LDFLAGS: -lopengl32

#include "gl.h"
*/
import "C"
import "unsafe"

const (
	TEXTURE_2D          = uint32(C.GL_TEXTURE_2D)
	TEXTURE_WRAP_S      = uint32(C.GL_TEXTURE_WRAP_S)
	TEXTURE_WRAP_T      = uint32(C.GL_TEXTURE_WRAP_T)
	TEXTURE_MIN_FILTER  = uint32(C.GL_TEXTURE_MIN_FILTER)
	TEXTURE_MAG_FILTER  = uint32(C.GL_TEXTURE_MAG_FILTER)
	CLAMP_TO_EDGE       = uint32(C.GL_CLAMP_TO_EDGE)
	LINEAR              = uint32(C.GL_LINEAR)
	LINES               = uint32(C.GL_LINES)
	QUADS               = uint32(C.GL_QUADS)
	RGBA                = uint32(C.GL_RGBA)
	UNSIGNED_BYTE       = uint32(C.GL_UNSIGNED_BYTE)
	SRC_ALPHA           = uint32(C.GL_SRC_ALPHA)
	ONE_MINUS_SRC_ALPHA = uint32(C.GL_ONE_MINUS_SRC_ALPHA)
	BLEND               = uint32(C.GL_BLEND)
	DEPTH_BUFFER_BIT    = uint32(C.GL_DEPTH_BUFFER_BIT)
	COLOR_BUFFER_BIT    = uint32(C.GL_COLOR_BUFFER_BIT)
)

func Begin(state uint32) {
	C.glBegin(C.uint(state))
}

func End() {
	C.glEnd()
}

func Color3f(r, g, b float32) {
	C.glColor3f(C.float(r), C.float(g), C.float(b))
}

func Color4f(r, g, b, a float32) {
	C.glColor4f(C.float(r), C.float(g), C.float(b), C.float(a))
}

func Vertex2f(x, y float32) {
	C.glVertex2f(C.float(x), C.float(y))
}

func ClearColor(r, g, b, a float32) {
	C.glClearColor(C.float(r), C.float(g), C.float(b), C.float(a))
}

func GenTextures(n int32, textures *uint32) {
	C.glGenTextures(C.int(n), (*C.uint)(unsafe.Pointer(textures)))
}

func DeleteTextures(n int, textures *uint32) {
	C.glDeleteTextures(C.int(n), (*C.uint)(unsafe.Pointer(textures)))
}

func BindTexture(target, texture uint32) {
	C.glBindTexture(C.uint(target), C.uint(texture))
}

func TexParameteri(target, pname, param uint32) {
	C.glTexParameteri(C.uint(target), C.uint(pname), C.int(param))
}

func TexCoord2f(s, t float32) {
	C.glTexCoord2f(C.float(s), C.float(t))
}

func TexImage2D(target, level, internalFormat uint32, width, height int, border, format, typ uint32, pixels []byte) {
	C.glTexImage2D(C.uint(target), C.int(level), C.int(internalFormat), C.int(width), C.int(height), C.int(border), C.uint(format), C.uint(typ), C.CBytes(pixels))
}

func Clear(mask uint32) {
	C.glClear(C.uint(mask))
}

func Enable(cap uint32) {
	C.glEnable(C.uint(cap))
}

func BlendFunc(sfactor, dfactor uint32) {
	C.glBlendFunc(C.uint(sfactor), C.uint(dfactor))
}

func Ortho(left, right, bottom, top, near, far float64) {
	C.glOrtho(C.double(left), C.double(right), C.double(bottom), C.double(top), C.double(near), C.double(far))
}

func Viewport(x, y, width, height int) {
	C.glViewport(C.int(x), C.int(y), C.int(width), C.int(height))
}

func LineWidth(width float32) {
	C.glLineWidth(C.float(width))
}
