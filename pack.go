package fltk

/*
#include "pack.h"
*/
import "C"
import "unsafe"

type Pack struct {
     Group
}

func NewPack(x, y, w, h int, text ...string) *Pack {
	i := &Pack{}
	initWidget(i, unsafe.Pointer(C.go_fltk_new_Pack(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return i
}
