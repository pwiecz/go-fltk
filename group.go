package fltk

/*
#include "group.h"
*/
import "C"
import "unsafe"

type Group struct {
	widget
}

func NewGroup(x, y, w, h int, text ...string) *Group {
	g := &Group{}
	initWidget(g, unsafe.Pointer(C.go_fltk_new_Group(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return g
}

func (g *Group) Begin()       { C.go_fltk_Group_begin((*C.GGroup)(g.ptr())) }
func (g *Group) End()         { C.go_fltk_Group_end((*C.GGroup)(g.ptr())) }
func (g *Group) Add(w Widget) { C.go_fltk_Group_add((*C.GGroup)(g.ptr()), w.getWidget().ptr()) }
func (g *Group) Resizable(w Widget) {
	C.go_fltk_Group_resizable((*C.GGroup)(g.ptr()), w.getWidget().ptr())
}
