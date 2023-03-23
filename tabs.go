package fltk

/*
#include "tabs.h"
*/
import "C"
import "unsafe"

type Tabs struct {
	Group
}

func NewTabs(x, y, w, h int, text ...string) *Tabs {
	i := &Tabs{}
	initWidget(i, unsafe.Pointer(C.go_fltk_new_Tabs(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return i
}

func (t *Tabs) Value() int {
	return int(C.go_fltk_Tabs_value((*C.Fl_Tabs)(t.ptr())))
}

func (t *Tabs) SetValue(value int) {
	C.go_fltk_Tabs_set_value((*C.Fl_Tabs)(t.ptr()), (C.int)(value))
}
