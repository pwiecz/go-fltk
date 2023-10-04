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
	dataMap      *browserDataMap
}

type browserDataMap struct {
	dataMap map[uintptr]interface{}
	id      uintptr
}

func newBrowserDataMap() *browserDataMap {
	return &browserDataMap{
		dataMap: make(map[uintptr]interface{}),
	}
}
func (m *browserDataMap) register(data interface{}) uintptr {
	m.id++
	m.dataMap[m.id] = data
	return m.id
}
func (m *browserDataMap) unregister(id uintptr) {
	delete(m.dataMap, id)
}
func (m *browserDataMap) data(id uintptr) interface{} {
	if id == 0 {
		return nil
	}
	if data, ok := m.dataMap[id]; ok {
		return data
	}
	return nil
}

var (
	ErrInvalidLine = errors.New("line doesn't exist")
)

func NewBrowser(x, y, w, h int, text ...string) *Browser {
	b := &Browser{}
	b.dataMap = newBrowserDataMap()
	b.icons = make(map[int]Image)
	initWidget(b, unsafe.Pointer(C.go_fltk_new_Browser(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return b
}

func (b *Browser) Add(str string) {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))

	C.go_fltk_Browser_add((*C.Fl_Browser)(b.ptr()), cStr, C.uintptr_t(0))
}

func (b *Browser) AddWithData(str string, data interface{}) {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))

	id := b.dataMap.register(data)

	C.go_fltk_Browser_add((*C.Fl_Browser)(b.ptr()), cStr, C.uintptr_t(id))
}

func (b *Browser) TopLine() int {
	return int(C.go_fltk_Browser_topline((*C.Fl_Browser)(b.ptr())))
}

func (b *Browser) SetBottomLine(line int) error {
	if line < 1 || line > b.Size() {
		return ErrInvalidLine
	}

	C.go_fltk_Browser_set_bottomline((*C.Fl_Browser)(b.ptr()), C.int(line))
	return nil
}

func (b *Browser) SetMiddleLine(line int) error {
	if line < 1 || line > b.Size() {
		return ErrInvalidLine
	}

	C.go_fltk_Browser_set_middleline((*C.Fl_Browser)(b.ptr()), C.int(line))
	return nil
}

func (b *Browser) SetTopLine(line int) error {
	if line < 1 || line > b.Size() {
		return ErrInvalidLine
	}

	C.go_fltk_Browser_set_topline((*C.Fl_Browser)(b.ptr()), C.int(line))
	return nil
}

func (b *Browser) Clear() {
	for k := range b.icons {
		delete(b.icons, k)
	}
	b.dataMap = newBrowserDataMap()
	C.go_fltk_Browser_clear((*C.Fl_Browser)(b.ptr()))
}

func (b *Browser) Remove(line int) error {
	if line < 1 || line > b.Size() {
		return ErrInvalidLine
	}
	delete(b.icons, line)

	// TODO: got the id from C++ is expensive, need a better way to delete go reference
	id := uintptr(C.go_fltk_Browser_data((*C.Fl_Browser)(b.ptr()), C.int(line)))
	b.dataMap.unregister(id)

	C.go_fltk_Browser_remove((*C.Fl_Browser)(b.ptr()), C.int(line))
	return nil
}

func (b *Browser) ColumnChar() rune {
	return rune(C.go_fltk_Browser_column_char((*C.Fl_Browser)(b.ptr())))
}

func (b *Browser) SetColumnChar(r rune) {
	cStr := C.CString(string(r))
	defer C.free(unsafe.Pointer(cStr))

	C.go_fltk_Browser_set_column_char((*C.Fl_Browser)(b.ptr()), *cStr)
}

func (b *Browser) HideLine(line int) error {
	if line < 1 || line > b.Size() {
		return ErrInvalidLine
	}

	C.go_fltk_Browser_hide_line((*C.Fl_Browser)(b.ptr()), C.int(line))
	return nil
}

func (b *Browser) Size() int {
	return int(C.go_fltk_Browser_size((*C.Fl_Browser)(b.ptr())))
}

func (b *Browser) Icon(line int) Image {
	return b.icons[line]
}

func (b *Browser) SetIcon(line int, i Image) {
	if i == nil {
		delete(b.icons, line)
		C.go_fltk_Browser_set_icon((*C.Fl_Browser)(b.ptr()), C.int(line), nil)
		return
	}
	b.icons[line] = i
	C.go_fltk_Browser_set_icon((*C.Fl_Browser)(b.ptr()), C.int(line), b.icons[line].getImage().ptr())
}

func (b *Browser) FormatChar() rune {
	return rune(C.go_fltk_Browser_format_char((*C.Fl_Browser)(b.ptr())))
}

func (b *Browser) SetFormatChar(r rune) {
	cStr := C.CString(string(r))
	defer C.free(unsafe.Pointer(cStr))

	C.go_fltk_Browser_set_format_char((*C.Fl_Browser)(b.ptr()), *cStr)
}

func (b *Browser) Displayed(line int) bool {
	return C.go_fltk_Browser_displayed((*C.Fl_Browser)(b.ptr()), C.int(line)) == 1
}

func (b *Browser) Text(line int) string {
	cStr := C.go_fltk_Browser_text((*C.Fl_Browser)(b.ptr()), C.int(line))
	return C.GoString(cStr)
}

func (b *Browser) SetColumnWidths(widths ...int) {
	cArr := make([]C.int, len(widths))
	for i, v := range widths {
		cArr[i] = C.int(v)
	}

	b.columnWidths = widths
	C.go_fltk_Browser_set_column_widths((*C.Fl_Browser)(b.ptr()), (*C.int)(&cArr[0]))
}

// Store column widths in Go instead of calling from C++ as it's complex and expensive
// to convert between them
func (b *Browser) ColumnWidths() []int {
	return b.columnWidths
}

func (b *Browser) Data(line int) interface{} {
	id := uintptr(C.go_fltk_Browser_data((*C.Fl_Browser)(b.ptr()), C.int(line)))
	return b.dataMap.data(id)
}

func (b *Browser) Value() int {
	return int(C.go_fltk_Browser_value((*C.Fl_Browser)(b.ptr())))
}

func (b *Browser) SetValue(line int) {
	C.go_fltk_Browser_set_value((*C.Fl_Browser)(b.ptr()), C.int(line))
}

func (b *Browser) SetSelected(line int, val bool) bool {
	if val {
		return C.go_fltk_Browser_select((*C.Fl_Browser)(b.ptr()), C.int(line), 1) != 0
	} else {
		return C.go_fltk_Browser_select((*C.Fl_Browser)(b.ptr()), C.int(line), 0) != 0
	}
}

func (b *Browser) IsSelected(line int) bool {
	return C.go_fltk_Browser_selected((*C.Fl_Browser)(b.ptr()), C.int(line)) != 0
}

type SelectBrowser struct {
	Browser
}

func NewSelectBrowser(x, y, w, h int, text ...string) *SelectBrowser {
	b := &SelectBrowser{}
	b.dataMap = newBrowserDataMap()
	initWidget(b, unsafe.Pointer(C.go_fltk_new_Select_Browser(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return b
}

type HoldBrowser struct {
	Browser
}

func NewHoldBrowser(x, y, w, h int, text ...string) *HoldBrowser {
	b := &HoldBrowser{}
	b.dataMap = newBrowserDataMap()
	initWidget(b, unsafe.Pointer(C.go_fltk_new_Hold_Browser(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return b
}

type MultiBrowser struct {
	Browser
}

func NewMultiBrowser(x, y, w, h int, text ...string) *MultiBrowser {
	b := &MultiBrowser{}
	b.dataMap = newBrowserDataMap()
	initWidget(b, unsafe.Pointer(C.go_fltk_new_Multi_Browser(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return b
}

type CheckBrowser struct {
	Group
	dataMap *browserDataMap
}

func NewCheckBrowser(x, y, w, h int, text ...string) *CheckBrowser {
	b := &CheckBrowser{}
	b.dataMap = newBrowserDataMap()
	initWidget(b, unsafe.Pointer(C.go_fltk_new_Check_Browser(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return b
}

func (b *CheckBrowser) Add(s string, checked bool) {
	str := C.CString(s)
	defer C.free(unsafe.Pointer(str))
	if checked {
		C.go_fltk_Check_Browser_add((*C.Fl_Check_Browser)(b.ptr()), str, 1)
	} else {
		C.go_fltk_Check_Browser_add((*C.Fl_Check_Browser)(b.ptr()), str, 0)
	}
}

func (b *CheckBrowser) SetChecked(item int, checked bool) {
	if checked {
		C.go_fltk_Check_Browser_set_checked((*C.Fl_Check_Browser)(b.ptr()), C.int(item), 1)
	} else {
		C.go_fltk_Check_Browser_set_checked((*C.Fl_Check_Browser)(b.ptr()), C.int(item), 0)
	}
}

func (b *CheckBrowser) IsChecked(item int) bool {
	return C.go_fltk_Check_Browser_is_checked((*C.Fl_Check_Browser)(b.ptr()), C.int(item)) != 0
}

func (b *CheckBrowser) CheckedCount() int {
	return int(C.go_fltk_Check_Browser_nchecked((*C.Fl_Check_Browser)(b.ptr())))
}

func (b *CheckBrowser) Remove(item int) {
	C.go_fltk_Check_Browser_remove((*C.Fl_Check_Browser)(b.ptr()), C.int(item))
}

func (b *CheckBrowser) Clear() {
	C.go_fltk_Check_Browser_clear((*C.Fl_Check_Browser)(b.ptr()))
}

func (b *CheckBrowser) ItemCount() int {
	return int(C.go_fltk_Check_Browser_nitems((*C.Fl_Check_Browser)(b.ptr())))
}

func (b *CheckBrowser) Text(item int) string {
	return C.GoString(C.go_fltk_Check_Browser_text((*C.Fl_Check_Browser)(b.ptr()), C.int(item)))
}
