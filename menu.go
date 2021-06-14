package fltk

/*
#include <stdlib.h>
#include "menu.h"
*/
import "C"
import "unsafe"

type menu struct {
	widget
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
func (m *menu) Size() int {
	return int(C.go_fltk_Menu_size((*C.Fl_Menu_)(m.ptr)))
}

type MenuButton struct {
	menu
	itemCallbacks     []uintptr
	destroyCallbackId uintptr
}

func NewMenuButton(x, y, w, h int, text ...string) *MenuButton {
	m := &MenuButton{}
	initWidget(m, unsafe.Pointer(C.go_fltk_new_MenuButton(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return m
}

type MenuType int

var (
	POPUP1   = MenuType(C.go_FL_POPUP1)
	POPUP2   = MenuType(C.go_FL_POPUP2)
	POPUP12  = MenuType(C.go_FL_POPUP12)
	POPUP3   = MenuType(C.go_FL_POPUP3)
	POPUP13  = MenuType(C.go_FL_POPUP13)
	POPUP23  = MenuType(C.go_FL_POPUP23)
	POPUP123 = MenuType(C.go_FL_POPUP123)
)

func (m *MenuButton) SetType(menuType MenuType) {
	C.go_fltk_MenuButton_set_type((*C.Fl_Menu_Button)(m.ptr), C.int(menuType))
}
func (m *MenuButton) Popup() {
	C.go_fltk_MenuButton_popup((*C.Fl_Menu_Button)(m.ptr))
}
func (m *MenuButton) Destroy() {
	for _, itemCallbackId := range m.itemCallbacks {
		globalCallbackMap.unregister(itemCallbackId)
	}
	m.itemCallbacks = m.itemCallbacks[:0]
	m.menu.Destroy()
}
func (m *MenuButton) Add(label string, callback func()) int {
	callbackId := globalCallbackMap.register(callback)
	m.itemCallbacks = append(m.itemCallbacks, callbackId)
	labelStr := C.CString(label)
	defer C.free(unsafe.Pointer(labelStr))
	return int(C.go_fltk_Menu_add((*C.Fl_Menu_)(m.ptr), labelStr, 0, C.int(callbackId), 0))
}
