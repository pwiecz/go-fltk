#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GGroup GGroup;
  typedef struct Fl_Widget Fl_Widget;

  extern GGroup *go_fltk_new_Group(int x, int y, int w, int h, const char *text);

  extern void go_fltk_Group_begin(GGroup *g);
  extern void go_fltk_Group_end(GGroup *g);
  extern void go_fltk_Group_add(GGroup *g, Fl_Widget *w);
  extern void go_fltk_Group_resizable(GGroup *g, Fl_Widget *w);
  extern void go_fltk_Group_draw_children(GGroup *g);

#ifdef __cplusplus
}
#endif
