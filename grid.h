#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GGrid GGrid;
  typedef struct Fl_Widget Fl_Widget;
  typedef struct Fl_Grid Fl_Grid;  

  extern GGrid *go_fltk_new_Grid(int x, int y, int w, int h, const char *text);

  extern void go_fltk_Grid_set_layout(Fl_Grid* grid, int rows, int columns, int margin,
                                      int gap);

  extern void go_fltk_Grid_set_show_grid(Fl_Grid *grid, int show);
  extern void go_fltk_Grid_set_show_grid_and_color(Fl_Grid *grid, int show,
                                                   unsigned int color);

  extern void go_fltk_Grid_set_column_gap(Fl_Grid *grid, int column, int gap);
  extern int go_fltk_Grid_column_gap(Fl_Grid *grid, int column);

  extern void go_fltk_Grid_set_column_weight(Fl_Grid *grid, int column, int weight);
  extern int go_fltk_Grid_column_weight(Fl_Grid* grid, int column);

  extern void go_fltk_Grid_set_row_gap(Fl_Grid *grid, int row, int gap);
  extern int go_fltk_Grid_row_gap(Fl_Grid *grid, int row);

  extern void go_fltk_Grid_set_row_weight(Fl_Grid *grid, int row, int weight);
  extern int go_fltk_Grid_row_weight(Fl_Grid *grid, int row);

  extern void go_fltk_Grid_set_widget(Fl_Grid* grid, Fl_Widget *widget, int row, int column,
                                      int align);
  extern void go_fltk_Grid_set_widget_with_span(Fl_Grid* grid, Fl_Widget *widget, int row,
                                                int column, int rowSpan,
                                                int columnSpan, int align);

  extern const int go_FL_GRID_BOTTOM;
  extern const int go_FL_GRID_BOTTOM_LEFT;
  extern const int go_FL_GRID_BOTTOM_RIGHT;
  extern const int go_FL_GRID_CENTER;
  extern const int go_FL_GRID_FILL;
  extern const int go_FL_GRID_HORIZONTAL;
  extern const int go_FL_GRID_LEFT;
  extern const int go_FL_GRID_PROPORTIONAL;
  extern const int go_FL_GRID_RIGHT;
  extern const int go_FL_GRID_TOP;
  extern const int go_FL_GRID_TOP_LEFT;
  extern const int go_FL_GRID_TOP_RIGHT;
  extern const int go_FL_GRID_VERTICAL;
  
#ifdef __cplusplus
}
#endif
