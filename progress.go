package fltk

/*
#include "progress.h"
*/
import "C"
import "unsafe"

type Progress struct {
	widget
}

func NewProgress(x, y, w, h int, text ...string) *Progress {
	p := &Progress{}
	initWidget(p, unsafe.Pointer(C.go_fltk_new_Progress(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return p
}

func (p *Progress) SetMaximum(max float64) {
	C.go_fltk_Progress_set_maximum((*C.GProgress)(p.ptr), C.double(max))
}
func (p *Progress) SetMinimum(max float64) {
	C.go_fltk_Progress_set_minimum((*C.GProgress)(p.ptr), C.double(max))
}
func (p *Progress) SetValue(value float64) {
	C.go_fltk_Progress_set_value((*C.GProgress)(p.ptr), C.double(value))
}
