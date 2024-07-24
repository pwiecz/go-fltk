//go:build windows

package fltk

/*
#include "window.h"
*/
import "C"

func (w *Window) RawHandle() uintptr {
	return uintptr(C.go_fltk_Window_win32_xid((*C.Fl_Window)(w.ptr())))
}
