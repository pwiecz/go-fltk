package fltk

/*
#include "gl_window.h"
*/
import "C"
import (
	"unsafe"
)

type GlWindow struct {
	Window
	deletionHandlerId uintptr
	drawFunId         uintptr
}

func NewGlWindow(x, y, w, h int, drawFun func()) *GlWindow {
	win := &GlWindow{}
	win.drawFunId = globalCallbackMap.register(drawFun)
	initWidget(win, unsafe.Pointer(C.go_fltk_new_GlWindow(C.int(x), C.int(y), C.int(w), C.int(h), C.uintptr_t(win.drawFunId))))
	win.deletionHandlerId = win.addDeletionHandler(win.onDelete)
	return win
}

func (w *GlWindow) MakeCurrent() {
	C.go_fltk_Gl_Window_make_current((*C.GGlWindow)(w.ptr()))
}
func (w *GlWindow) onDelete() {
	if w.deletionHandlerId > 0 {
		globalCallbackMap.unregister(w.deletionHandlerId)
	}
	w.deletionHandlerId = 0
	if w.drawFunId > 0 {
		globalCallbackMap.unregister(w.drawFunId)
	}
	w.drawFunId = 0

}
func (w *GlWindow) Destroy() {
	if w.drawFunId > 0 {
		globalCallbackMap.unregister(w.drawFunId)
	}
	w.drawFunId = 0
	w.Window.Destroy()
}
func (w *GlWindow) ContextValid() bool {
	return C.go_fltk_Gl_Window_context_valid((*C.GGlWindow)(w.ptr())) != 0
}
func (w *GlWindow) Valid() bool {
	return C.go_fltk_Gl_Window_valid((*C.GGlWindow)(w.ptr())) != 0
}
func (w *GlWindow) CanDo() bool {
	return C.go_fltk_Gl_Window_can_do((*C.GGlWindow)(w.ptr())) != 0
}
func (w *GlWindow) SetMode(mode int) {
	C.go_fltk_Gl_Window_set_mode((*C.GGlWindow)(w.ptr()), C.int(mode))
}
