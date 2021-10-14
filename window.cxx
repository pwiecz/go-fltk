#include "window.h"

#include "FL/Fl_Window.H"


Fl_Window *go_fltk_new_Window(int w, int h) {
  return new Fl_Window(w, h);
}

int go_fltk_Window_shown(Fl_Window *w) {
  return w->shown();
}

void go_fltk_Window_show(Fl_Window *w) {
  w->show();
}

int go_fltk_Window_x_root(Fl_Window* w) {
  return w->x_root();
}

int go_fltk_Window_y_root(Fl_Window* w) {
  return w->y_root();
}

void go_fltk_Window_set_label(Fl_Window *w, const char *label) {
  w->copy_label(label);
}

void go_fltk_Window_set_cursor(Fl_Window* w, int cursor) {
  w->cursor((Fl_Cursor)cursor);
}

const int go_FL_CURSOR_DEFAULT = (int)FL_CURSOR_DEFAULT;
const int go_FL_CURSOR_ARROW = (int)FL_CURSOR_ARROW;
const int go_FL_CURSOR_CROSS = (int)FL_CURSOR_CROSS;
const int go_FL_CURSOR_WAIT = (int)FL_CURSOR_WAIT;
const int go_FL_CURSOR_INSERT = (int)FL_CURSOR_INSERT;
const int go_FL_CURSOR_HAND = (int)FL_CURSOR_HAND;
const int go_FL_CURSOR_HELP = (int)FL_CURSOR_HELP;
const int go_FL_CURSOR_MOVE = (int)FL_CURSOR_MOVE;
const int go_FL_CURSOR_NS = (int)FL_CURSOR_NS;
const int go_FL_CURSOR_WE = (int)FL_CURSOR_WE;
const int go_FL_CURSOR_NWSE = (int)FL_CURSOR_NWSE;
const int go_FL_CURSOR_NESW = (int)FL_CURSOR_NESW;
const int go_FL_CURSOR_N = (int)FL_CURSOR_N;
const int go_FL_CURSOR_NE = (int)FL_CURSOR_NE;
const int go_FL_CURSOR_E = (int)FL_CURSOR_E;
const int go_FL_CURSOR_SE = (int)FL_CURSOR_SE;
const int go_FL_CURSOR_S = (int)FL_CURSOR_S;
const int go_FL_CURSOR_SW = (int)FL_CURSOR_SW;
const int go_FL_CURSOR_W = (int)FL_CURSOR_W;
const int go_FL_CURSOR_NW = (int)FL_CURSOR_NW;
const int go_FL_CURSOR_NONE = (int)FL_CURSOR_NONE;
