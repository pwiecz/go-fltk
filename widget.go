package fltk

/*
#include <stdlib.h>
#include "widget.h"
*/
import "C"
import "unsafe"

type widget struct {
	ptr        *C.Fl_Widget
	callbackId uintptr
	eventHandlerId int
}

type Widget interface {
	getWidget() *widget
}

func initWidget(w Widget, p unsafe.Pointer) {
	w.getWidget().ptr = (*C.Fl_Widget)(p)
}

func (w *widget) SetCallback(f func()) {
	if w.callbackId > 0 {
		globalCallbackMap.unregister(w.callbackId)
	}
	w.callbackId = globalCallbackMap.register(f)
	C.go_fltk_Widget_set_callback(w.ptr, C.uintptr_t(w.callbackId))
}
func (w *widget) SetCallbackCondition(when CallbackCondition) {
	C.go_fltk_Widget_when(w.ptr, C.int(when))
}
func (w *widget) SetEventHandler(handler func(Event) bool) {
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = globalEventHandlerMap.register(handler)
	if C.go_fltk_Widget_set_event_handler(w.ptr, C.int(w.eventHandlerId)) == 0 {
		panic("This widget does not support event handling")
	}
}

func (w *widget) getWidget() *widget { return w }
func (w *widget) Destroy() {
	if w.callbackId > 0 {
		globalCallbackMap.unregister(w.callbackId)
	}
	w.callbackId = 0
	if w.eventHandlerId > 0 {
		globalEventHandlerMap.unregister(w.eventHandlerId)
	}
	w.eventHandlerId = 0
	if w.ptr != nil {
		C.go_fltk_delete_widget(w.ptr)
	}
	w.ptr = nil
}

func (w *widget) SetBox(box BoxType)           { C.go_fltk_Widget_set_box(w.ptr, C.int(box)) }
func (w *widget) SetLabelFont(font Font)       { C.go_fltk_Widget_set_labelfont(w.ptr, C.int(font)) }
func (w *widget) SetLabelSize(size int)        { C.go_fltk_Widget_set_labelsize(w.ptr, C.int(size)) }
func (w *widget) SetLabelType(ltype LabelType) { C.go_fltk_Widget_set_labeltype(w.ptr, C.int(ltype)) }
func (w *widget) SetLabelColor(col Color)      { C.go_fltk_Widget_set_labelcolor(w.ptr, C.uint(col)) }
func (w *widget) ClearVisibleFocus()           { C.go_fltk_Widget_clear_visible_focus(w.ptr) }
func (w *widget) X() int                       { return int(C.go_fltk_Widget_x(w.ptr)) }
func (w *widget) Y() int                       { return int(C.go_fltk_Widget_y(w.ptr)) }
func (w *widget) W() int                       { return int(C.go_fltk_Widget_w(w.ptr)) }
func (w *widget) H() int                       { return int(C.go_fltk_Widget_h(w.ptr)) }
func (w *widget) SetAlign(align Align)         { C.go_fltk_Widget_set_align(w.ptr, C.uint(align)) }
func (w *widget) MeasureLabel() (int, int) {
	var width, height C.int
	C.go_fltk_Widget_measure_label(w.ptr, &width, &height)
	return int(width), int(height)
}
func (w *widget) SetPosition(x, y int) { C.go_fltk_Widget_set_position(w.ptr, C.int(x), C.int(y)) }
func (w *widget) Resize(x, y, width, height int) {
	C.go_fltk_Widget_resize(w.ptr, C.int(x), C.int(y), C.int(width), C.int(height))
}
func (w *widget) Redraw()                  { C.go_fltk_Widget_redraw(w.ptr) }
func (w *widget) Deactivate()              { C.go_fltk_Widget_deactivate(w.ptr) }
func (w *widget) Activate()                { C.go_fltk_Widget_activate(w.ptr) }
func (w *widget) SetType(widgetType uint8) { C.go_fltk_Widget_set_type(w.ptr, C.uchar(widgetType)) }
func (w *widget) Show()                    { C.go_fltk_Widget_show(w.ptr) }
func (w *widget) Hide()                    { C.go_fltk_Widget_hide(w.ptr) }
func (w *widget) SelectionColor() Color    { return Color(C.go_fltk_Widget_selection_color(w.ptr)) }
func (w *widget) SetSelectionColor(color Color) {
	C.go_fltk_Widget_set_selection_color(w.ptr, C.uint(color))
}
func (w *widget) SetColor(color Color) {
	C.go_fltk_Widget_set_color(w.ptr, C.uint(color))
}
func (w *widget) SetLabel(label string) {
	labelStr := C.CString(label)
	defer C.free(unsafe.Pointer(labelStr))
	C.go_fltk_Widget_set_label(w.ptr, labelStr)
}
func (w *widget) SetImage(i Image)   { C.go_fltk_Widget_set_image(w.ptr, i.getImage().ptr) }
func (w *widget) SetDeimage(i Image) { C.go_fltk_Widget_set_deimage(w.ptr, i.getImage().ptr) }
func (w *widget) Box() BoxType       { return BoxType(C.go_fltk_Widget_box(w.ptr)) }
func (w *widget) LabelColor() Color  { return Color(C.go_fltk_Widget_labelcolor(w.ptr)) }
func (w *widget) Align() Align       { return Align(C.go_fltk_Widget_align(w.ptr)) }
func (w *widget) Type() uint8        { return uint8(C.go_fltk_Widget_type(w.ptr)) }
func (w *widget) Label() string {
	return C.GoString(C.go_fltk_Widget_label(w.ptr))
}
func (w *widget) Color() Color {
	return Color(C.go_fltk_Widget_color(w.ptr))
}
func (w *widget) LabelFont() Font      { return Font(C.go_fltk_Widget_labelfont(w.ptr)) }
func (w *widget) LabelSize() int       { return int(C.go_fltk_Widget_labelsize(w.ptr)) }
func (w *widget) LabelType() LabelType { return LabelType(C.go_fltk_Widget_labeltype(w.ptr)) }
func (w *widget) Parent() *Group {
	ptr := C.go_fltk_Widget_parent(w.ptr)
	wid := widget{(*C.Fl_Widget)(ptr), 0, 0}
	return &Group{wid}
}
