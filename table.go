package fltk

/*
#include "table.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type table struct {
	Group
}

func (t *table) SetRowCount(rowCount int) {
	C.go_fltk_Table_set_row_count((*C.Fl_Table)(t.ptr), C.int(rowCount))
}
func (t *table) SetRowHeight(row, height int) {
	C.go_fltk_Table_set_row_height((*C.Fl_Table)(t.ptr), C.int(row), C.int(height))
}
func (t *table) SetRowHeightAll(height int) {
	C.go_fltk_Table_set_row_height_all((*C.Fl_Table)(t.ptr), C.int(height))
}
func (t *table) EnableRowHeaders() {
	C.go_fltk_Table_set_row_header((*C.Fl_Table)(t.ptr), 1)
}
func (t *table) DisableRowHeaders() {
	C.go_fltk_Table_set_row_header((*C.Fl_Table)(t.ptr), 0)
}
func (t *table) AllowRowResizing() {
	C.go_fltk_Table_set_row_resize((*C.Fl_Table)(t.ptr), 1)
}
func (t *table) DisallowRowResizing() {
	C.go_fltk_Table_set_row_resize((*C.Fl_Table)(t.ptr), 0)
}
func (t *table) SetColumnCount(columnCount int) {
	C.go_fltk_Table_set_column_count((*C.Fl_Table)(t.ptr), C.int(columnCount))
}
func (t *table) SetColumnWidth(column, width int) {
	C.go_fltk_Table_set_column_width((*C.Fl_Table)(t.ptr), C.int(column), C.int(width))
}
func (t *table) SetColumnWidthAll(width int) {
	C.go_fltk_Table_set_column_width_all((*C.Fl_Table)(t.ptr), C.int(width))
}
func (t *table) EnableColumnHeaders() {
	C.go_fltk_Table_set_column_header((*C.Fl_Table)(t.ptr), 1)
}
func (t *table) DisableColumnHeaders() {
	C.go_fltk_Table_set_column_header((*C.Fl_Table)(t.ptr), 0)
}
func (t *table) AllowColumnResizing() {
	C.go_fltk_Table_set_column_resize((*C.Fl_Table)(t.ptr), 1)
}
func (t *table) DisallowColumnResizing() {
	C.go_fltk_Table_set_column_resize((*C.Fl_Table)(t.ptr), 0)
}
func (t *table) CallbackRow() int {
	return int(C.go_fltk_Table_callback_row((*C.Fl_Table)(t.ptr)))
}
func (t *table) CallbackContext() TableContext {
	return TableContext(C.go_fltk_Table_callback_context((*C.Fl_Table)(t.ptr)))
}
func (t *table) Selection() (int, int, int, int) {
	var top, left, bottom, right C.int
	C.go_fltk_Table_get_selection((*C.Fl_Table)(t.ptr), &top, &left, &bottom, &right)
	return int(top), int(left), int(bottom), int(right)
}
func (t *table) VisibleCells() (int, int, int, int) {
	var top, bottom, left, right C.int
	C.go_fltk_Table_visible_cells((*C.Fl_Table)(t.ptr), &top, &bottom, &left, &right)
	return int(top), int(left), int(bottom), int(right)
}
func (t *table) SetTopRow(row int) {
	C.go_fltk_Table_set_top_row((*C.Fl_Table)(t.ptr), C.int(row))
}

type TableRow struct {
	table
	eventHandler       int
	resizeHandlerId    uintptr
	drawCellCallbackId int
}

type tableCallbackMap struct {
	callbackMap map[int]func(TableContext, int, int, int, int, int, int)
	id          int
}

func newTableCallbackMap() *tableCallbackMap {
	return &tableCallbackMap{
		callbackMap: make(map[int]func(TableContext, int, int, int, int, int, int)),
	}
}
func (m *tableCallbackMap) register(fn func(TableContext, int, int, int, int, int, int)) int {
	m.id++
	m.callbackMap[m.id] = fn
	return m.id
}
func (m *tableCallbackMap) unregister(id int) {
	delete(m.callbackMap, id)
}
func (m *tableCallbackMap) invoke(id int, context TableContext, r, c, x, y, w, h int) {
	if id == 0 {
		return
	}
	if callback, ok := m.callbackMap[id]; ok && callback != nil {
		callback(context, r, c, x, y, w, h)
	}
}

var globalTableCallbackMap = newTableCallbackMap()

type TableContext int

var (
	ContextNone      = TableContext(C.go_FL_CONTEXT_NONE)
	ContextStartPage = TableContext(C.go_FL_CONTEXT_STARTPAGE)
	ContextEndPage   = TableContext(C.go_FL_CONTEXT_ENDPAGE)
	ContextRowHeader = TableContext(C.go_FL_CONTEXT_ROW_HEADER)
	ContextColHeader = TableContext(C.go_FL_CONTEXT_COL_HEADER)
	ContextCell      = TableContext(C.go_FL_CONTEXT_CELL)
	ContextTable     = TableContext(C.go_FL_CONTEXT_TABLE)
	ContextRCResize  = TableContext(C.go_FL_CONTEXT_RC_RESIZE)
)

func NewTableRow(x, y, w, h int) *TableRow {
	i := &TableRow{}
	initWidget(i, unsafe.Pointer(C.go_fltk_new_TableRow(C.int(x), C.int(y), C.int(w), C.int(h))))
	return i
}

func (t *TableRow) Destroy() {
	if t.drawCellCallbackId > 0 {
		globalTableCallbackMap.unregister(t.drawCellCallbackId)
	}
	t.drawCellCallbackId = 0
	if t.resizeHandlerId > 0 {
		globalCallbackMap.unregister(t.resizeHandlerId)
	}
	t.resizeHandlerId = 0
	if t.eventHandler > 0 {
		globalEventHandlerMap.unregister(t.eventHandler)
	}
	t.eventHandler = 0
	t.table.Destroy()
}
func (t *TableRow) IsRowSelected(row int) bool {
	return C.go_fltk_TableRow_row_selected((*C.GTableRow)(t.ptr), C.int(row)) != 0
}
func (t *TableRow) SetDrawCellCallback(callback func(TableContext, int, int, int, int, int, int)) {
	if t.drawCellCallbackId > 0 {
		globalTableCallbackMap.unregister(t.drawCellCallbackId)
	}
	t.drawCellCallbackId = globalTableCallbackMap.register(callback)
	C.go_fltk_TableRow_set_draw_cell_callback((*C.GTableRow)(t.ptr), C.int(t.drawCellCallbackId))
}
func (t *TableRow) SetEventHandler(handler func(Event) bool) {
	if t.eventHandler >= 0 {
		globalEventHandlerMap.unregister(t.eventHandler)
	}
	t.eventHandler = globalEventHandlerMap.register(handler)
	C.go_fltk_TableRow_set_event_handler((*C.GTableRow)(t.ptr), C.int(t.eventHandler))
}
func (t *TableRow) SetResizeHandler(handler func()) {
	if t.resizeHandlerId > 0 {
		globalCallbackMap.unregister(t.resizeHandlerId)
	}
	t.resizeHandlerId = globalCallbackMap.register(handler)
	C.go_fltk_TableRow_set_resize_handler((*C.GTableRow)(t.ptr), C.uintptr_t(t.resizeHandlerId))
}

type SelectionFlag int

var (
	Deselect        = SelectionFlag(C.go_FL_DESELECT)
	Select          = SelectionFlag(C.go_FL_SELECT)
	ToggleSelection = SelectionFlag(C.go_FL_TOGGLE_SELECTION)
)

func (t *TableRow) SelectAllRows(flag SelectionFlag) {
	C.go_fltk_TableRow_select_all_rows((*C.GTableRow)(t.ptr), C.int(flag))
}
func (t *TableRow) SelectRow(row int, flag SelectionFlag) {
	C.go_fltk_TableRow_select_row((*C.GTableRow)(t.ptr), C.int(row), C.int(flag))
}
func (t *TableRow) FindCell(ctx TableContext, row int, col int) (int, int, int, int, error) {
	var x, y, w, h C.int
	ret := C.go_fltk_TableRow_find_cell((*C.GTableRow)(t.ptr), C.int(ctx), C.int(row), C.int(col), &x, &y, &w, &h)
	err := errors.New("No cell was found")
	if ret == 0 {
		err = nil
	}
	return int(x), int(y), int(w), int(h), err
}

type RowSelectMode int

var (
	SelectNone   = RowSelectMode(C.go_FL_SELECT_NONE)
	SelectSingle = RowSelectMode(C.go_FL_SELECT_SINGLE)
	SelectMulti  = RowSelectMode(C.go_FL_SELECT_MULTI)
)

func (t *TableRow) SetType(tableType RowSelectMode) {
	C.go_fltk_TableRow_set_type((*C.GTableRow)(t.ptr), C.int(tableType))
}

//export _go_drawTableHandler
func _go_drawTableHandler(id, context, r, c, x, y, w, h C.int) {
	globalTableCallbackMap.invoke(int(id), TableContext(context), int(r), int(c), int(x), int(y), int(w), int(h))
}
