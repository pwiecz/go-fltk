#include "table.h"

#include <FL/Fl_Table_Row.H>

#include "event_handler.h"

#include "_cgo_export.h"


class GTableRow : public EventHandler<Fl_Table_Row> {
public:
  GTableRow(int x, int y, int w, int h)
    : EventHandler<Fl_Table_Row>(x, y, w, h) {}
  
  void set_draw_cell_callback(int drawFunId) {
    m_drawFunId = drawFunId;
  }
  void draw_cell(TableContext context, int R, int C, int X, int Y, int W, int H) final {
    if (m_drawFunId > 0) {
      _go_drawTableHandler(m_drawFunId, (int)context, R, C, X, Y, W, H);
    }
  }

  int find_cell_(int ctx, int r, int c, int *x, int *y, int *w, int *h) {
    int X = 0, Y = 0, W = 0, H = 0;
    int ret = this->find_cell((Fl_Table::TableContext)ctx, r, c, X, Y, W, H);
    *x = X, *y = Y, *w = W, *h = H;
    return ret;
  }
	
  int row_from_cursor() {
	int row = 0;
	int col = 0;
	ResizeFlag rflag = RESIZE_NONE;
	TableContext ctx = cursor2rowcol(row, col, rflag);
	if (ctx == CONTEXT_COL_HEADER)
		row = -1;
	else if (ctx != CONTEXT_CELL && ctx != CONTEXT_ROW_HEADER)
		row = -2;
	return row;
  }
	
  int column_from_cursor() {
	int row = 0;
	int col = 0;
	ResizeFlag rflag = RESIZE_NONE;
	TableContext ctx = cursor2rowcol(row, col, rflag);
	if (ctx == CONTEXT_ROW_HEADER)
		col = -1;
	else if (ctx != CONTEXT_CELL && ctx != CONTEXT_COL_HEADER)
		col = -2;
	return col;
  }
	
private:
  int m_drawFunId = 0;
};

GTableRow *go_fltk_new_TableRow(int x, int y, int w, int h) {
  return new GTableRow(x, y, w, h);
}
int go_fltk_TableRow_row_selected(GTableRow* t, int row) {
  return t->row_selected(row);
}
void go_fltk_TableRow_set_draw_cell_callback(GTableRow* t, int drawFunId) {
  t->set_draw_cell_callback(drawFunId);
}
void go_fltk_TableRow_set_type(GTableRow* t, int type) {
  t->type((Fl_Table_Row::TableRowSelectMode)type);
}
void go_fltk_TableRow_select_all_rows(GTableRow* t, int flag) {
  t->select_all_rows(flag);
}
void go_fltk_TableRow_select_row(GTableRow* t, int row, int flag) {
  t->select_row(row, flag);
}
int go_fltk_TableRow_find_cell(GTableRow* t, int ctx, int r, int c, int *x, int *y, int *w, int *h) {
  return t->find_cell_(ctx, r, c, x, y, w, h);
}
void go_fltk_Table_set_row_count(Fl_Table* t, int rowCount) {
  t->rows(rowCount);
}
int go_fltk_Table_row_count(Fl_Table* t) {
	return t->rows();
}
void go_fltk_Table_set_row_height(Fl_Table* t, int row, int height) {
  t->row_height(row, height);
}
void go_fltk_Table_set_row_height_all(Fl_Table* t, int height) {
  t->row_height_all(height);
}
void go_fltk_Table_set_row_header(Fl_Table* t, int header) {
  t->row_header(header);
}
void go_fltk_Table_set_row_resize(Fl_Table* t, int resize) {
  t->row_resize(resize);
}
void go_fltk_Table_set_column_count(Fl_Table* t, int columnCount) {
  t->cols(columnCount);
}
void go_fltk_Table_set_column_width(Fl_Table* t, int column, int width) {
  t->col_width(column, width);
}
void go_fltk_Table_set_column_width_all(Fl_Table* t, int width) {
  t->col_width_all(width);
}
void go_fltk_Table_set_column_header(Fl_Table* t, int header) {
  t->col_header(header);
}
void go_fltk_Table_set_column_resize(Fl_Table* t, int resize) {
  t->col_resize(resize);
}
int go_fltk_Table_callback_row(Fl_Table* t) {
  return t->callback_row();
}
int go_fltk_Table_callback_column(Fl_Table* t) {
  return t->callback_col();
}
int go_fltk_Table_callback_context(Fl_Table* t) {
  return t->callback_context();
}
void go_fltk_Table_get_selection(Fl_Table* t, int* top, int* left, int* bottom, int* right) {
  t->get_selection(*top, *left, *bottom, *right);
}
void go_fltk_Table_visible_cells(Fl_Table* t, int* top, int* bottom, int* left, int* right) {
  t->visible_cells(*top, *bottom, *left, *right);
}
void go_fltk_Table_set_top_row(Fl_Table* t, int row) {
  t->top_row(row);
}
int go_fltk_Table_top_row(Fl_Table* t) {
  return t->top_row();
}
int go_fltk_Table_scrollbar_size(Fl_Table* t) {
  return t->scrollbar_size();
}
void go_fltk_Table_set_scrollbar_size(Fl_Table* t, int size) {
  t->scrollbar_size(size);
}
void go_fltk_Table_set_column_header_height(Fl_Table* t, int size) {
	t->col_header_height(size);
}
int go_fltk_Table_column_header_height(Fl_Table* t) {
	return t->col_header_height();
}
void go_fltk_Table_set_row_header_width(Fl_Table* t, int size) {
	t->row_header_width(size);
}
int go_fltk_Table_row_header_width(Fl_Table* t) {
	return t->row_header_width();
}
int go_fltk_Table_column_from_cursor(GTableRow* t) {
	return t->column_from_cursor();
}
int go_fltk_Table_row_from_cursor(GTableRow* t) {
	return t->row_from_cursor();
}

const int go_FL_CONTEXT_NONE = (int)Fl_Table::CONTEXT_NONE;
const int go_FL_CONTEXT_STARTPAGE = (int)Fl_Table::CONTEXT_STARTPAGE;
const int go_FL_CONTEXT_ENDPAGE = (int)Fl_Table::CONTEXT_ENDPAGE;
const int go_FL_CONTEXT_ROW_HEADER = (int)Fl_Table::CONTEXT_ROW_HEADER;
const int go_FL_CONTEXT_COL_HEADER = (int)Fl_Table::CONTEXT_COL_HEADER;
const int go_FL_CONTEXT_CELL = (int)Fl_Table::CONTEXT_CELL;
const int go_FL_CONTEXT_TABLE = (int)Fl_Table::CONTEXT_TABLE;
const int go_FL_CONTEXT_RC_RESIZE = (int)Fl_Table::CONTEXT_RC_RESIZE;

const int go_FL_SELECT_NONE = (int)Fl_Table_Row::SELECT_NONE;
const int go_FL_SELECT_SINGLE = (int)Fl_Table_Row::SELECT_SINGLE;
const int go_FL_SELECT_MULTI = (int)Fl_Table_Row::SELECT_MULTI;

const int go_FL_DESELECT = 0;
const int go_FL_SELECT = 1;
const int go_FL_TOGGLE_SELECTION = 2;
