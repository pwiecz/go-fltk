package fltk

/*
#include <stdlib.h>
#include "input_choice.h"
*/
import "C"
import "unsafe"

type InputChoice struct {
	Group
	input             *Input
	menuButton        *MenuButton
	deletionHandlerId uintptr
}

func NewInputChoice(x, y, w, h int, text ...string) *InputChoice {
	c := &InputChoice{}
	inputChoice := C.go_fltk_new_Input_Choice(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))
	initWidget(c, unsafe.Pointer(inputChoice))
	c.input = &Input{}
	initUnownedWidget(c.input, unsafe.Pointer(C.go_fltk_Input_Choice_input((*C.Fl_Input_Choice)(inputChoice))))
	c.menuButton = &MenuButton{}
	initUnownedWidget(c.menuButton, unsafe.Pointer(C.go_fltk_Input_Choice_menubutton((*C.Fl_Input_Choice)(inputChoice))))
	c.deletionHandlerId = c.addDeletionHandler(c.onDelete)
	return c
}

func (c *InputChoice) onDelete() {
	c.input.onDelete()
	c.menuButton.onDelete()
	if c.deletionHandlerId > 0 {
		globalCallbackMap.unregister(c.deletionHandlerId)
	}
	c.deletionHandlerId = 0
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
	return c.input
}
func (c *InputChoice) MenuButton() *MenuButton {
	return c.menuButton
}
