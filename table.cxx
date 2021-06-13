#include "table.h"

#include <FL/Fl_Table_Row.H>

#include "_cgo_export.h"


class GTableRow : public Fl_Table_Row {
public:
  GTableRow(int x, int y, int w, int h)
    : Fl_Table_Row(x, y, w, h) {}
  
  void set_draw_cell_callback(int drawFunId) {
    m_drawFunId = drawFunId;
  }
  void draw_cell(TableContext context, int R, int C, int X, int Y, int W, int H) final {
    if (m_drawFunId > 0) {
      _go_drawTableHandler(m_drawFunId, (int)context, R, C, X, Y, W, H);
    }
  }
  void set_event_handler(int handlerId) {
    m_eventHandlerId = handlerId;
  }
  int handle(int event) final {
    if (m_eventHandlerId) {
      const int ret = _go_eventHandler(m_eventHandlerId, event);
      if (ret != 0) {
	return ret;
      }
    }
    return Fl_Table_Row::handle(event);
  }

  void set_resize_handler(uintptr_t resizeHandlerId) {
    m_resizeHandlerId = resizeHandlerId;
  }

  void resize(int x, int y, int w, int h) final {
    Fl_Table_Row::resize(x, y, w, h);
    if (m_resizeHandlerId != 0) {
      _go_callbackHandler(m_resizeHandlerId);
    }
  }

private:
  int m_drawFunId = 0;
  int m_eventHandlerId = 0;
  uintptr_t m_resizeHandlerId = 0;
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
void go_fltk_TableRow_set_event_handler(GTableRow* t, int eventHandler) {
  t->set_event_handler(eventHandler);
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
void go_fltk_Table_set_row_count(Fl_Table* t, int rowCount) {
  t->rows(rowCount);
}
void go_fltk_Table_set_column_count(Fl_Table* t, int columnCount) {
  t->cols(columnCount);
}
void go_fltk_Table_set_column_width(Fl_Table* t, int column, int width) {
  t->col_width(column, width);
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

void go_fltk_TableRow_set_resize_handler(GTableRow* t, uintptr_t handlerId) {
  t->set_resize_handler(handlerId);
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
