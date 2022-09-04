package fltk

/*
#include "slider.h"
*/
import "C"
import "unsafe"

type Slider struct {
	valuator
}

type SliderType uint8

var (
	VERT_SLIDER      = SliderType(C.go_FL_VERT_SLIDER)
	HOR_SLIDER       = SliderType(C.go_FL_HOR_SLIDER)
	VERT_FILL_SLIDER = SliderType(C.go_FL_VERT_FILL_SLIDER)
	HOR_FILL_SLIDER  = SliderType(C.go_FL_HOR_FILL_SLIDER)
	VERT_NICE_SLIDER = SliderType(C.go_FL_VERT_NICE_SLIDER)
	HOR_NICE_SLIDER  = SliderType(C.go_FL_HOR_NICE_SLIDER)
)

func NewSlider(x, y, w, h int, text ...string) *Slider {
	s := &Slider{}
	initWidget(s, unsafe.Pointer(C.go_fltk_new_Slider(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return s
}

func (s *Slider) SetType(sliderType SliderType) {
	s.widget.SetType(uint8(sliderType))
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
