package fltk

/*
#include "gl_window.h"
*/
import "C"
import "unsafe"

type GlWindow struct {
	Window
	drawFunId       uintptr
	resizeHandlerId uintptr
	eventHandler    int
}

func NewGlWindow(x, y, w, h int, drawFun func()) *GlWindow {
	win := &GlWindow{eventHandler: -1}
	win.drawFunId = globalCallbackMap.register(drawFun)
	initWidget(win, unsafe.Pointer(C.go_fltk_new_GlWindow(C.int(x), C.int(y), C.int(w), C.int(h), C.uintptr_t(win.drawFunId))))
	return win
}

func (w *GlWindow) Destroy() {
	if w.drawFunId > 0 {
		globalCallbackMap.unregister(w.drawFunId)
	}
	w.drawFunId = 0
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = 0
	if w.eventHandler > 0 {
		globalEventHandlerMap.unregister(w.eventHandler)
	}
	w.eventHandler = 0
	w.Window.Destroy()
}
func (w *GlWindow) ContextValid() bool {
	return int(C.go_fltk_Gl_Window_context_valid((*C.GGlWindow)(w.ptr))) != 0
}
func (w *GlWindow) Valid() bool {
	return int(C.go_fltk_Gl_Window_valid((*C.GGlWindow)(w.ptr))) != 0
}
func (w *GlWindow) SetEventHandler(handler func(Event) bool) {
	if w.eventHandler > 0 {
		globalEventHandlerMap.unregister(w.eventHandler)
	}
	w.eventHandler = globalEventHandlerMap.register(handler)
	C.go_fltk_Gl_Window_set_event_handler((*C.GGlWindow)(w.ptr), C.int(w.eventHandler))
}

func (w *GlWindow) SetResizeHandler(handler func()) {
	if w.resizeHandlerId > 0 {
		globalCallbackMap.unregister(w.resizeHandlerId)
	}
	w.resizeHandlerId = globalCallbackMap.register(handler)
	C.go_fltk_Gl_Window_set_resize_handler((*C.GGlWindow)(w.ptr), C.uintptr_t(w.resizeHandlerId))
}
func (w *GlWindow) SetMode(mode int) {
	C.go_fltk_Gl_Window_set_mode((*C.GGlWindow)(w.ptr), C.int(mode))
}
