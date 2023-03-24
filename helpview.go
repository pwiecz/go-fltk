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
	hv := &HelpView{}
	initWidget(hv, unsafe.Pointer(C.go_fltk_new_HelpView(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return hv
}

func (h *HelpView) Directory() string {
	return C.GoString(C.go_fltk_HelpView_directory((*C.Fl_Help_View)(h.ptr())))
}

func (h *HelpView) Filename() string {
	return C.GoString(C.go_fltk_HelpView_filename((*C.Fl_Help_View)(h.ptr())))
}

func (h *HelpView) Find(str string, i ...int) int {
	if len(i) < 1 {
		i = append(i, 0)
	}

	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	return int(C.go_fltk_HelpView_find((*C.Fl_Help_View)(h.ptr()), cStr, C.int(i[0])))
}

func (h *HelpView) Load(f string) {
	fStr := C.CString(f)
	defer C.free(unsafe.Pointer(fStr))
	C.go_fltk_HelpView_load((*C.Fl_Help_View)(h.ptr()), fStr)
}

func (h *HelpView) LeftLine() int {
	return int(C.go_fltk_HelpView_leftline((*C.Fl_Help_View)(h.ptr())))
}

func (h *HelpView) SetLeftLine(i int) {
	C.go_fltk_HelpView_set_leftline((*C.Fl_Help_View)(h.ptr()), C.int(i))
}

func (h *HelpView) TopLine() int {
	return int(C.go_fltk_HelpView_topline((*C.Fl_Help_View)(h.ptr())))
}

func (h *HelpView) SetTopLine(i int) {
	C.go_fltk_HelpView_set_topline((*C.Fl_Help_View)(h.ptr()), C.int(i))
}

func (h *HelpView) SetTopLineString(str string) {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	C.go_fltk_HelpView_set_toplinestring((*C.Fl_Help_View)(h.ptr()), cStr)
}

func (h *HelpView) Value() string {
	return C.GoString(C.go_fltk_HelpView_value((*C.Fl_Help_View)(h.ptr())))
}

func (h *HelpView) SetValue(str string) {
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	C.go_fltk_HelpView_set_value((*C.Fl_Help_View)(h.ptr()), cStr)
}

func (h *HelpView) TextSize(size int) {
	C.go_fltk_HelpView_set_textsize((*C.Fl_Help_View)(h.ptr()), C.int(size))
}

func (h *HelpView) TextFont(font Font) {
	C.go_fltk_HelpView_set_textfont((*C.Fl_Help_View)(h.ptr()), C.int(font))
}

func (h *HelpView) TextColor(col Color) {
	C.go_fltk_HelpView_set_textcolor((*C.Fl_Help_View)(h.ptr()), C.uint(col))
}
