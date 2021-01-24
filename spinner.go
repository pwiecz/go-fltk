package fltk

/*
#include "enumerations.h"
#include "input.h"
#include "spinner.h"
*/
import "C"
import "unsafe"

type Spinner struct {
	Widget
}

func NewSpinner(x, y, w, h int, text ...string) *Spinner {
	s := &Spinner{}
	initWidget(s, unsafe.Pointer(C.go_fltk_new_Spinner(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return s
}

type SpinnerInputType uint8

var SPINNER_INT_INPUT = SpinnerInputType(C.go_FL_INT_INPUT)
var SPINNER_FLOAT_INPUT = SpinnerInputType(C.go_FL_FLOAT_INPUT)

func (s *Spinner) SetType(inputType SpinnerInputType) {
	C.go_fltk_Spinner_set_type((*C.Fl_Spinner)(s.ptr), C.uchar(inputType))
}
func (s *Spinner) SetMaximum(max float64) {
	C.go_fltk_Spinner_set_maximum((*C.Fl_Spinner)(s.ptr), C.double(max))
}
func (s *Spinner) SetMinimum(min float64) {
	C.go_fltk_Spinner_set_minimum((*C.Fl_Spinner)(s.ptr), C.double(min))
}
func (s *Spinner) SetStep(step float64) {
	C.go_fltk_Spinner_set_step((*C.Fl_Spinner)(s.ptr), C.double(step))
}
func (s *Spinner) SetValue(val float64) {
	C.go_fltk_Spinner_set_value((*C.Fl_Spinner)(s.ptr), C.double(val))
}
func (s *Spinner) Value() float64 {
	return (float64)(C.go_fltk_Spinner_value((*C.Fl_Spinner)(s.ptr)))
}
