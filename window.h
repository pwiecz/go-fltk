#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Window Fl_Window;

  extern Fl_Window *go_fltk_new_Window(int w, int h);

  extern int go_fltk_Window_shown(Fl_Window* w);
  extern void go_fltk_Window_show(Fl_Window *w);
  extern void go_fltk_Window_set_label(Fl_Window *w, const char *label);
  extern void go_fltk_Window_set_cursor(Fl_Window *w, int cursor);

  extern const int go_FL_CURSOR_DEFAULT;
  extern const int go_FL_CURSOR_ARROW;
  extern const int go_FL_CURSOR_CROSS;
  extern const int go_FL_CURSOR_WAIT;
  extern const int go_FL_CURSOR_INSERT;
  extern const int go_FL_CURSOR_HAND;
  extern const int go_FL_CURSOR_HELP;
  extern const int go_FL_CURSOR_MOVE;
  extern const int go_FL_CURSOR_NS;
  extern const int go_FL_CURSOR_WE;
  extern const int go_FL_CURSOR_NWSE;
  extern const int go_FL_CURSOR_NESW;
  extern const int go_FL_CURSOR_N;
  extern const int go_FL_CURSOR_NE;
  extern const int go_FL_CURSOR_E;
  extern const int go_FL_CURSOR_SE;
  extern const int go_FL_CURSOR_S;
  extern const int go_FL_CURSOR_SW;
  extern const int go_FL_CURSOR_W;
  extern const int go_FL_CURSOR_NW;
  extern const int go_FL_CURSOR_NONE;

#ifdef __cplusplus
}
#endif
