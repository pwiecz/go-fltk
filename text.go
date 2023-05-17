package fltk

/*
#include <stdlib.h>
#include "text.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type StyleTableEntry struct {
	Color Color
	Font  Font
	Size  int
}

type TextBuffer struct {
	cPtr *C.Fl_Text_Buffer
}

var ErrTextBufferDestroyed = errors.New("text buffer is destroyed")

func NewTextBuffer() *TextBuffer {
	ptr := C.go_fltk_new_TextBuffer()
	return &TextBuffer{ptr}
}

func (b *TextBuffer) ptr() *C.Fl_Text_Buffer {
	if b.cPtr == nil {
		panic(ErrTextBufferDestroyed)
	}
	return b.cPtr
}
func (b *TextBuffer) Destroy() {
	C.go_fltk_TextBuffer_delete(b.ptr())
	b.cPtr = nil
}

func (b *TextBuffer) SetText(txt string) {
	txtstr := C.CString(txt)
	defer C.free(unsafe.Pointer(txtstr))
	C.go_fltk_TextBuffer_set_text(b.ptr(), txtstr)
}

func (b *TextBuffer) Append(txt string) {
	txtstr := C.CString(txt)
	defer C.free(unsafe.Pointer(txtstr))
	C.go_fltk_TextBuffer_append(b.ptr(), txtstr)
}

func (b *TextBuffer) Text() string {
	cStr := C.go_fltk_TextBuffer_text(b.ptr())
	defer C.free(unsafe.Pointer(cStr))
	return C.GoString(cStr)
}

// GetTextRange - get text from start and end position
func (b *TextBuffer) GetTextRange(start, end int) string {
	cStr := C.go_fltk_TextBuffer_text_range(b.ptr(), C.int(start), C.int(end))
	defer C.free(unsafe.Pointer(cStr))
	return C.GoString(cStr)
}

// Highlight - highlight text from start and end position
func (b *TextBuffer) Highlight(start, end int) {
	C.go_fltk_TextBuffer_highlight(b.ptr(), C.int(start), C.int(end))
}

// UnHighlight all highlighted text
func (b *TextBuffer) UnHighlight() {
	C.go_fltk_TextBuffer_unhighlight(b.ptr())
}

// ReplaceRange - replace text within the start, end range with text
func (b *TextBuffer) ReplaceRange(start, end int, txt string) {
	txtstr := C.CString(txt)
	defer C.free(unsafe.Pointer(txtstr))
	C.go_fltk_TextBuffer_replace(b.ptr(), C.int(start), C.int(end), txtstr)
}

// ReplaceSelection - replace text within current selection with text
func (b *TextBuffer) ReplaceSelection(txt string) {
	txtstr := C.CString(txt)
	defer C.free(unsafe.Pointer(txtstr))
	C.go_fltk_TextBuffer_replace_selection(b.ptr(), txtstr)
}

// SearchForward - search forward for searchStr from position start. return the position if found otherwise -1
func (b *TextBuffer) SearchForward(start int, searchStr string, matchCase bool) (foundPos int) {
	txtstr := C.CString(searchStr)
	defer C.free(unsafe.Pointer(txtstr))
	_matchCase := 1
	if !matchCase {
		_matchCase = 0
	}
	var C_foundPos C.int
	_isFound := C.go_fltk_TextBuffer_search_forward(b.ptr(), C.int(start), txtstr, &C_foundPos, C.int(_matchCase))
	if int(_isFound) == 0 {
		foundPos = -1
	} else {
		foundPos = int(C_foundPos)
	}
	return foundPos
}

// SearchBackward - search backward for searchStr from position start. return the position if found otherwise -1
func (b *TextBuffer) SearchBackward(start int, searchStr string, matchCase bool) (foundPos int) {
	txtstr := C.CString(searchStr)
	defer C.free(unsafe.Pointer(txtstr))
	_matchCase := 1
	if !matchCase {
		_matchCase = 0
	}
	var C_foundPos C.int
	_isFound := C.go_fltk_TextBuffer_search_backward(b.ptr(), C.int(start), txtstr, &C_foundPos, C.int(_matchCase))
	if int(_isFound) == 0 {
		foundPos = -1
	} else {
		foundPos = int(C_foundPos)
	}
	return foundPos
}

// Search - search text in the buffer for string searchStr. Return the position found or -1
func (b *TextBuffer) Search(start int, searchStr string, searchBackward, matchCase bool) (foundPos int) {
	if searchBackward {
		return b.SearchBackward(start, searchStr, matchCase)
	} else {
		return b.SearchForward(start, searchStr, matchCase)
	}
}

// Select text between start and end
func (b *TextBuffer) Select(start, end int) {
	C.go_fltk_TextBuffer_select(b.ptr(), C.int(start), C.int(end))
}

// Selected Check if any text is selected
func (b *TextBuffer) IsSelected(start, end int) bool {
	return C.go_fltk_TextBuffer_selected(b.ptr()) != 0
}

// GetSelectionPosition - Get position (start, end) of the currently selected text
func (b *TextBuffer) GetSelectionPosition() (int, int) {
	var _start, _end C.int
	C.go_fltk_TextBuffer_selection_position(b.ptr(), &_start, &_end)
	return int(_start), int(_end)
}

// GetSelectionText return the text within the current selection
func (b *TextBuffer) GetSelectionText() string {
	cStr := C.go_fltk_TextBuffer_selection_text(b.ptr())
	defer C.free(unsafe.Pointer(cStr))
	return C.GoString(cStr)
}

// UnSelect - unselect any selections in the buffer
func (b *TextBuffer) UnSelect() {
	C.go_fltk_TextBuffer_unselect(b.ptr())
}

// SetTabWidth - set the TAB distance (width)
func (b *TextBuffer) SetTabWidth(tabWidth int) {
	C.go_fltk_TextBuffer_set_tab_distance(b.ptr(), C.int(tabWidth))
}

func (b *TextBuffer) GetTabWidth() int {
	w := C.go_fltk_TextBuffer_tab_distance(b.ptr())
	return int(w)
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
	C.go_fltk_TextDisplay_set_buffer((*C.Fl_Text_Display)(t.ptr()), buf.ptr())
}

// wrapMargin is not needed if WrapMode is WRAP_NONE or WRAP_AT_BOUNDS
func (t *TextDisplay) SetWrapMode(wrap WrapMode, wrapMargin ...int) {
	if len(wrapMargin) < 1 {
		wrapMargin = append(wrapMargin, 0)
	}

	C.go_fltk_TextDisplay_set_wrap_mode((*C.Fl_Text_Display)(t.ptr()), C.int(wrap), C.int(wrapMargin[0]))
}

func (t *TextDisplay) Buffer() *TextBuffer {
	ptr := C.go_fltk_TextDisplay_buffer((*C.Fl_Text_Display)(t.ptr()))
	return &TextBuffer{ptr}
}

// MoveRight moves the current insert position right one character.
// Returns true if the cursor moved, false if the end of the text was reached
func (t *TextDisplay) MoveRight() bool {
	return (int)(C.go_fltk_TextDisplay_move_right((*C.Fl_Text_Display)(t.ptr()))) != 0
}

// MoveLeft moves the current insert position left one character.
func (t *TextDisplay) MoveLeft() bool {
	return (int)(C.go_fltk_TextDisplay_move_left((*C.Fl_Text_Display)(t.ptr()))) != 0
}

// MoveUp moves the current insert position up one line.
func (t *TextDisplay) MoveUp() bool {
	return (int)(C.go_fltk_TextDisplay_move_up((*C.Fl_Text_Display)(t.ptr()))) != 0
}

// MoveDown moves the current insert position down one line.
func (t *TextDisplay) MoveDown() bool {
	return (int)(C.go_fltk_TextDisplay_move_down((*C.Fl_Text_Display)(t.ptr()))) != 0
}

// ShowInsertPosition scrolls the text buffer to show the current insert position.
func (t *TextDisplay) ShowInsertPosition() {
	C.go_fltk_TextDisplay_show_insert_position((*C.Fl_Text_Display)(t.ptr()))
}

// TextColor gets the default color of text in the widget
func (t *TextDisplay) TextColor() Color {
	return Color(C.go_fltk_TextDisplay_text_color((*C.Fl_Text_Display)(t.ptr())))
}

// SetTextColor sets the default color of text in the widget
func (t *TextDisplay) SetTextColor(color Color) {
	C.go_fltk_TextDisplay_set_text_color((*C.Fl_Text_Display)(t.ptr()), C.uint(color))
}

// TextSize gets the default size of text in the widget
func (t *TextDisplay) TextSize() int {
	return (int)(C.go_fltk_TextDisplay_text_size((*C.Fl_Text_Display)(t.ptr())))
}

// SetTextSize sets the default size of text in the widget
func (t *TextDisplay) SetTextSize(size int) {
	C.go_fltk_TextDisplay_set_text_size((*C.Fl_Text_Display)(t.ptr()), C.int(size))
}

// SetInsertPosition set the insert position to a new position
func (t *TextDisplay) SetInsertPosition(newPos int) {
	C.go_fltk_TextDisplay_insert_position((*C.Fl_Text_Display)(t.ptr()), C.int(newPos))
}

// GetInsertPosition - return the current insert position
func (t *TextDisplay) GetInsertPosition() int {
	return (int)(C.go_fltk_TextDisplay_get_insert_position((*C.Fl_Text_Display)(t.ptr())))
}

// InsertText - Insert text at the cursor position
func (t *TextDisplay) InsertText(txt string) {
	txtstr := C.CString(txt)
	defer C.free(unsafe.Pointer(txtstr))
	C.go_fltk_TextDisplay_insert_text((*C.Fl_Text_Display)(t.ptr()), txtstr)
}

// Overstrike - Not sure what it does, the fltk doc does not match with the name meaning
func (t *TextDisplay) Overstrike(txt string) {
	txtstr := C.CString(txt)
	defer C.free(unsafe.Pointer(txtstr))
	C.go_fltk_TextDisplay_overstrike((*C.Fl_Text_Display)(t.ptr()), txtstr)
}

func (t *TextDisplay) SetHighlightData(buf *TextBuffer, entries []StyleTableEntry) {
	var colors []C.uint
	var fonts []C.int
	var sizes []C.int
	var attrs []C.uint
	var bgcolors []C.uint
	for i := 0; i < len(entries); i++ {
		colors = append(colors, C.uint(entries[i].Color))
		fonts = append(fonts, C.int(entries[i].Font))
		sizes = append(sizes, C.int(entries[i].Size))
		attrs = append(attrs, 0)
		bgcolors = append(bgcolors, 0)
	}
	C.go_fltk_TextDisplay_set_highlight_data((*C.Fl_Text_Display)(t.ptr()), buf.ptr(), &colors[0], &fonts[0], &sizes[0], &attrs[0], &bgcolors[0], C.int(len(entries)))
}

type TextEditor struct {
	TextDisplay
}

// Copy copy of selected text or the current character in the current buffer of editor 'e'. (kf_copy)
func (t *TextEditor) Copy() {
	C.go_fltk_TextEditor_copy((*C.Fl_Text_Editor)(t.ptr()))
}

// Insert - Togglesthe insert mode (kf_insert)
func (t *TextEditor) Insert() {
	C.go_fltk_TextEditor_insert((*C.Fl_Text_Editor)(t.ptr()))
}

func (t *TextEditor) Cut() {
	C.go_fltk_TextEditor_cut((*C.Fl_Text_Editor)(t.ptr()))
}

func (t *TextEditor) Delete() {
	C.go_fltk_TextEditor_delete((*C.Fl_Text_Editor)(t.ptr()))
}

func (t *TextEditor) Paste() {
	C.go_fltk_TextEditor_paste((*C.Fl_Text_Editor)(t.ptr()))
}

func (t *TextEditor) SelectAll() {
	C.go_fltk_TextEditor_select_all((*C.Fl_Text_Editor)(t.ptr()))
}

// NewTextEditor returns a TextEditor.
//
// Example:
//		textBuffer := fltk.NewTextBuffer()
//		textEditor := fltk.NewTextEditor(x, y, width, height)
//		textEditor.SetBuffer(textBuffer)
//		textBuffer.SetText("Initial Text")
//		fmt.Println(textBuffer.Text()) // Prints: Initial Text
//
// Note that the text buffer pointer must be kept around for as long as the
// text editor is in use.
func NewTextEditor(x, y, w, h int, text ...string) *TextEditor {
	t := &TextEditor{}
	initWidget(t, unsafe.Pointer(C.go_fltk_new_TextEditor(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return t
}
