#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Scroll Fl_Scroll;
  typedef struct GScroll GScroll;

  extern GScroll *go_fltk_new_Scroll(int x, int y, int w, int h, const char *text);
  extern void go_fltk_Scroll_scroll_to(Fl_Scroll* scroll, int x, int y);
  extern int go_fltk_Scroll_x_position(Fl_Scroll* scroll);
  extern int go_fltk_Scroll_y_position(Fl_Scroll* scroll);

  extern const unsigned char go_FL_SCROLL_HORIZONTAL;
  extern const unsigned char go_FL_SCROLL_VERTICAL;
  extern const unsigned char go_FL_SCROLL_BOTH;
  extern const unsigned char go_FL_SCROLL_HORIZONTAL_ALWAYS;
  extern const unsigned char go_FL_SCROLL_VERTICAL_ALWAYS;
  extern const unsigned char go_FL_SCROLL_BOTH_ALWAYS;

#ifdef __cplusplus
}
#endif
