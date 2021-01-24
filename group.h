#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Group Fl_Group;
  typedef struct Fl_Widget Fl_Widget;

  extern Fl_Group *go_fltk_new_Group(int x, int y, int w, int h, const char *text);

  extern void go_fltk_Group_begin(Fl_Group *g);
  extern void go_fltk_Group_end(Fl_Group *g);
  extern void go_fltk_Group_add(Fl_Group *g, Fl_Widget *w);
  extern void go_fltk_Group_resizable(Fl_Group *g, Fl_Widget *w);

#ifdef __cplusplus
}
#endif
