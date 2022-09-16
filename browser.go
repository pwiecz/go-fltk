package fltk

/*
#include <stdlib.h>
#include "browser.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type Browser struct {
	Group
	icons        map[int]Image
	columnWidths []int
	dataMap      map[uintptr]interface{}
	lastDataID   uintptr
}

var (
	InvalidLine = errors.New("line doesn't exist")
)

func NewBrowser(x, y, w, h int, text ...string) *Browser {
	b := &Browser{}
	b.dataMap = make(map[uintptr]interface{})
	b.icons = make(map[int]Image)
	initGroup(b, unsafe.Pointer(C.go_fltk_new_Browser(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
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

	b.lastDataID++
	id := b.lastDataID
	b.dataMap[id] = data

	C.go_fltk_Browser_add((*C.GBrowser)(b.ptr()), cStr, C.uintptr_t(id))
}

func (b *Browser) TopLine() int {
	return int(C.go_fltk_Browser_topline((*C.GBrowser)(b.ptr())))
}

func (b *Browser) SetBottomLine(line int) error {
	if line < 1 || line > b.Size() {
		return InvalidLine
	}

	C.go_fltk_Browser_set_bottomline((*C.GBrowser)(b.ptr()), C.int(line))
	return nil
}

func (b *Browser) SetMiddleLine(line int) error {
	if line < 1 || line > b.Size() {
		return InvalidLine
	}

	C.go_fltk_Browser_set_middleline((*C.GBrowser)(b.ptr()), C.int(line))
	return nil
}

func (b *Browser) SetTopLine(line int) error {
	if line < 1 || line > b.Size() {
		return InvalidLine
	}

	C.go_fltk_Browser_set_topline((*C.GBrowser)(b.ptr()), C.int(line))
	return nil
}

func (b *Browser) Clear() {
	for k := range b.icons {
		delete(b.icons, k)
	}
	b.dataMap = make(map[uintptr]interface{})
	C.go_fltk_Browser_clear((*C.GBrowser)(b.ptr()))
}

func (b *Browser) Remove(line int) error {
	if line < 1 || line > b.Size() {
		return InvalidLine
	}
	delete(b.icons, line)

	// TODO: got the id from C++ is expensive, need a better way to delete go reference
	id := uintptr(C.go_fltk_Browser_data((*C.GBrowser)(b.ptr()), C.int(line)))
	delete(b.dataMap, id)

	C.go_fltk_Browser_remove((*C.GBrowser)(b.ptr()), C.int(line))
	return nil
}

func (b *Browser) ColumnChar() rune {
	return rune(C.go_fltk_Browser_column_char((*C.GBrowser)(b.ptr())))
}

func (b *Browser) SetColumnChar(r rune) {
	cStr := C.CString(string(r))
	defer C.free(unsafe.Pointer(cStr))

	C.go_fltk_Browser_set_column_char((*C.GBrowser)(b.ptr()), *cStr)
}

func (b *Browser) HideLine(line int) error {
	if line < 1 || line > b.Size() {
		return InvalidLine
	}

	C.go_fltk_Browser_hide_line((*C.GBrowser)(b.ptr()), C.int(line))
	return nil
}

func (b *Browser) Size() int {
	return int(C.go_fltk_Browser_size((*C.GBrowser)(b.ptr())))
}

func (b *Browser) Icon(line int) Image {
	return b.icons[line]
}

func (b *Browser) SetIcon(line int, i Image) {
	if i == nil {
		delete(b.icons, line)
		C.go_fltk_Browser_set_icon((*C.GBrowser)(b.ptr()), C.int(line), nil)
		return
	}

	b.icons[line] = i
	C.go_fltk_Browser_set_icon((*C.GBrowser)(b.ptr()), C.int(line), b.icons[line].getImage().ptr())
}

func (b *Browser) FormatChar() rune {
	return rune(C.go_fltk_Browser_format_char((*C.GBrowser)(b.ptr())))
}

func (b *Browser) SetFormatChar(r rune) {
	cStr := C.CString(string(r))
	defer C.free(unsafe.Pointer(cStr))

	C.go_fltk_Browser_set_format_char((*C.GBrowser)(b.ptr()), *cStr)
}

func (b *Browser) Displayed(line int) bool {
	if C.go_fltk_Browser_displayed((*C.GBrowser)(b.ptr()), C.int(line)) == 1 {
		return true
	}

	return false
}

func (b *Browser) Text(line int) string {
	cStr := C.go_fltk_Browser_text((*C.GBrowser)(b.ptr()), C.int(line))
	//defer C.free(unsafe.Pointer(cStr))

	return C.GoString(cStr)
}

func (b *Browser) SetColumnWidths(widths ...int) {
	cArr := make([]C.int, len(widths), len(widths))
	for i, v := range widths {
		cArr[i] = C.int(v)
	}

	b.columnWidths = widths
	C.go_fltk_Browser_set_column_widths((*C.GBrowser)(b.ptr()), (*C.int)(&cArr[0]))
}

// Store column widths in Go instead of calling from C++ as it's complex and expensive
// to convert between them
func (b *Browser) ColumnWidths() []int {
	return b.columnWidths
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
