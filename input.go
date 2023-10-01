package fltk

/*
#include <stdlib.h>
#include "input.h"
*/
import "C"
import "unsafe"

type Input struct {
	widget
}

func NewInput(x, y, w, h int, text ...string) *Input {
	i := &Input{}
	initWidget(i, unsafe.Pointer(C.go_fltk_new_Input(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return i
}

func (i *Input) Value() string {
	return C.GoString(C.go_fltk_Input_value((*C.Fl_Input)(i.ptr())))
}
func (i *Input) SetValue(value string) bool {
	valueStr := C.CString(value)
	defer C.free(unsafe.Pointer(valueStr))
	return C.go_fltk_Input_set_value((*C.Fl_Input)(i.ptr()), valueStr) != 0
}
func (i *Input) Resize(x int, y int, w int, h int) {
	C.go_fltk_Input_resize((*C.Fl_Input)(i.ptr()), C.int(x), C.int(y), C.int(w), C.int(h))
}
func (i *Input) SetInsertPosition(pos, mark int) {
	C.go_fltk_Input_set_insert_position((*C.Fl_Input)(i.ptr()), C.int(pos), C.int(mark))
}
func (i *Input) InsertPosition() int {
	return int(C.go_fltk_Input_insert_position((*C.Fl_Input)(i.ptr())))
}
func (i *Input) Mark() int {
	return int(C.go_fltk_Input_mark((*C.Fl_Input)(i.ptr())))
}

type Output struct {
	Input
}

func NewOutput(x, y, w, h int, text ...string) *Output {
	i := &Output{}
	initWidget(i, unsafe.Pointer(C.go_fltk_new_Output(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return i
}

type FloatInput struct {
	Input
}

func NewFloatInput(x, y, w, h int, text ...string) *FloatInput {
	i := &FloatInput{}
	initWidget(i, unsafe.Pointer(C.go_fltk_new_Float_Input(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return i
}

func NewIntInput(x, y, w, h int, text ...string) *FloatInput {
	i := &FloatInput{}
	initWidget(i, unsafe.Pointer(C.go_fltk_new_Int_Input(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return i
}
