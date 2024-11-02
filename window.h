#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Window Fl_Window;
  typedef struct GWindow GWindow;
  typedef struct Fl_RGB_Image Fl_RGB_Image;

  extern GWindow *go_fltk_new_Window(int w, int h, const char* title);
  extern GWindow *go_fltk_new_Window_with_position(int x, int y, int w, int h, const char* title);

  extern int go_fltk_Window_shown(Fl_Window* w);
  extern void go_fltk_Window_show(Fl_Window *w);
  extern int go_fltk_Window_x_root(Fl_Window* w);
  extern int go_fltk_Window_y_root(Fl_Window* w);
  extern void go_fltk_Window_set_xclass(Fl_Window* w, const char *wmclass);
  extern void go_fltk_Window_set_label(Fl_Window *w, const char *label);
  extern void go_fltk_Window_set_cursor(Fl_Window *w, int cursor);
  extern void go_fltk_Window_set_fullscreen(Fl_Window *w, int flag);
  extern int go_fltk_Window_fullscreen_active(Fl_Window *w);
  extern void go_fltk_Window_set_modal(Fl_Window *w);
  extern void go_fltk_Window_set_non_modal(Fl_Window *w);
  extern void go_fltk_Window_set_icons(Fl_Window* w, const Fl_RGB_Image* images[], int length);
  extern void go_fltk_Window_size_range(Fl_Window* w, int minW, int minH, int maxW, int maxH, int deltaX, int deltaY, int aspectRatio);
  extern void* go_fltk_Window_xid(Fl_Window* w);
		
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
