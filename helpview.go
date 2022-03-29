package fltk

/*
#include <stdlib.h>
#include "helpview.h"
*/
import "C"
import "unsafe"

type HelpView struct {
	widget
}

func NewHelpView(x, y, w, h int, text ...string) *HelpView {
	t := &HelpView{}
	initWidget(t, unsafe.Pointer(C.go_fltk_new_HelpView(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return t
}

func (h *HelpView) Destroy() {
	if h.ptr != nil {
		C.go_fltk_HelpView_delete((*C.Fl_Help_View)(h.ptr))
	}
}

func (h *HelpView) Directory() string {
	return C.GoString(C.go_fltk_HelpView_directory((*C.Fl_Help_View)(h.ptr)))
}

func (h *HelpView) Filename() string {
	return C.GoString(C.go_fltk_HelpView_filename((*C.Fl_Help_View)(h.ptr)))
}

func (h *HelpView) Find(str string, i ...int) int {
	if len(i) < 1 {
		i = append(i, 0)
	}

	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	return int(C.go_fltk_HelpView_find((*C.Fl_Help_View)(h.ptr), cStr, C.int(i[0])))
}

func (h *HelpView) Load(f string) {
	fStr := C.CString(f)
	defer C.free(unsafe.Pointer(fStr))
	C.go_fltk_HelpView_load((*C.Fl_Help_View)(h.ptr), fStr)
}
