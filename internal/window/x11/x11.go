/*
 * Copyright (C) 2024 vuelto-org
 *
 * This file is part of the Vuelto project, licensed under the VL-Cv1.1 License.
 * Primary License: GNU GPLv3 or later (see <https://www.gnu.org/licenses/>).
 * If unmaintained, this software defaults to the MIT License as per Vuelto License V1.1,
 * at which point the copyright no longer applies.
 *
 * Distributed WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 */

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
