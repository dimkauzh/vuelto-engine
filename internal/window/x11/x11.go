package x11

/*
#cgo LDFLAGS: -lX11


#include <X11/Xlib.h>
#include <GL/gl.h>
#include <GL/glx.h>
*/
import "C"

func XOpenDisplay() {
	C.XOpenDisplay(nil)
}

func XCloseDisplay(display *C.Display) {
	C.XCloseDisplay(nil)
}
