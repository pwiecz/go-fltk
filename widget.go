package fltk

/*
#include "widget.h"
*/
import "C"
import "unsafe"

type widget struct {
	ptr        *C.Fl_Widget
	callbackId uintptr
	callback   func()
}

type Widget interface {
	GetWidget() *widget
}

func initWidget(w Widget, p unsafe.Pointer) {
	w.GetWidget().ptr = (*C.Fl_Widget)(p)
}

func (w *widget) SetCallback(f func()) {
	if w.callbackId > 0 {
		globalCallbackMap.unregister(w.callbackId)
	}
	w.callbackId = globalCallbackMap.register(f)
	C.go_fltk_Widget_set_callback(w.ptr, unsafe.Pointer(w.callbackId))
}

func (w *widget) GetWidget() *widget           { return w }
func (w *widget) SetBox(box BoxType)           { C.go_fltk_Widget_set_box(w.ptr, C.int(box)) }
func (w *widget) SetLabelFont(font Font)       { C.go_fltk_Widget_set_labelfont(w.ptr, C.int(font)) }
func (w *widget) SetLabelSize(size int)        { C.go_fltk_Widget_set_labelsize(w.ptr, C.int(size)) }
func (w *widget) SetLabelType(ltype LabelType) { C.go_fltk_Widget_set_labeltype(w.ptr, C.int(ltype)) }
func (w *widget) HandleCallback() {
	if w.callback != nil {
		w.callback()
	}
}
func (w *widget) X() int               { return int(C.go_fltk_Widget_x(w.ptr)) }
func (w *widget) Y() int               { return int(C.go_fltk_Widget_y(w.ptr)) }
func (w *widget) W() int               { return int(C.go_fltk_Widget_w(w.ptr)) }
func (w *widget) H() int               { return int(C.go_fltk_Widget_h(w.ptr)) }
func (w *widget) SetAlign(align Align) { C.go_fltk_Widget_set_align(w.ptr, C.uint(align)) }
func (w *widget) MeasureLabel() (int, int) {
	var ww, hh C.int
	ww = 0
	C.go_fltk_Widget_measure_label(w.ptr, &ww, &hh)
	return int(ww), int(hh)
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
func (w *widget) SetSelectionColor(color uint) {
	C.go_fltk_Widget_set_selection_color(w.ptr, C.uint(color))
}
