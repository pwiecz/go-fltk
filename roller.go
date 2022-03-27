package fltk

/*
#include "roller.h"
*/
import "C"
import "unsafe"

type Roller struct {
	valuator
}

func NewRoller(x, y, w, h int, text ...string) *Roller {
	r := &Roller{}
	initWidget(r, unsafe.Pointer(C.go_fltk_new_Roller(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return r
}
