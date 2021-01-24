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

type glWindowEventCallbackMap struct {
	callbackMap map[int]func(Event) bool
	id          int
}

func newGlWindowEventCallbackMap() *glWindowEventCallbackMap {
	return &glWindowEventCallbackMap{
		callbackMap: make(map[int]func(Event) bool),
	}
}
func (m *glWindowEventCallbackMap) register(fn func(Event) bool) int {
	m.id++
	m.callbackMap[m.id] = fn
	return m.id
}
func (m *glWindowEventCallbackMap) unregister(id int) {
	delete(m.callbackMap, id)
}
func (m *glWindowEventCallbackMap) invoke(id int, event Event) bool {
	return m.callbackMap[id](event)
}

var globalGlWindowEventCallbackMap = newGlWindowEventCallbackMap()

func NewGlWindow(x, y, w, h int, drawFun func()) *GlWindow {
	win := &GlWindow{eventHandler: -1}
	drawFunId := globalCallbackMap.register(drawFun)
	initWidget(win, unsafe.Pointer(C.go_fltk_new_GlWindow(C.int(x), C.int(y), C.int(w), C.int(h), unsafe.Pointer(drawFunId))))
	return win
}

//export _go_glWindowEventHandler
func _go_glWindowEventHandler(handlerId C.int, event C.int) C.int {
	if globalGlWindowEventCallbackMap.invoke(int(handlerId), Event(event)) {
		return 1
	}
	return 0
}

func (w *GlWindow) ContextValid() bool {
	return int(C.go_fltk_Gl_Window_context_valid((*C.Fl_Gl_Window)(w.ptr))) != 0
}
func (w *GlWindow) Valid() bool {
	return int(C.go_fltk_Gl_Window_valid((*C.Fl_Gl_Window)(w.ptr))) != 0
}
func (w *GlWindow) SetEventHandler(handler func(Event) bool) {
	if w.eventHandler >= 0 {
		globalGlWindowEventCallbackMap.unregister(w.eventHandler)
	}
	w.eventHandler = globalGlWindowEventCallbackMap.register(handler)
	C.go_fltk_Gl_Window_set_event_handler((*C.Fl_Gl_Window)(w.ptr), C.int(w.eventHandler))
}
