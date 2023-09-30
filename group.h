#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GGroup GGroup;
  typedef struct Fl_Widget Fl_Widget;
  typedef struct Fl_Group Fl_Group;

  extern GGroup *go_fltk_new_Group(int x, int y, int w, int h, const char *text);

  extern void go_fltk_Group_begin(Fl_Group *g);
  extern void go_fltk_Group_end(Fl_Group *g);
  extern void go_fltk_Group_add(Fl_Group *g, Fl_Widget *w);
  extern void go_fltk_Group_remove(Fl_Group *g, Fl_Widget *w);
  extern void go_fltk_Group_resizable(Fl_Group *g, Fl_Widget *w);
  extern void go_fltk_Group_draw_children(Fl_Group *g);
  extern Fl_Widget* go_fltk_Group_child(Fl_Group *g, int index);
  extern int go_fltk_Group_child_count(Fl_Group* g);

#ifdef __cplusplus
}
#endif
