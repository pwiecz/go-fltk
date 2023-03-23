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
	C.go_fltk_Progress_set_maximum((*C.Fl_Progress)(p.ptr()), C.double(max))
}
func (p *Progress) Maximum() float64 {
	return float64(C.go_fltk_Progress_maximum((*C.Fl_Progress)(p.ptr())))
}
func (p *Progress) SetMinimum(max float64) {
	C.go_fltk_Progress_set_minimum((*C.Fl_Progress)(p.ptr()), C.double(max))
}
func (p *Progress) Minimum() float64 {
	return float64(C.go_fltk_Progress_minimum((*C.Fl_Progress)(p.ptr())))
}
func (p *Progress) SetValue(value float64) {
	C.go_fltk_Progress_set_value((*C.Fl_Progress)(p.ptr()), C.double(value))
}
func (p *Progress) Value() float64 {
	return float64(C.go_fltk_Progress_value((*C.Fl_Progress)(p.ptr())))
}
