package fltk

/*
#include "drawings.h"
*/
import "C"

func Color(color uint) {
	C.go_fltk_color(C.uint(color))
}

func Draw(text string, x, y, w, h int, align Align) {
	C.go_fltk_draw(C.CString(text), C.int(x), C.int(y), C.int(w), C.int(h), C.uint(align))
}

func DrawBox(boxType BoxType, x, y, w, h int, color uint) {
	C.go_fltk_draw_box(
		C.int(boxType), C.int(x), C.int(y), C.int(w), C.int(h), C.uint(color))
}
