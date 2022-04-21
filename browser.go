package fltk

/*
#include <stdlib.h>
#include "browser.h"
*/
import "C"
import "unsafe"

type Browser struct {
	widget
}

func NewBrowser(x, y, w, h int, text ...string) *Browser {
	b := &Browser{}
	initWidget(b, unsafe.Pointer(C.go_fltk_new_Browser(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return b
}

func (b *Browser) Add(str string) {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))

	C.go_fltk_Browser_add((*C.GBrowser)(b.ptr), cStr, unsafe.Pointer(&cStr))
}

func (b *Browser) BottomLine(i int) {
	C.go_fltk_Browser_bottomline((*C.GBrowser)(b.ptr), C.int(i))
}

func (b *Browser) Clear() {
	C.go_fltk_Browser_clear((*C.GBrowser)(b.ptr))
}
