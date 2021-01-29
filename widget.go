package fltk

/*
#include "widget.h"
*/
import "C"
import "unsafe"

var widgets = make(map[*C.Fl_Widget]Widgety)

type Widgety interface {
	GetWidget() *Widget
}

type Widget struct {
	ptr        *C.Fl_Widget
	callbackId uintptr
	callback   func()
}

func initWidget(w Widgety, p unsafe.Pointer) *Widget {
	w.GetWidget().ptr = (*C.Fl_Widget)(p)
	w.GetWidget().callback = emptyCallback
	widgets[w.GetWidget().ptr] = w
	return w.GetWidget()
}

//func NewWidget(x, y, w, h int, text... string) *Widget {
//	return initWidget(&Widget{}, unsafe.Pointer(C.go_fltk_new_Widget(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
//}

func (w *Widget) SetCallback(f func()) {
	if w.callbackId > 0 {
		globalCallbackMap.unregister(w.callbackId)
	}
	w.callbackId = globalCallbackMap.register(f)
	C.go_fltk_Widget_set_callback(w.ptr, unsafe.Pointer(w.callbackId))
}

func (w *Widget) GetWidget() *Widget           { return w }
func (w *Widget) SetBox(box BoxType)           { C.go_fltk_Widget_set_box(w.ptr, C.int(box)) }
func (w *Widget) SetLabelFont(font Font)       { C.go_fltk_Widget_set_labelfont(w.ptr, C.int(font)) }
func (w *Widget) SetLabelSize(size int)        { C.go_fltk_Widget_set_labelsize(w.ptr, C.int(size)) }
func (w *Widget) SetLabelType(ltype LabelType) { C.go_fltk_Widget_set_labeltype(w.ptr, C.int(ltype)) }
func (w *Widget) HandleCallback()              { w.callback() }
func (w *Widget) X() int                       { return int(C.go_fltk_Widget_x(w.ptr)) }
func (w *Widget) Y() int                       { return int(C.go_fltk_Widget_y(w.ptr)) }
func (w *Widget) W() int                       { return int(C.go_fltk_Widget_w(w.ptr)) }
func (w *Widget) H() int                       { return int(C.go_fltk_Widget_h(w.ptr)) }
func (w *Widget) SetAlign(align Align)         { C.go_fltk_Widget_set_align(w.ptr, C.uint(align)) }
func (w *Widget) MeasureLabel() (int, int) {
	var ww, hh C.int
	ww = 0
	C.go_fltk_Widget_measure_label(w.ptr, &ww, &hh)
	return int(ww), int(hh)
}
func (w *Widget) SetPosition(x, y int) { C.go_fltk_Widget_set_position(w.ptr, C.int(x), C.int(y)) }
func (w *Widget) Resize(x, y, width, height int) {
	C.go_fltk_Widget_resize(w.ptr, C.int(x), C.int(y), C.int(width), C.int(height))
}
func (w *Widget) Redraw()                  { C.go_fltk_Widget_redraw(w.ptr) }
func (w *Widget) Deactivate()              { C.go_fltk_Widget_deactivate(w.ptr) }
func (w *Widget) Activate()                { C.go_fltk_Widget_activate(w.ptr) }
func (w *Widget) SetType(widgetType uint8) { C.go_fltk_Widget_set_type(w.ptr, C.uchar(widgetType)) }
func (w *Widget) Show()                    { C.go_fltk_Widget_show(w.ptr) }
func (w *Widget) Hide()                    { C.go_fltk_Widget_hide(w.ptr) }
func (w *Widget) SetSelectionColor(color uint) {
	C.go_fltk_Widget_set_selection_color(w.ptr, C.uint(color))
}
