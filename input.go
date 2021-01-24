package fltk

/*
#include "input.h"
*/
import "C"
import "unsafe"

type Input struct {
	Widget
}

func NewInput(x, y, w, h int, text ...string) *Input {
	i := &Input{}
	initWidget(i, unsafe.Pointer(C.go_fltk_new_Input(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return i
}

func (i *Input) Value() string {
	return C.GoString(C.go_fltk_Input_value((*C.Fl_Input)(i.ptr)))
}
func (i *Input) SetValue(value string) bool {
	return C.go_fltk_Input_set_value((*C.Fl_Input)(i.ptr), C.CString(value)) != C.int(0)
}
func (i *Input) Resize(x int, y int, w int, h int) {
	C.go_fltk_Input_resize((*C.Fl_Input)(i.ptr), C.int(x), C.int(y), C.int(w), C.int(h))
}
