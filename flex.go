package fltk

/*
#include "flex.h"
*/
import "C"
import "unsafe"

type Flex struct {
	Group
}

func NewFlex(x, y, w, h int, text ...string) *Flex {
	p := &Flex{}
	initGroup(p, unsafe.Pointer(C.go_fltk_new_Flex(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return p
}

type FlexType uint8

var (
	COLUMN = FlexType(C.go_FL_FLEX_COLUMN)
	ROW    = FlexType(C.go_FL_FLEX_ROW)
)

func (f *Flex) SetType(flexType FlexType) {
	f.widget.SetType(uint8(flexType))
}
func (f *Flex) SetGap(spacing int) {
	C.go_fltk_Flex_set_gap((*C.GFlex)(f.ptr()), C.int(spacing))
}
func (f *Flex) Fixed(w Widget, size int) {
	C.go_fltk_Flex_fixed((*C.GFlex)(f.ptr()), w.getWidget().ptr(), C.int(size))
}
func (f *Flex) End() {
	C.go_fltk_Flex_end((*C.GFlex)(f.ptr()))
}
