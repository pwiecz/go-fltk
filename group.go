package fltk

/*
#include "group.h"
*/
import "C"
import (
	"unsafe"
)

type Group struct {
	widget
}

func NewGroup(x, y, w, h int, text ...string) *Group {
	g := &Group{}
	initWidget(g, unsafe.Pointer(C.go_fltk_new_Group(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return g
}

func (g *Group) Begin() {
	C.go_fltk_Group_begin((*C.Fl_Group)(g.ptr()))
}
func (g *Group) End() {
	C.go_fltk_Group_end((*C.Fl_Group)(g.ptr()))
}

func (g *Group) Add(w Widget) {
	C.go_fltk_Group_add((*C.Fl_Group)(g.ptr()), w.getWidget().ptr())
}
func (g *Group) Remove(w Widget) {
	C.go_fltk_Group_remove((*C.Fl_Group)(g.ptr()), w.getWidget().ptr())
}

func (g *Group) Resizable(w Widget) {
	C.go_fltk_Group_resizable((*C.Fl_Group)(g.ptr()), w.getWidget().ptr())
}
func (g *Group) DrawChildren() {
	C.go_fltk_Group_draw_children((*C.Fl_Group)(g.ptr()))
}

func (g *Group) Child(index int) *widget {
	child := C.go_fltk_Group_child((*C.Fl_Group)(g.ptr()), C.int(index))
	if child == nil {
		return nil
	}
	widget := &widget{}
	initUnownedWidget(widget, unsafe.Pointer(child))
	return widget
}
