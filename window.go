package fltk

/*
#include <stdlib.h>
#include "window.h"
*/
import "C"
import "unsafe"

type Window struct {
	Group
}

func NewWindow(w, h int) *Window {
	win := &Window{}
	initWidget(win, unsafe.Pointer(C.go_fltk_new_Window(C.int(w), C.int(h))))
	return win
}

//func (w *Window) Destroy() {C.go_fltk_Window_destroy(w.ptr)}
func (w *Window) Show() { C.go_fltk_Window_show((*C.Fl_Window)(w.ptr)) }
func (w *Window) SetLabel(label string) {
	labelStr := C.CString(label)
	defer C.free(unsafe.Pointer(labelStr))
	C.go_fltk_Window_set_label((*C.Fl_Window)(w.ptr), labelStr)
}
