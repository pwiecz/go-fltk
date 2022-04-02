package fltk

/*
#include "box.h"
*/
import "C"
import "unsafe"

type Box struct {
	widget
	eventHandler int
}

func NewBox(boxType BoxType, x, y, w, h int, text ...string) *Box {
	b := &Box{}
	initWidget(b, unsafe.Pointer(C.go_fltk_new_Box(C.int(boxType), C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return b
}
func (b *Box) SetEventHandler(handler func(Event) bool) {
	if b.eventHandler > 0 {
		globalEventHandlerMap.unregister(b.eventHandler)
	}
	b.eventHandler = globalEventHandlerMap.register(handler)
	C.go_fltk_Box_set_event_handler((*C.GBox)(b.ptr), C.int(b.eventHandler))
}
