package fltk

/*
#include <stdlib.h>
#include "browser.h"
*/
import "C"
import "unsafe"

type Browser struct {
	widget
	dataMap map[uintptr]interface{}
	lastID  uintptr
}

func NewBrowser(x, y, w, h int, text ...string) *Browser {
	b := &Browser{}
	b.dataMap = make(map[uintptr]interface{})
	initWidget(b, unsafe.Pointer(C.go_fltk_new_Browser(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return b
}

func (b *Browser) Add(str string) {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))

	C.go_fltk_Browser_add((*C.GBrowser)(b.ptr()), cStr, C.uintptr_t(0))
}

func (b *Browser) AddWithData(str string, data interface{}) {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))

	b.lastID++
	id := b.lastID
	b.dataMap[id] = data

	C.go_fltk_Browser_add((*C.GBrowser)(b.ptr()), cStr, C.uintptr_t(id))
}

func (b *Browser) BottomLine(i int) { C.go_fltk_Browser_bottomline((*C.GBrowser)(b.ptr()), C.int(i)) }

func (b *Browser) Clear() {
	b.dataMap = make(map[uintptr]interface{})
	C.go_fltk_Browser_clear((*C.GBrowser)(b.ptr()))
}

func (b *Browser) Data(line int) interface{} {
	id := uintptr(C.go_fltk_Browser_data((*C.GBrowser)(b.ptr()), C.int(line)))
	return b.dataMap[id]
}

func (b *Browser) Value() int {
	return int(C.go_fltk_Browser_value((*C.GBrowser)(b.ptr())))
}

func (b *Browser) SetValue(line int) {
	C.go_fltk_Browser_set_value((*C.GBrowser)(b.ptr()), C.int(line))
}

type SelectBrowser struct {
	Browser
}

func NewSelectBrowser(x, y, w, h int, text ...string) *SelectBrowser {
	b := &SelectBrowser{}
	b.dataMap = make(map[uintptr]interface{})
	initWidget(b, unsafe.Pointer(C.go_fltk_new_Select_Browser(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return b
}

type HoldBrowser struct {
	Browser
}

func NewHoldBrowser(x, y, w, h int, text ...string) *HoldBrowser {
	b := &HoldBrowser{}
	b.dataMap = make(map[uintptr]interface{})
	initWidget(b, unsafe.Pointer(C.go_fltk_new_Hold_Browser(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return b
}

type MultiBrowser struct {
	Browser
}

func NewMultiBrowser(x, y, w, h int, text ...string) *MultiBrowser {
	b := &MultiBrowser{}
	b.dataMap = make(map[uintptr]interface{})
	initWidget(b, unsafe.Pointer(C.go_fltk_new_Multi_Browser(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return b
}
