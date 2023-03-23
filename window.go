package fltk

/*
#include <stdlib.h>
#include "window.h"
*/
import "C"
import "unsafe"

type Window struct {
	Group
	icons []*RgbImage
}

func NewWindow(w, h int) *Window {
	win := &Window{}
	initWidget(win, unsafe.Pointer(C.go_fltk_new_Window(C.int(w), C.int(h))))
	return win
}
func (w *Window) IsShown() bool {
	return C.go_fltk_Window_shown((*C.Fl_Window)(w.ptr())) != 0
}
func (w *Window) Show() { C.go_fltk_Window_show((*C.Fl_Window)(w.ptr())) }
func (w *Window) XRoot() int {
	return int(C.go_fltk_Window_x_root((*C.Fl_Window)(w.ptr())))
}
func (w *Window) YRoot() int {
	return int(C.go_fltk_Window_y_root((*C.Fl_Window)(w.ptr())))
}
func (w *Window) SetLabel(label string) {
	labelStr := C.CString(label)
	defer C.free(unsafe.Pointer(labelStr))
	C.go_fltk_Window_set_label((*C.Fl_Window)(w.ptr()), labelStr)
}
func (w *Window) SetCursor(cursor Cursor) {
	C.go_fltk_Window_set_cursor((*C.Fl_Window)(w.ptr()), C.int(cursor))
}

func (w *Window) SetFullscreen(flag bool) {
	f := 0
	if flag {
		f = 1
	}
	C.go_fltk_Window_set_fullscreen((*C.Fl_Window)(w.ptr()), C.int(f))
}

func (w *Window) FullscreenActive() bool {
	return C.go_fltk_Window_fullscreen_active((*C.Fl_Window)(w.ptr())) != 0
}

func (w *Window) SetModal() {
	C.go_fltk_Window_set_modal((*C.Fl_Window)(w.ptr()))
}

func (w *Window) SetNonModal() {
	C.go_fltk_Window_set_non_modal((*C.Fl_Window)(w.ptr()))
}

func (w *Window) SetIcons(icons []*RgbImage) {
	images := make([]*C.Fl_RGB_Image, 0, len(icons))
	for _, icon := range icons {
		images = append(images, (*C.Fl_RGB_Image)(icon.iPtr))
	}
	C.go_fltk_Window_set_icons((*C.Fl_Window)(w.ptr()), &images[0], C.int(len(images)))
	w.icons = icons
}

type Cursor int

var (
	CURSOR_DEFAULT = Cursor(C.go_FL_CURSOR_DEFAULT)
	CURSOR_ARROW   = Cursor(C.go_FL_CURSOR_ARROW)
	CURSOR_CROSS   = Cursor(C.go_FL_CURSOR_CROSS)
	CURSOR_WAIT    = Cursor(C.go_FL_CURSOR_WAIT)
	CURSOR_INSERT  = Cursor(C.go_FL_CURSOR_INSERT)
	CURSOR_HAND    = Cursor(C.go_FL_CURSOR_HAND)
	CURSOR_HELP    = Cursor(C.go_FL_CURSOR_HELP)
	CURSOR_MOVE    = Cursor(C.go_FL_CURSOR_MOVE)
	CURSOR_NS      = Cursor(C.go_FL_CURSOR_NS)
	CURSOR_WE      = Cursor(C.go_FL_CURSOR_WE)
	CURSOR_NWSE    = Cursor(C.go_FL_CURSOR_NWSE)
	CURSOR_NESW    = Cursor(C.go_FL_CURSOR_NESW)
	CURSOR_N       = Cursor(C.go_FL_CURSOR_N)
	CURSOR_NE      = Cursor(C.go_FL_CURSOR_NE)
	CURSOR_E       = Cursor(C.go_FL_CURSOR_E)
	CURSOR_SE      = Cursor(C.go_FL_CURSOR_SE)
	CURSOR_S       = Cursor(C.go_FL_CURSOR_S)
	CURSOR_SW      = Cursor(C.go_FL_CURSOR_SW)
	CURSOR_W       = Cursor(C.go_FL_CURSOR_W)
	CURSOR_NW      = Cursor(C.go_FL_CURSOR_NW)
)
