package fltk

/*
#include <stdlib.h>
#include "input_choice.h"
*/
import "C"
import "unsafe"

type InputChoice struct {
	Group
}

func NewInputChoice(x, y, w, h int, text ...string) *InputChoice {
	c := &InputChoice{}
	initWidget(c, unsafe.Pointer(C.go_fltk_new_Input_Choice(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return c
}


func (c *InputChoice) Clear() {
	C.go_fltk_Input_Choice_clear((*C.Fl_Input_Choice)(c.ptr()))
}

func (c *InputChoice) Value() string {
	return C.GoString(C.go_fltk_Input_Choice_value((*C.Fl_Input_Choice)(c.ptr())))
}

func (c *InputChoice) SetValue(label string) {
	labelStr := C.CString(label)
	defer C.free(unsafe.Pointer(labelStr))
	C.go_fltk_Input_Choice_set_value((*C.Fl_Input_Choice)(c.ptr()), labelStr)
}

func (c *InputChoice) SetValueIndex(index int) {
	C.go_fltk_Input_Choice_set_value_index((*C.Fl_Input_Choice)(c.ptr()), C.int(index))
}

func (c *InputChoice) UpdateMenuButton() bool {
	return C.go_fltk_Input_Choice_update_menubutton((*C.Fl_Input_Choice)(c.ptr())) == 1
}

func (c *InputChoice) Input() *Input {
	input := &Input{}
	cInput := C.go_fltk_Input_Choice_input((*C.Fl_Input_Choice)(c.ptr()))
	initUnownedWidget(input, unsafe.Pointer(cInput))
	return input
}
func (c *InputChoice) MenuButton() *MenuButton {
	menuButton := &MenuButton{}
	cMenuButton := C.go_fltk_Input_Choice_menubutton((*C.Fl_Input_Choice)(c.ptr()))
	initUnownedWidget(menuButton, unsafe.Pointer(cMenuButton))
	return menuButton
}
