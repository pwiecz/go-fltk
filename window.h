#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Window Fl_Window;

  extern Fl_Window *go_fltk_new_Window(int w, int h);

  extern void go_fltk_Window_show(Fl_Window *w);
  extern void go_fltk_Window_set_label(Fl_Window *w, const char *label);

#ifdef __cplusplus
}
#endif
