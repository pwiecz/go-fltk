#pragma once

#ifdef __cplusplus
extern "C" {
#endif
  typedef struct Fl_Widget Fl_Widget;

  extern void go_fltk_enable_tooltips();

  extern void go_fltk_disable_tooltips();

  extern int go_fltk_tooltips_enabled();

  extern float go_fltk_tooltip_delay();

  extern void go_fltk_set_tooltip_delay(float delay);

  extern void go_fltk_tooltip_enter_area(Fl_Widget *wi, int x, int y, int w, int h, const char *tip);
#ifdef __cplusplus
}
#endif
