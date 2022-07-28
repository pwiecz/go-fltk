package fltk

/*
#include <stdlib.h>
#include "browser.h"
*/
import "C"
import "unsafe"
import "fmt"

type Browser struct {
	widget
	icons  map[int]Image
}

func NewBrowser(x, y, w, h int, text ...string) *Browser {
	b := &Browser{}
	b.icons = make(map[int]Image)
	initWidget(b, unsafe.Pointer(C.go_fltk_new_Browser(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return b
}

func (b *Browser) Add(str string) {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))

	C.go_fltk_Browser_add((*C.GBrowser)(b.ptr()), cStr, unsafe.Pointer(&cStr))
}

func (b *Browser) BottomLine(i int) {
	C.go_fltk_Browser_bottomline((*C.GBrowser)(b.ptr()), C.int(i))
}

func (b *Browser) Clear() {
	C.go_fltk_Browser_clear((*C.GBrowser)(b.ptr()))
}

func (b *Browser) Remove(i int) {
	C.go_fltk_Browser_remove((*C.GBrowser)(b.ptr()), C.int(i))
}

func (b *Browser) ColumnChar() rune {
	return rune(C.go_fltk_Browser_column_char((*C.GBrowser)(b.ptr())))
}

func (b *Browser) SetColumnChar(r rune) {
	cStr := C.CString(string(r))
	defer C.free(unsafe.Pointer(cStr))

	C.go_fltk_Browser_set_column_char((*C.GBrowser)(b.ptr()), *cStr)
}

//TODO: fix crash whenever this is called
/*
func (b *Browser) HideLine(line int) {
	C.go_fltk_Browser_hide_line((*C.GBrowser)(b.ptr()), C.int(line))
}
*/

func (b *Browser) Icon(line int) Image {
		fmt.Println(b.icons)
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
