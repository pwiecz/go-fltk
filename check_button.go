package fltk

/*
#include "check_button.h"
*/
import "C"
import "unsafe"

type CheckButton struct {
	Button
}

func NewCheckButton(x, y, w, h int, text ...string) *CheckButton {
	i := &CheckButton{}
	initWidget(i, unsafe.Pointer(C.go_fltk_new_Check_Button(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return i
}
