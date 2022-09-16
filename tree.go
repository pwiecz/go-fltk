package fltk

/*
#include <stdlib.h>
#include "tree.h"
*/
import "C"
import "unsafe"

type Tree struct {
	Group
}

func NewTree(x, y, w, h int, text ...string) *Tree {
	t := &Tree{}
	initGroup(t, unsafe.Pointer(C.go_fltk_new_Tree(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return t
}

func (t *Tree) SetShowRoot(show bool) {
	if show {
		C.go_fltk_Tree_set_show_root((*C.GTree)(t.ptr()), 1)
	} else {
		C.go_fltk_Tree_set_show_root((*C.GTree)(t.ptr()), 0)
	}
}

type TreeItem struct {
	ptr *C.Fl_Tree_Item
}

func (t *Tree) Add(path string) TreeItem {
	pathStr := C.CString(path)
	defer C.free(unsafe.Pointer(pathStr))
	itemPtr := C.go_fltk_Tree_add((*C.GTree)(t.ptr()), pathStr)
	return TreeItem{ptr: itemPtr}
}
func (t TreeItem) SetWidget(w Widget) {
	C.go_fltk_Tree_Item_set_widget(t.ptr, w.getWidget().ptr())
}

type TreeItemDrawMode uint
var (
	TreeItemDrawDefault = TreeItemDrawMode(C.go_FL_TREE_ITEM_DRAW_DEFAULT)
	TreeItemDrawLabelAndWidget = TreeItemDrawMode(C.go_FL_TREE_ITEM_DRAW_LABEL_AND_WIDGET)
	TreeItemHeightFromWidget = TreeItemDrawMode(C.go_FL_TREE_ITEM_HEIGHT_FROM_WIDGET)
)
func (t *Tree) SetItemDrawMode(drawMode TreeItemDrawMode) {
	C.go_fltk_Tree_set_item_draw_mode((*C.GTree)(t.ptr()), C.uint(drawMode))
}

type TreeConnector int
var (
	TreeConnectorNone = TreeConnector(C.go_FL_TREE_CONNECTOR_NONE)
	TreeConnectorDotted = TreeConnector(C.go_FL_TREE_CONNECTOR_DOTTED)
	TreeConnectorSolid = TreeConnector(C.go_FL_TREE_CONNECTOR_SOLID)
)
func (t *Tree) SetConnectorStyle(style TreeConnector) {
	C.go_fltk_Tree_set_connector_style((*C.GTree)(t.ptr()), C.int(style))
}
