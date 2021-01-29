package fltk

/*
#include "box.h"
*/
import "C"
import "unsafe"

type Box struct {
	widget
}

func NewBox(boxType BoxType, x, y, w, h int, text ...string) *Box {
	b := &Box{}
	initWidget(b, unsafe.Pointer(C.go_fltk_new_Box(C.int(boxType), C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return b
}
