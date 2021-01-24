package fltk

/*
#include "table.h"
*/
import "C"
import "unsafe"

type table struct {
	Group
}

func (t *table) SetRowCount(rowCount int) {
	C.go_fltk_Table_set_row_count((*C.Fl_Table)(t.ptr), C.int(rowCount))
}
func (t *table) SetColumnCount(columnCount int) {
	C.go_fltk_Table_set_column_count((*C.Fl_Table)(t.ptr), C.int(columnCount))
}

type TableRow struct {
	table
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
	m.callbackMap[id](context, r, c, x, y, w, h)
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

func NewTableRow(x, y, w, h int, drawFun func(TableContext, int, int, int, int, int, int)) *TableRow {
	i := &TableRow{}
	drawFunId := globalTableCallbackMap.register(drawFun)
	initWidget(i, unsafe.Pointer(C.go_fltk_new_TableRow(C.int(x), C.int(y), C.int(w), C.int(h), C.int(drawFunId))))
	return i
}

//export _go_drawTableHandler
func _go_drawTableHandler(id, context, r, c, x, y, w, h C.int) {
	globalTableCallbackMap.invoke(int(id), TableContext(context), int(r), int(c), int(x), int(y), int(w), int(h))
}
