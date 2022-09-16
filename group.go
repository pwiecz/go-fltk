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
	// Child widgets in an unspecified order
	children          []Widget
	deletionHandlerId uintptr
}

type groupInterface interface {
	Widget
	getGroup() *Group
}

var toplevelGroup *Group = &Group{}
var currentGroup groupInterface = toplevelGroup

func initGroup(g groupInterface, p unsafe.Pointer) {
	initWidget(g, p)
	group := g.getGroup()
	group.deletionHandlerId = group.addDeletionHandler(group.onDelete)
	currentGroup = g
}

func NewGroup(x, y, w, h int, text ...string) *Group {
	g := &Group{}
	initGroup(g, unsafe.Pointer(C.go_fltk_new_Group(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return g
}

func (g *Group) getGroup() *Group {
	return g
}
func (g *Group) Begin() {
	C.go_fltk_Group_begin((*C.GGroup)(g.ptr()))
	currentGroup = g
}
func (g *Group) End() {
	C.go_fltk_Group_end((*C.GGroup)(g.ptr()))
	currentGroup = g.parent
}

func (g *Group) removeChild(child Widget) {
	childWidget := child.getWidget()
	childWidget.parent = nil
	childPtr := childWidget.ptr()
	for i, c := range g.children {
		if c.getWidget().ptr() == childPtr {
			childrenCount := len(g.children)
			g.children[i] = g.children[childrenCount-1]
			g.children[childrenCount-1] = nil
			g.children = g.children[:childrenCount-1]
			return
		}
	}
}

func (g *Group) Add(w Widget) {
	C.go_fltk_Group_add((*C.GGroup)(g.ptr()), w.getWidget().ptr())
	if ww := w.getWidget(); ww.parent != nil {
		ww.parent.getGroup().removeChild(w)
	}
	w.getWidget().parent = g
	g.children = append(g.children, w)
}
func (g *Group) Remove(w Widget) {
	C.go_fltk_Group_remove((*C.GGroup)(g.ptr()), w.getWidget().ptr())
	g.removeChild(w)
	w.getWidget().parent = toplevelGroup
}

func (g *Group) Resizable(w Widget) {
	C.go_fltk_Group_resizable((*C.GGroup)(g.ptr()), w.getWidget().ptr())
}
func (g *Group) DrawChildren() {
	C.go_fltk_Group_draw_children((*C.GGroup)(g.ptr()))
}

func (g *Group) onDelete() {
	if g.deletionHandlerId > 0 {
		globalCallbackMap.unregister(g.deletionHandlerId)
	}
	g.deletionHandlerId = 0
	for i := range g.children {
		g.children[i].getWidget().parent = nil
		g.children[i] = nil
	}
	g.children = g.children[:0]
	if currentGroup == g {
		currentGroup = nil
	}
}
