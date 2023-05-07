#pragma once

#include <stdint.h>


#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Table Fl_Table;
  typedef struct GTableRow GTableRow;

  extern GTableRow* go_fltk_new_TableRow(int x, int y, int w, int h);

  extern void go_fltk_Table_set_row_count(Fl_Table* t, int rowCount);
  extern int go_fltk_Table_row_count(Fl_Table* t);
  extern void go_fltk_Table_set_row_height(Fl_Table* t, int row, int height);
  extern void go_fltk_Table_set_row_height_all(Fl_Table* t, int height);
  extern void go_fltk_Table_set_row_header(Fl_Table* t, int header);
  extern void go_fltk_Table_set_row_resize(Fl_Table* t, int resize);
  extern void go_fltk_Table_set_column_count(Fl_Table* t, int columnCount);
  extern void go_fltk_Table_set_column_width(Fl_Table* t, int column, int width);
  extern void go_fltk_Table_set_column_width_all(Fl_Table* t, int width);
  extern void go_fltk_Table_set_column_header(Fl_Table* t, int header);
  extern void go_fltk_Table_set_column_resize(Fl_Table* t, int resize);
  extern int go_fltk_Table_callback_row(Fl_Table* t);
  extern int go_fltk_Table_callback_column(Fl_Table* t);
  extern int go_fltk_Table_callback_context(Fl_Table* t);
  extern void go_fltk_Table_get_selection(Fl_Table* t, int* top, int* left, int* bottom, int* right);
  extern void go_fltk_Table_visible_cells(Fl_Table* t, int* top, int* bottom, int* left, int* right);
  extern void go_fltk_Table_set_top_row(Fl_Table* t, int row);
  extern int go_fltk_Table_top_row(Fl_Table* t);
  extern int go_fltk_Table_scrollbar_size(Fl_Table* t);
  extern void go_fltk_Table_set_scrollbar_size(Fl_Table* t, int size);
  extern int go_fltk_Table_row_header_width(Fl_Table* t);
  extern void go_fltk_Table_set_row_header_width(Fl_Table* t, int size);
  extern int go_fltk_Table_column_header_height(Fl_Table* t);
  extern void go_fltk_Table_set_column_header_height(Fl_Table* t, int size);
		
  extern int go_fltk_TableRow_row_selected(GTableRow* t, int row);
  extern void go_fltk_TableRow_set_draw_cell_callback(GTableRow* t, int drawCellCallback);
  extern void go_fltk_TableRow_set_type(GTableRow* t, int tableType);
  extern void go_fltk_TableRow_select_all_rows(GTableRow* t, int flag);
  extern void go_fltk_TableRow_select_row(GTableRow* t, int row, int flag);
  extern int go_fltk_TableRow_find_cell(GTableRow* t, int ctx, int row, int col, int *x, int *y, int *w, int *h);
  extern int go_fltk_Table_column_from_cursor(GTableRow* t);
  extern int go_fltk_Table_row_from_cursor(GTableRow* t);
		
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

  extern const int go_FL_DESELECT;
  extern const int go_FL_SELECT;
  extern const int go_FL_TOGGLE_SELECTION;

#ifdef __cplusplus
}
#endif
