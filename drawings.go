package fltk

/*
#include <stdlib.h>
#include "drawings.h"
*/
import "C"
import "unsafe"

func Color(color uint) {
	C.go_fltk_color(C.uint(color))
}

func Draw(text string, x, y, w, h int, align Align) {
	textStr := C.CString(text)
	defer C.free(unsafe.Pointer(textStr))
	C.go_fltk_draw(textStr, C.int(x), C.int(y), C.int(w), C.int(h), C.uint(align))
}

func DrawBox(boxType BoxType, x, y, w, h int, color uint) {
	C.go_fltk_draw_box(
		C.int(boxType), C.int(x), C.int(y), C.int(w), C.int(h), C.uint(color))
}
