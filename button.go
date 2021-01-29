package fltk

/*
#include "button.h"
*/
import "C"
import "unsafe"

type Button struct {
	Widget
}

func NewButton(x, y, w, h int, text ...string) *Button {
	i := &Button{}
	initWidget(i, unsafe.Pointer(C.go_fltk_new_Button(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return i
}

func (b *Button) Value() bool {
	return C.go_fltk_Button_value((*C.Fl_Button)(b.ptr)) != C.char(0)
}
func (b *Button) SetValue(val bool) {
	if val {
		C.go_fltk_Button_set_value((*C.Fl_Button)(b.ptr), 1)
	} else {
		C.go_fltk_Button_set_value((*C.Fl_Button)(b.ptr), 0)
	}
}
