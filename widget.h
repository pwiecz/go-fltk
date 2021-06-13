#pragma once

#include <stdint.h>


#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Widget Fl_Widget;

  extern void go_fltk_Widget_destroy(Fl_Widget *w);
  extern void go_fltk_Widget_set_box(Fl_Widget *w, int box);
  extern void go_fltk_Widget_set_labelfont(Fl_Widget *w, int font);
  extern void go_fltk_Widget_set_labelsize(Fl_Widget *w, int size);
  extern void go_fltk_Widget_set_labeltype(Fl_Widget *w, int type);
  extern void go_fltk_Widget_set_callback(Fl_Widget *w, uintptr_t id);
  extern void go_fltk_Widget_when(Fl_Widget* w, int when);
  extern int go_fltk_Widget_x(Fl_Widget *w);
  extern int go_fltk_Widget_y(Fl_Widget *w);
  extern int go_fltk_Widget_w(Fl_Widget *w);
  extern int go_fltk_Widget_h(Fl_Widget *w);
  extern void go_fltk_Widget_set_align(Fl_Widget* w, unsigned int align);
  extern void go_fltk_Widget_measure_label(Fl_Widget* w, int* ww, int *hh);
  extern void go_fltk_Widget_set_position(Fl_Widget* w, int x, int y);
  extern void go_fltk_Widget_resize(Fl_Widget* w, int x, int y, int width, int height);
  extern void go_fltk_Widget_redraw(Fl_Widget* w);
  extern void go_fltk_Widget_deactivate(Fl_Widget* w);
  extern void go_fltk_Widget_activate(Fl_Widget* w);
  extern void go_fltk_Widget_set_type(Fl_Widget* w, unsigned char type);
  extern void go_fltk_Widget_show(Fl_Widget* w);
  extern void go_fltk_Widget_hide(Fl_Widget* w);
  extern unsigned int go_fltk_Widget_selection_color(Fl_Widget* w);
  extern void go_fltk_Widget_set_selection_color(Fl_Widget* w, unsigned int color);
  extern void go_fltk_Widget_set_color(Fl_Widget* w, unsigned int color);
  extern void go_fltk_Widget_set_label(Fl_Widget* w, const char* label);

#ifdef __cplusplus
}
#endif
