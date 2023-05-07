package fltk

/*
#include "button.h"
*/
import "C"
import "unsafe"

type Button struct {
	widget
}

func NewButton(x, y, w, h int, text ...string) *Button {
	b := &Button{}
	initWidget(b, unsafe.Pointer(C.go_fltk_new_Button(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return b
}

func (b *Button) Value() bool {
	return C.go_fltk_Button_value((*C.Fl_Button)(b.ptr())) != C.char(0)
}
func (b *Button) SetValue(val bool) {
	if val {
		C.go_fltk_Button_set_value((*C.Fl_Button)(b.ptr()), 1)
	} else {
		C.go_fltk_Button_set_value((*C.Fl_Button)(b.ptr()), 0)
	}
}
func (b *Button) SetDownBox(box BoxType) {
	C.go_fltk_Button_set_down_box((*C.Fl_Button)(b.ptr()), C.int(box))
}
func (b *Button) SetShortcut(shortcut int) {
	C.go_fltk_Button_set_shortcut((*C.Fl_Button)(b.ptr()), C.int(shortcut))
}
func (b *Button) Shortcut() int {
	return int(C.go_fltk_Button_shortcut((*C.Fl_Button)(b.ptr())))
}

type CheckButton struct {
	Button
}

func NewCheckButton(x, y, w, h int, text ...string) *CheckButton {
	i := &CheckButton{}
	initWidget(i, unsafe.Pointer(C.go_fltk_new_Check_Button(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return i
}

type RadioButton struct {
	Button
}

func NewRadioButton(x, y, w, h int, text ...string) *RadioButton {
	i := &RadioButton{}
	initWidget(i, unsafe.Pointer(C.go_fltk_new_Radio_Button(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return i
}

type RadioRoundButton struct {
	Button
}

func NewRadioRoundButton(x, y, w, h int, text ...string) *RadioRoundButton {
	i := &RadioRoundButton{}
	initWidget(i, unsafe.Pointer(C.go_fltk_new_Radio_Round_Button(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return i
}

type ReturnButton struct {
	Button
}

func NewReturnButton(x, y, w, h int, text ...string) *ReturnButton {
	i := &ReturnButton{}
	initWidget(i, unsafe.Pointer(C.go_fltk_new_Return_Button(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return i
}

type ToggleButton struct {
	Button
}

func NewToggleButton(x, y, w, h int, text ...string) *ToggleButton {
	i := &ToggleButton{}
	initWidget(i, unsafe.Pointer(C.go_fltk_new_Toggle_Button(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return i
}
