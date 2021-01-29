package fltk

/*
#include "text_editor.h"
*/
import "C"
import "unsafe"

type TextEditor struct {
     widget
}

func NewTextEditor(x, y, w, h int, text... string) *TextEditor {
	t := &TextEditor{}
	initWidget(t, unsafe.Pointer(C.go_fltk_new_TextEditor(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return t
}
