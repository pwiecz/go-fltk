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
	b := &CheckButton{}
	initWidget(b, unsafe.Pointer(C.go_fltk_new_Check_Button(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return b
}

type RadioButton struct {
	Button
}

func NewRadioButton(x, y, w, h int, text ...string) *RadioButton {
	b := &RadioButton{}
	initWidget(b, unsafe.Pointer(C.go_fltk_new_Radio_Button(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return b
}

type RadioRoundButton struct {
	Button
}

func NewRadioRoundButton(x, y, w, h int, text ...string) *RadioRoundButton {
	b := &RadioRoundButton{}
	initWidget(b, unsafe.Pointer(C.go_fltk_new_Radio_Round_Button(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return b
}

type ReturnButton struct {
	Button
}

func NewReturnButton(x, y, w, h int, text ...string) *ReturnButton {
	b := &ReturnButton{}
	initWidget(b, unsafe.Pointer(C.go_fltk_new_Return_Button(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return b
}

type ToggleButton struct {
	Button
}

func NewToggleButton(x, y, w, h int, text ...string) *ToggleButton {
	b := &ToggleButton{}
	initWidget(b, unsafe.Pointer(C.go_fltk_new_Toggle_Button(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return b
}

type LightButton struct {
	Button
}

func NewLightButton(x, y, w, h int, text ...string) *LightButton {
	b := &LightButton{}
	initWidget(b, unsafe.Pointer(C.go_fltk_new_Light_Button(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return b
}
