package fltk

/*
#include <stdlib.h>
#include "menu.h"
*/
import "C"
import "unsafe"

type menu struct {
	widget
	deletionHandlerId uintptr
	itemCallbacks     []uintptr
}

func (m *menu) init() {
	if m.deletionHandlerId > 0 {
		panic("menu already initialized")
	}
	m.deletionHandlerId = m.addDeletionHandler(m.onDelete)
}
func (m *menu) onDelete() {
	if m.deletionHandlerId > 0 {
		globalCallbackMap.unregister(m.deletionHandlerId)
	}
	m.deletionHandlerId = 0
	for _, itemCallbackId := range m.itemCallbacks {
		globalCallbackMap.unregister(itemCallbackId)
	}
	m.itemCallbacks = m.itemCallbacks[:0]
}
func (m *menu) Destroy() {
	for _, itemCallbackId := range m.itemCallbacks {
		globalCallbackMap.unregister(itemCallbackId)
	}
	m.itemCallbacks = m.itemCallbacks[:0]
	m.widget.Destroy()
}
func (m *menu) Add(label string, callback func()) int {
	callbackId := globalCallbackMap.register(callback)
	m.itemCallbacks = append(m.itemCallbacks, callbackId)
	labelStr := C.CString(label)
	defer C.free(unsafe.Pointer(labelStr))
	return int(C.go_fltk_Menu_add((*C.Fl_Menu_)(m.ptr()), labelStr, 0, C.int(callbackId), 0))
}
func (m *menu) AddEx(label string, shortcut int, callback func(), flags int) int {
	callbackId := globalCallbackMap.register(callback)
	m.itemCallbacks = append(m.itemCallbacks, callbackId)
	labelStr := C.CString(label)
	defer C.free(unsafe.Pointer(labelStr))
	return int(C.go_fltk_Menu_add((*C.Fl_Menu_)(m.ptr()), labelStr, C.int(shortcut), C.int(callbackId), C.int(flags)))
}
func (m *menu) Insert(index int, label string, callback func()) int {
	callbackId := globalCallbackMap.register(callback)
	m.itemCallbacks = append(m.itemCallbacks, callbackId)
	labelStr := C.CString(label)
	defer C.free(unsafe.Pointer(labelStr))
	return int(C.go_fltk_Menu_insert((*C.Fl_Menu_)(m.ptr()), C.int(index), labelStr, 0, C.int(callbackId), 0))
}
func (m *menu) InsertEx(index int, label string, shortcut int, callback func(), flags int) int {
	callbackId := globalCallbackMap.register(callback)
	m.itemCallbacks = append(m.itemCallbacks, callbackId)
	labelStr := C.CString(label)
	defer C.free(unsafe.Pointer(labelStr))
	return int(C.go_fltk_Menu_insert((*C.Fl_Menu_)(m.ptr()), C.int(index), labelStr, C.int(shortcut), C.int(callbackId), C.int(flags)))
}
func (m *menu) Remove(index int) {
	C.go_fltk_Menu_remove((*C.Fl_Menu_)(m.ptr()), C.int(index))
}
func (m *menu) Replace(index int, label string) {
	labelStr := C.CString(label)
	defer C.free(unsafe.Pointer(labelStr))
	C.go_fltk_Menu_replace((*C.Fl_Menu_)(m.ptr()), C.int(index), labelStr)
}
func (m *menu) FindIndex(label string) int {
	labelStr := C.CString(label)
	defer C.free(unsafe.Pointer(labelStr))
	return int(C.go_fltk_Menu_find_index((*C.Fl_Menu_)(m.ptr()), labelStr))
}
func (m *menu) SetValue(value int) {
	C.go_fltk_Menu_set_value((*C.Fl_Menu_)(m.ptr()), C.int(value))
}
func (m *menu) Value() int {
	return int(C.go_fltk_Menu_value((*C.Fl_Menu_)(m.ptr())))
}
func (m *menu) Size() int {
	return int(C.go_fltk_Menu_size((*C.Fl_Menu_)(m.ptr())))
}

type MenuButton struct {
	menu
}

func NewMenuButton(x, y, w, h int, text ...string) *MenuButton {
	m := &MenuButton{}
	initWidget(m, unsafe.Pointer(C.go_fltk_new_MenuButton(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	m.menu.init()
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

var (
	MENU_INACTIVE  = int(C.go_FL_MENU_INACTIVE)
	MENU_TOGGLE    = int(C.go_FL_MENU_TOGGLE)
	MENU_VALUE     = int(C.go_FL_MENU_VALUE)
	MENU_RADIO     = int(C.go_FL_MENU_RADIO)
	MENU_INVISIBLE = int(C.go_FL_MENU_INVISIBLE)
	SUBMENU        = int(C.go_FL_SUBMENU)
	MENU_DIVIDER   = int(C.go_FL_MENU_DIVIDER)
)

func (m *MenuButton) SetType(menuType MenuType) {
	C.go_fltk_MenuButton_set_type((*C.GMenu_Button)(m.ptr()), C.int(menuType))
}
func (m *MenuButton) Popup() {
	C.go_fltk_MenuButton_popup((*C.GMenu_Button)(m.ptr()))
}
func (m *MenuButton) Destroy() {
	m.menu.Destroy()
}

type MenuBar struct {
	menu
}

func NewMenuBar(x, y, w, h int, text ...string) *MenuBar {
	m := &MenuBar{}
	initWidget(m, unsafe.Pointer(C.go_fltk_new_MenuBar(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	m.menu.init()
	return m
}
