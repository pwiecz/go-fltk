package fltk

/*
#include <stdlib.h>
#include "menu.h"
*/
import "C"
import "unsafe"

type menu struct {
	Widget
}

func (m *menu) Add(label string, callback func()) int {
	callbackId := globalCallbackMap.register(callback)
	labelStr := C.CString(label)
	defer C.free(unsafe.Pointer(labelStr))
	return int(C.go_fltk_Menu_add((*C.Fl_Menu_)(m.ptr), labelStr, 0, C.int(callbackId), 0))
}
func (m *menu) SetValue(value int) {
	C.go_fltk_Menu_set_value((*C.Fl_Menu_)(m.ptr), C.int(value))
}
func (m *menu) Value() int {
	return int(C.go_fltk_Menu_value((*C.Fl_Menu_)(m.ptr)))
}
