#include "grid.h"

#include <FL/Fl_Grid.H>

#include "event_handler.h"


class GGrid : public EventHandler<Fl_Grid> {
public:
  GGrid(int x, int y, int w, int h, const char* label)
  : EventHandler<Fl_Grid>(x, y, w, h, label) {}
};

GGrid* go_fltk_new_Grid(int x, int y, int w, int h, const char* label) {
  return new GGrid(x, y, w, h, label);
}

void go_fltk_Grid_set_layout(Fl_Grid* grid, int rows, int columns, int margin,
                             int gap) {
  grid->layout(rows, columns, margin, gap);
}

void go_fltk_Grid_set_show_grid(Fl_Grid *grid, int show) {
  grid->show_grid(show);
}
void go_fltk_Grid_set_show_grid_and_color(Fl_Grid *grid, int show,
                                          unsigned int color) {
  grid->show_grid(show, (Fl_Color)color);
}

void go_fltk_Grid_set_column_gap(Fl_Grid *grid, int column, int gap) {
  grid->col_gap(column, gap);
}
int go_fltk_Grid_column_gap(Fl_Grid *grid, int column) {
  return grid->col_gap(column);
}

void go_fltk_Grid_set_column_weight(Fl_Grid *grid, int column, int weight) {
  grid->col_weight(column, weight);
}
int go_fltk_Grid_column_weight(Fl_Grid* grid, int column) {
  return grid->col_weight(column);
}

void go_fltk_Grid_set_row_gap(Fl_Grid *grid, int row, int gap) {
  grid->row_gap(row, gap);
}
int go_fltk_Grid_row_gap(Fl_Grid *grid, int row) {
  return grid->row_gap(row);
}

void go_fltk_Grid_set_row_weight(Fl_Grid *grid, int row, int weight) {
  grid->row_weight(row, weight);
}
int go_fltk_Grid_row_weight(Fl_Grid *grid, int row) {
  return grid->row_weight(row);
}

void go_fltk_Grid_set_widget(Fl_Grid* grid, Fl_Widget *widget, int row, int column, int align) {
  grid->widget(widget, row, column, (Fl_Grid_Align)align);
}

void go_fltk_Grid_set_widget_with_span(Fl_Grid* grid, Fl_Widget *widget, int row,
                                       int column, int rowSpan,
                                       int columnSpan, int align) {
  grid->widget(widget, row, column, rowSpan, columnSpan, (Fl_Grid_Align)align);
}

const int go_FL_GRID_BOTTOM = (int)FL_GRID_BOTTOM;
const int go_FL_GRID_BOTTOM_LEFT = (int)FL_GRID_BOTTOM_LEFT;
const int go_FL_GRID_BOTTOM_RIGHT = (int)FL_GRID_BOTTOM_RIGHT;
const int go_FL_GRID_CENTER = (int)FL_GRID_CENTER;
const int go_FL_GRID_FILL = (int)FL_GRID_FILL;
const int go_FL_GRID_HORIZONTAL = (int)FL_GRID_HORIZONTAL;
const int go_FL_GRID_LEFT = (int)FL_GRID_LEFT;
const int go_FL_GRID_PROPORTIONAL = (int)FL_GRID_PROPORTIONAL;
const int go_FL_GRID_RIGHT = (int)FL_GRID_RIGHT;
const int go_FL_GRID_TOP = (int)FL_GRID_TOP;
const int go_FL_GRID_TOP_LEFT = (int)FL_GRID_TOP_LEFT;
const int go_FL_GRID_TOP_RIGHT = (int)FL_GRID_TOP_RIGHT;
const int go_FL_GRID_VERTICAL = (int)FL_GRID_VERTICAL;
