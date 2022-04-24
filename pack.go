package fltk

/*
#include "pack.h"
*/
import "C"
import "unsafe"

type Pack struct {
	Group
}

func NewPack(x, y, w, h int, text ...string) *Pack {
	p := &Pack{}
	initWidget(p, unsafe.Pointer(C.go_fltk_new_Pack(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return p
}

type PackType uint8

var (
	VERTICAL   = PackType(C.go_FL_PACK_VERTICAL)
	HORIZONTAL = PackType(C.go_FL_PACK_HORIZONTAL)
)

func (p *Pack) SetType(packType PackType) {
	p.widget.SetType(uint8(packType))
}
func (p *Pack) SetSpacing(spacing int) {
	C.go_fltk_Pack_set_spacing((*C.GPack)(p.ptr()), C.int(spacing))
}
