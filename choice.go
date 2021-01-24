package fltk

/*
#include "choice.h"
*/
import "C"
import "unsafe"

type Choice struct {
	menu
}

func NewChoice(x, y, w, h int, text ...string) *Choice {
	c := &Choice{}
	initWidget(c, unsafe.Pointer(C.go_fltk_new_Choice(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return c
}
