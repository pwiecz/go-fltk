package fltk

/*
#include "scroll.h"
*/
import "C"
import "unsafe"

type Scroll struct {
	Group
}

func NewScroll(x, y, w, h int, text ...string) *Scroll {
	s := &Scroll{}
	initWidget(s, unsafe.Pointer(C.go_fltk_new_Scroll(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return s
}

type ScrollType uint8

var (
	SCROLL_HORIZONTAL        = ScrollType(C.go_FL_SCROLL_HORIZONTAL)
	SCROLL_VERTICAL          = ScrollType(C.go_FL_SCROLL_VERTICAL)
	SCROLL_BOTH              = ScrollType(C.go_FL_SCROLL_BOTH)
	SCROLL_HORIZONTAL_ALWAYS = ScrollType(C.go_FL_SCROLL_HORIZONTAL_ALWAYS)
	SCROLL_VERTICAL_ALWAYS   = ScrollType(C.go_FL_SCROLL_VERTICAL_ALWAYS)
	SCROLL_BOTH_ALWAYS       = ScrollType(C.go_FL_SCROLL_BOTH_ALWAYS)
)

func (s *Scroll) SetType(scrollType ScrollType) {
	s.widget.SetType(uint8(scrollType))
}
