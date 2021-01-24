package fltk

/*
#include "tabs.h"
*/
import "C"
import "unsafe"

type Tabs struct {
	Group
}

func NewTabs(x, y, w, h int, text... string) *Tabs {
	i := &Tabs{}
	initWidget(i, unsafe.Pointer(C.go_fltk_new_Tabs(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return i
}
