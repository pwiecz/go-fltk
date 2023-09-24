package fltk

/*
#include "wizard.h"
*/
import "C"
import "unsafe"

type Wizard struct {
	Group
}

func NewWizard(x, y, w, h int, text ...string) *Wizard {
	wiz := &Wizard{}
	initWidget(wiz, unsafe.Pointer(C.go_fltk_new_Wizard(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return wiz
}

func (w *Wizard) Next() {
	C.go_fltk_Wizard_next((*C.Fl_Wizard)(w.ptr()))
}

func (w *Wizard) Prev() {
	C.go_fltk_Wizard_prev((*C.Fl_Wizard)(w.ptr()))
}

func (w *Wizard) SetValue(widget Widget) {
	C.go_fltk_Wizard_set_value((*C.Fl_Wizard)(w.ptr()), widget.getWidget().ptr())
}

func (w *Wizard) Value() *widget {
	value := C.go_fltk_Wizard_value((*C.Fl_Wizard)(w.ptr()))
	if value == nil {
		return nil
	}
	widget := &widget{}
	initUnownedWidget(widget, unsafe.Pointer(value))
	return widget
}
