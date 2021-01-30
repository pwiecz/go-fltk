package fltk

/*
#include "gl_window.h"
*/
import "C"
import "unsafe"

type GlWindow struct {
	Window
	eventHandler int
}

func NewGlWindow(x, y, w, h int, drawFun func()) *GlWindow {
	win := &GlWindow{eventHandler: -1}
	drawFunId := globalCallbackMap.register(drawFun)
	initWidget(win, unsafe.Pointer(C.go_fltk_new_GlWindow(C.int(x), C.int(y), C.int(w), C.int(h), unsafe.Pointer(drawFunId))))
	return win
}

func (w *GlWindow) ContextValid() bool {
	return int(C.go_fltk_Gl_Window_context_valid((*C.Fl_Gl_Window)(w.ptr))) != 0
}
func (w *GlWindow) Valid() bool {
	return int(C.go_fltk_Gl_Window_valid((*C.Fl_Gl_Window)(w.ptr))) != 0
}
func (w *GlWindow) SetEventHandler(handler func(Event) bool) {
	if w.eventHandler >= 0 {
		globalEventHandlerMap.unregister(w.eventHandler)
	}
	w.eventHandler = globalEventHandlerMap.register(handler)
	C.go_fltk_Gl_Window_set_event_handler((*C.Fl_Gl_Window)(w.ptr), C.int(w.eventHandler))
}
