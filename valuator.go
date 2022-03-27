package fltk

/*
#include "valuator.h"
*/
import "C"

type valuator struct {
	widget
}

func (v *valuator) SetMinimum(value float64) {
	C.go_fltk_Valuator_set_minimum((*C.Fl_Valuator)(v.ptr), C.double(value))
}

func (v *valuator) SetMaximum(value float64) {
	C.go_fltk_Valuator_set_maximum((*C.Fl_Valuator)(v.ptr), C.double(value))
}

func (v *valuator) SetStep(value float64) {
	C.go_fltk_Valuator_set_step((*C.Fl_Valuator)(v.ptr), C.double(value))
}

func (v *valuator) Value() float64 {
	return float64(C.go_fltk_Valuator_value((*C.Fl_Valuator)(v.ptr)))
}

func (v *valuator) SetValue(value float64) {
	C.go_fltk_Valuator_set_value((*C.Fl_Valuator)(v.ptr), C.double(value))
}
