#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Table Fl_Table;
  typedef struct GTableRow GTableRow;

  extern GTableRow* go_fltk_new_TableRow(int x, int y, int w, int h);

  extern void go_fltk_Table_set_row_count(Fl_Table* t, int rowCount);
  extern void go_fltk_Table_set_column_count(Fl_Table* t, int columnCount);
  extern void go_fltk_Table_set_column_width(Fl_Table* t, int column, int width);
  extern void go_fltk_Table_set_column_header(Fl_Table* t, int header);
  extern void go_fltk_Table_set_column_resize(Fl_Table* t, int resize);
  extern int go_fltk_TableRow_row_selected(GTableRow* t, int row);
  extern void go_fltk_TableRow_set_draw_cell_callback(GTableRow* t, int drawCellCallback);
  extern void go_fltk_TableRow_set_event_handler(GTableRow* t, int eventHandler);
  extern void go_fltk_TableRow_set_type(GTableRow* t, int tableType);

  extern const int go_FL_CONTEXT_NONE;
  extern const int go_FL_CONTEXT_STARTPAGE;
  extern const int go_FL_CONTEXT_ENDPAGE;
  extern const int go_FL_CONTEXT_ROW_HEADER;
  extern const int go_FL_CONTEXT_COL_HEADER;
  extern const int go_FL_CONTEXT_CELL;
  extern const int go_FL_CONTEXT_TABLE;
  extern const int go_FL_CONTEXT_RC_RESIZE;

  extern const int go_FL_SELECT_NONE;
  extern const int go_FL_SELECT_SINGLE;
  extern const int go_FL_SELECT_MULTI;

#ifdef __cplusplus
}
#endif
