package fltk

/*
#include "slider.h"
*/
import "C"
import "unsafe"

type Slider struct {
	valuator
}

func NewSlider(x, y, w, h int, text ...string) *Slider {
	s := &Slider{}
	initWidget(s, unsafe.Pointer(C.go_fltk_new_Slider(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return s
}

type ValueSlider struct {
	Slider
}

func NewValueSlider(x, y, w, h int, text ...string) *ValueSlider {
	s := &ValueSlider{}
	initWidget(s, unsafe.Pointer(C.go_fltk_new_Value_Slider(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return s
}

func (s *ValueSlider) SetTextFont(font Font) {
	C.go_fltk_Value_Slider_set_textfont((*C.GValue_Slider)(s.ptr()), C.int(font))
}
func (s *ValueSlider) SetTextSize(size int) {
	C.go_fltk_Value_Slider_set_textsize((*C.GValue_Slider)(s.ptr()), C.int(size))
}
