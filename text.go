package fltk

/*
#include <stdlib.h>
#include <stdint.h>
#include "text.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

// --- Modify-Callback-Map ---

type modifyCallbackMap struct {
	cbMap map[uintptr]func(int, int, int, int, string)
	id    uintptr
}

func newModifyCallbackMap() *modifyCallbackMap {
	return &modifyCallbackMap{
		cbMap: make(map[uintptr]func(int, int, int, int, string)),
	}
}
func (m *modifyCallbackMap) register(fn func(int, int, int, int, string)) uintptr {
	m.id++
	m.cbMap[m.id] = fn
	return m.id
}
func (m *modifyCallbackMap) unregister(id uintptr) {
	delete(m.cbMap, id)
}
func (m *modifyCallbackMap) invoke(id uintptr, pos, nInserted, nDeleted, nRestyled int, deletedText string) {
	if callback, ok := m.cbMap[id]; ok && callback != nil {
		callback(pos, nInserted, nDeleted, nRestyled, deletedText)
	}
}
func (m *modifyCallbackMap) isEmpty() bool {
	return len(m.cbMap) == 0
}
func (m *modifyCallbackMap) size() int {
	return len(m.cbMap)
}
func (m *modifyCallbackMap) clear() {
	for id := range m.cbMap {
		delete(m.cbMap, id)
	}
}

var globalModifyCallbackMap = newModifyCallbackMap()

//export _go_modifyCallbackHandler
func _go_modifyCallbackHandler(id C.uintptr_t, pos, nInserted, nDeleted, nRestyled C.int, deletedText *C.char) {
	globalModifyCallbackMap.invoke(uintptr(id), int(pos), int(nInserted), int(nDeleted), int(nRestyled), C.GoString(deletedText))
}

type StyleTableEntry struct {
	Color Color
	Font  Font
	Size  int
}

type TextBuffer struct {
	cPtr       *C.Fl_Text_Buffer
	handlerIds []uintptr
}

var ErrTextBufferDestroyed = errors.New("text buffer is destroyed")
var ErrNoTextBufferAssociated = errors.New("there is no text buffer associated")

func NewTextBuffer() *TextBuffer {
	ptr := C.go_fltk_new_TextBuffer()
	return &TextBuffer{cPtr: ptr}
}

func (b *TextBuffer) ptr() *C.Fl_Text_Buffer {
	if b.cPtr == nil {
		panic(ErrTextBufferDestroyed)
	}
	return b.cPtr
}
func (b *TextBuffer) Destroy() {
	for _, id := range b.handlerIds {
		globalModifyCallbackMap.unregister(id)
	}
	b.handlerIds = nil

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

func (b *TextBuffer) Insert(pos int, txt string) {
	txtstr := C.CString(txt)
	defer C.free(unsafe.Pointer(txtstr))
	C.go_fltk_TextBuffer_insert(b.ptr(), C.int(pos), txtstr)
}

func (b *TextBuffer) Remove(start, end int) {
	C.go_fltk_TextBuffer_remove(b.ptr(), C.int(start), C.int(end))
}

func (b *TextBuffer) CharAt(pos int) rune {
	return rune(C.go_fltk_TextBuffer_char_at(b.ptr(), C.int(pos)))
}

func (b *TextBuffer) NextChar(pos int) int {
	return int(C.go_fltk_TextBuffer_next_char(b.ptr(), C.int(pos)))
}

func (b *TextBuffer) PrevChar(pos int) int {
	return int(C.go_fltk_TextBuffer_prev_char(b.ptr(), C.int(pos)))
}

func (b *TextBuffer) LineStart(pos int) int {
	return int(C.go_fltk_TextBuffer_line_start(b.ptr(), C.int(pos)))
}

func (b *TextBuffer) LineEnd(pos int) int {
	return int(C.go_fltk_TextBuffer_line_end(b.ptr(), C.int(pos)))
}

func (b *TextBuffer) LineText(pos int) string {
	cStr := C.go_fltk_TextBuffer_line_text(b.ptr(), C.int(pos))
	defer C.free(unsafe.Pointer(cStr))
	return C.GoString(cStr)
}

func (b *TextBuffer) CountLines(start, end int) int {
	return int(C.go_fltk_TextBuffer_count_lines(b.ptr(), C.int(start), C.int(end)))
}

func (b *TextBuffer) SkipLines(start, nLines int) int {
	return int(C.go_fltk_TextBuffer_skip_lines(b.ptr(), C.int(start), C.int(nLines)))
}

func (b *TextBuffer) RewindLines(start, nLines int) int {
	return int(C.go_fltk_TextBuffer_rewind_lines(b.ptr(), C.int(start), C.int(nLines)))
}

func (b *TextBuffer) Length() int {
	return int(C.go_fltk_TextBuffer_length(b.ptr()))
}

func (b *TextBuffer) AddModifyCallback(cb func(int, int, int, int, string)) {
	handlerId := globalModifyCallbackMap.register(cb)
	b.handlerIds = append(b.handlerIds, handlerId)
	C.go_fltk_TextBuffer_add_modify_callback(b.ptr(), C.uintptr_t(handlerId))
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

// IsSelected checks if any text is selected
func (b *TextBuffer) IsSelected() bool {
	return C.go_fltk_TextBuffer_selected(b.ptr()) != 0
}

// GetSelectionPosition gets position (start, end) of the currently selected text
func (b *TextBuffer) GetSelectionPosition() (int, int) {
	var _start, _end C.int
	C.go_fltk_TextBuffer_selection_position(b.ptr(), &_start, &_end)
	return int(_start), int(_end)
}

// GetSelectionText returns the text within the current selection
func (b *TextBuffer) GetSelectionText() string {
	cStr := C.go_fltk_TextBuffer_selection_text(b.ptr())
	defer C.free(unsafe.Pointer(cStr))
	return C.GoString(cStr)
}

// UnSelect unselects any selections in the buffer
func (b *TextBuffer) UnSelect() {
	C.go_fltk_TextBuffer_unselect(b.ptr())
}

// SetTabWidth sets the TAB distance (width)
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

// Translate a pixel position into a character index.
func (t *TextDisplay) XYToPosition(x, y int) int {
	return int(C.go_fltk_TextDisplay_xy_to_position((*C.Fl_Text_Display)(t.ptr()), C.int(x), C.int(y)))
}

// Convert a character index into a pixel position.
func (t *TextDisplay) PositionToXY(pos int) (int, int) {
	var x, y C.int
	C.go_fltk_TextDisplay_position_to_xy((*C.Fl_Text_Display)(t.ptr()), C.int(pos), &x, &y)
	return int(x), int(y)
}

func (t *TextDisplay) Buffer() *TextBuffer {
	ptr := C.go_fltk_TextDisplay_buffer((*C.Fl_Text_Display)(t.ptr()))
	if ptr == nil {
		return nil
	}
	return &TextBuffer{cPtr: ptr}
}

// MoveRight moves the current insert position right one character.
// Returns true if the cursor moved, false if the end of the text was reached
func (t *TextDisplay) MoveRight() bool {
	if t.Buffer() == nil {
		panic(ErrNoTextBufferAssociated)
	}
	return (int)(C.go_fltk_TextDisplay_move_right((*C.Fl_Text_Display)(t.ptr()))) != 0
}

// MoveLeft moves the current insert position left one character.
func (t *TextDisplay) MoveLeft() bool {
	if t.Buffer() == nil {
		panic(ErrNoTextBufferAssociated)
	}
	return (int)(C.go_fltk_TextDisplay_move_left((*C.Fl_Text_Display)(t.ptr()))) != 0
}

// MoveUp moves the current insert position up one line.
func (t *TextDisplay) MoveUp() bool {
	if t.Buffer() == nil {
		panic(ErrNoTextBufferAssociated)
	}
	return (int)(C.go_fltk_TextDisplay_move_up((*C.Fl_Text_Display)(t.ptr()))) != 0
}

// MoveDown moves the current insert position down one line.
func (t *TextDisplay) MoveDown() bool {
	return (int)(C.go_fltk_TextDisplay_move_down((*C.Fl_Text_Display)(t.ptr()))) != 0
}

// ShowInsertPosition scrolls the text buffer to show the current insert position.
func (t *TextDisplay) ShowInsertPosition() {
	if t.Buffer() == nil {
		panic(ErrNoTextBufferAssociated)
	}
	C.go_fltk_TextDisplay_show_insert_position((*C.Fl_Text_Display)(t.ptr()))
}

// HideCursor hides the text cursor.
func (t *TextDisplay) HideCursor() {
	C.go_fltk_TextDisplay_hide_cursor((*C.Fl_Text_Display)(t.ptr()))
}

// TextColor gets the default color of text in the widget.
func (t *TextDisplay) TextColor() Color {
	return Color(C.go_fltk_TextDisplay_text_color((*C.Fl_Text_Display)(t.ptr())))
}

// SetTextColor sets the default color of text in the widget.
func (t *TextDisplay) SetTextColor(color Color) {
	C.go_fltk_TextDisplay_set_text_color((*C.Fl_Text_Display)(t.ptr()), C.uint(color))
}

// TextFont gets the default font used when drawing text in the widget.
func (t *TextDisplay) TextFont() Font {
	return Font(C.go_fltk_TextDisplay_text_font((*C.Fl_Text_Display)(t.ptr())))
}

// SetTextFont sets the default font used when drawing text in the widget.
func (t *TextDisplay) SetTextFont(font Font) {
	C.go_fltk_TextDisplay_set_text_font((*C.Fl_Text_Display)(t.ptr()), C.int(font))
}

// TextSize gets the default size of text in the widget.
func (t *TextDisplay) TextSize() int {
	return (int)(C.go_fltk_TextDisplay_text_size((*C.Fl_Text_Display)(t.ptr())))
}

// SetTextSize sets the default size of text in the widget.
func (t *TextDisplay) SetTextSize(size int) {
	C.go_fltk_TextDisplay_set_text_size((*C.Fl_Text_Display)(t.ptr()), C.int(size))
}

// SetInsertPosition set the insert position to a new position.
func (t *TextDisplay) SetInsertPosition(newPos int) {
	if t.Buffer() == nil {
		panic(ErrNoTextBufferAssociated)
	}
	C.go_fltk_TextDisplay_insert_position((*C.Fl_Text_Display)(t.ptr()), C.int(newPos))
}

// GetInsertPosition - return the current insert position.
func (t *TextDisplay) GetInsertPosition() int {
	if t.Buffer() == nil {
		panic(ErrNoTextBufferAssociated)
	}
	return (int)(C.go_fltk_TextDisplay_get_insert_position((*C.Fl_Text_Display)(t.ptr())))
}

// InsertText - Insert text at the cursor position.
func (t *TextDisplay) InsertText(txt string) {
	if t.Buffer() == nil {
		panic(ErrNoTextBufferAssociated)
	}
	txtstr := C.CString(txt)
	defer C.free(unsafe.Pointer(txtstr))
	C.go_fltk_TextDisplay_insert_text((*C.Fl_Text_Display)(t.ptr()), txtstr)
}

// Overstrike - Not sure what it does, the fltk doc does not match with the name meaning
func (t *TextDisplay) Overstrike(txt string) {
	if t.Buffer() == nil {
		panic(ErrNoTextBufferAssociated)
	}
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

// SetLinenumberWidth enabled/disables and sets the width used by line numbers.
//
// A width of 0 pixels disables line numbers. A width > 0 enables line
// numbers and sets that as the width to be used.
func (t *TextDisplay) SetLinenumberWidth(w int) {
	C.go_fltk_TextDisplay_set_linenumber_width((*C.Fl_Text_Display)(t.ptr()), C.int(w))
}

// SetLinenumberSize sets the font size used for line numbers (if enabled; see SetLinenumberWidth).
func (t *TextDisplay) SetLinenumberSize(s int) {
	C.go_fltk_TextDisplay_set_linenumber_size((*C.Fl_Text_Display)(t.ptr()), C.int(s))
}

// SetLinenumberFgcolor sets the foreground color used for line numbers (if enabled; see SetLinenumberWidth).
func (t *TextDisplay) SetLinenumberFgcolor(color Color) {
	C.go_fltk_TextDisplay_set_linenumber_fgcolor((*C.Fl_Text_Display)(t.ptr()), C.uint(color))
}

// SetLinenumberBgcolor sets the background color used for line numbers (if enabled; see SetLinenumberWidth).
func (t *TextDisplay) SetLinenumberBgcolor(color Color) {
	C.go_fltk_TextDisplay_set_linenumber_bgcolor((*C.Fl_Text_Display)(t.ptr()), C.uint(color))
}

// SetLinenumberFont sets the font used for line numbers (if enabled; see SetLinenumberWidth).
func (t *TextDisplay) SetLinenumberFont(font Font) {
	C.go_fltk_TextDisplay_set_linenumber_font((*C.Fl_Text_Display)(t.ptr()), C.int(font))
}

// SetLinenumberAlign sets the alignment used for line numbers (if enabled; see SetLinenumberWidth).
func (t *TextDisplay) SetLinenumberAlign(align Align) {
	C.go_fltk_TextDisplay_set_linenumber_align((*C.Fl_Text_Display)(t.ptr()), C.int(align))
}

type TextEditor struct {
	TextDisplay
}

// Copy copy of selected text or the current character in the current buffer of the editor (kf_copy)
func (t *TextEditor) Copy() {
	if t.Buffer() == nil {
		panic(ErrNoTextBufferAssociated)
	}
	C.go_fltk_TextEditor_copy((*C.Fl_Text_Editor)(t.ptr()))
}

// Insert toggles the insert mode (kf_insert)
func (t *TextEditor) Insert() {
	if t.Buffer() == nil {
		panic(ErrNoTextBufferAssociated)
	}
	C.go_fltk_TextEditor_insert((*C.Fl_Text_Editor)(t.ptr()))
}

// Cut cuts the selected text from the editor's buffer into the clipboard.
func (t *TextEditor) Cut() {
	if t.Buffer() == nil {
		panic(ErrNoTextBufferAssociated)
	}
	C.go_fltk_TextEditor_cut((*C.Fl_Text_Editor)(t.ptr()))
}

// Delete deletes the selected text from the editor's buffer.
func (t *TextEditor) Delete() {
	if t.Buffer() == nil {
		panic(ErrNoTextBufferAssociated)
	}
	C.go_fltk_TextEditor_delete((*C.Fl_Text_Editor)(t.ptr()))
}

// Paste pastes the clipboard's text into the editor's buffer at the
// insertion point.
func (t *TextEditor) Paste() {
	if t.Buffer() == nil {
		panic(ErrNoTextBufferAssociated)
	}
	C.go_fltk_TextEditor_paste((*C.Fl_Text_Editor)(t.ptr()))
}

// SelectAll selects all the editor's text.
func (t *TextEditor) SelectAll() {
	if t.Buffer() == nil {
		panic(ErrNoTextBufferAssociated)
	}
	C.go_fltk_TextEditor_select_all((*C.Fl_Text_Editor)(t.ptr()))
}

// Undo undoes the last edit.
// Returns true iff any change was made.
func (t *TextEditor) Undo() bool {
	if t.Buffer() == nil {
		panic(ErrNoTextBufferAssociated)
	}
	return C.go_fltk_TextEditor_undo((*C.Fl_Text_Editor)(t.ptr())) != 0
}

// Redo redoes the last undone edit.
// Returns true iff any change was made.
func (t *TextEditor) Redo() bool {
	if t.Buffer() == nil {
		panic(ErrNoTextBufferAssociated)
	}
	return C.go_fltk_TextEditor_redo((*C.Fl_Text_Editor)(t.ptr())) != 0
}

// NewTextEditor returns a TextEditor.
//
// Example:
//
//	textBuffer := fltk.NewTextBuffer()
//	textEditor := fltk.NewTextEditor(x, y, width, height)
//	textEditor.SetBuffer(textBuffer)
//	textBuffer.SetText("Initial Text")
//	fmt.Println(textBuffer.Text()) // Prints: Initial Text
//
// Note that the text buffer pointer must be kept around for as long as the
// text editor is in use.
func NewTextEditor(x, y, w, h int, text ...string) *TextEditor {
	t := &TextEditor{}
	initWidget(t, unsafe.Pointer(C.go_fltk_new_TextEditor(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return t
}
