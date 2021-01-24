#include "table.h"

#include <Fl/Fl_Table_Row.H>

#include "_cgo_export.h"


class GTableRow : public Fl_Table_Row {
public:
  GTableRow(int x, int y, int w, int h, int drawFunId)
      : Fl_Table_Row(x, y, w, h)
      , m_drawFunId(drawFunId) {}
  void draw_cell(TableContext context, int R, int C, int X, int Y, int W,
                 int H) override {
    _go_drawTableHandler(m_drawFunId, (int)context, R, C, X, Y, W, H);
  }

private:
  const int m_drawFunId;
};

Fl_Table_Row *go_fltk_new_TableRow(int x, int y, int w, int h, int drawFunId) {
  return new GTableRow(x, y, w, h, drawFunId);
}

void go_fltk_Table_set_row_count(Fl_Table* t, int rowCount) {
  t->rows(rowCount);
}
void go_fltk_Table_set_column_count(Fl_Table* t, int columnCount) {
  t->cols(columnCount);
}

const int go_FL_CONTEXT_NONE = Fl_Table::CONTEXT_NONE;
const int go_FL_CONTEXT_STARTPAGE = Fl_Table::CONTEXT_STARTPAGE;
const int go_FL_CONTEXT_ENDPAGE = Fl_Table::CONTEXT_ENDPAGE;
const int go_FL_CONTEXT_ROW_HEADER = Fl_Table::CONTEXT_ROW_HEADER;
const int go_FL_CONTEXT_COL_HEADER = Fl_Table::CONTEXT_COL_HEADER;
const int go_FL_CONTEXT_CELL = Fl_Table::CONTEXT_CELL;
const int go_FL_CONTEXT_TABLE = Fl_Table::CONTEXT_TABLE;
const int go_FL_CONTEXT_RC_RESIZE = Fl_Table::CONTEXT_RC_RESIZE;
