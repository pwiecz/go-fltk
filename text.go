package fltk

/*
#include <stdlib.h>
#include "text.h"
*/
import "C"
import "unsafe"

type TextBuffer struct {
	ptr *C.Fl_Text_Buffer
}

func NewTextBuffer() *TextBuffer {
	ptr := C.go_fltk_new_TextBuffer()
	return &TextBuffer{ptr}
}

func (b *TextBuffer) SetText(txt string) {
	txtstr := C.CString(txt)
	defer C.free(unsafe.Pointer(txtstr))
	C.go_fltk_TextBuffer_set_text(b.ptr, txtstr)
}

func (b *TextBuffer) Append(txt string) {
	txtstr := C.CString(txt)
	defer C.free(unsafe.Pointer(txtstr))
	C.go_fltk_TextBuffer_append(b.ptr, txtstr)
}

func (b *TextBuffer) Text() string {
	return C.GoString(C.go_fltk_TextBuffer_text(b.ptr))
}

type TextDisplay struct {
	widget
}

func NewTextDisplay(x, y, w, h int, text ...string) *TextDisplay {
	t := &TextDisplay{}
	initWidget(t, unsafe.Pointer(C.go_fltk_new_TextDisplay(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return t
}

func (t *TextDisplay) SetBuffer(buf *TextBuffer) {
	C.go_fltk_TextDisplay_set_buffer((*C.Fl_Text_Display)(t.ptr), buf.ptr)
}

// wrapMargin is not needed if WrapMode is WRAP_NONE or WRAP_AT_BOUNDS
func (t *TextDisplay) SetWrapMode(wrap WrapMode, wrapMargin ...int) {
    if len(wrapMargin) < 1 {
        wrapMargin = append(wrapMargin, 0)
    }

	C.go_fltk_TextDisplay_set_wrap_mode((*C.Fl_Text_Display)(t.ptr), C.int(wrap), C.int(wrapMargin[0]))
}

func (t *TextDisplay) Buffer() *TextBuffer {
	ptr := C.go_fltk_TextDisplay_buffer((*C.Fl_Text_Display)(t.ptr))
	return &TextBuffer{ptr}
}

type TextEditor struct {
	TextDisplay
}

func NewTextEditor(x, y, w, h int, text ...string) *TextEditor {
	t := &TextEditor{}
	initWidget(t, unsafe.Pointer(C.go_fltk_new_TextEditor(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return t
}
