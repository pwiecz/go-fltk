#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GWindow GWindow;
  typedef struct Fl_RGB_Image Fl_RGB_Image;

  extern GWindow *go_fltk_new_Window(int w, int h);

  extern int go_fltk_Window_shown(GWindow* w);
  extern void go_fltk_Window_show(GWindow *w);
  extern int go_fltk_Window_x_root(GWindow* w);
  extern int go_fltk_Window_y_root(GWindow* w);
  extern void go_fltk_Window_set_label(GWindow *w, const char *label);
  extern void go_fltk_Window_set_cursor(GWindow *w, int cursor);
  extern void go_fltk_Window_set_fullscreen(GWindow *w, int flag);
  extern int go_fltk_Window_fullscreen_active(GWindow *w);
  extern void go_fltk_Window_set_modal(GWindow *w);
  extern void go_fltk_Window_set_non_modal(GWindow *w);
  extern void go_fltk_Window_set_icons(GWindow* w, const Fl_RGB_Image* images[], int length);

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
