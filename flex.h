#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GFlex GFlex;
  typedef struct Fl_Widget Fl_Widget;

  extern GFlex *go_fltk_new_Flex(int x, int y, int w, int h, const char *text);

  extern void go_fltk_Flex_set_gap(GFlex* flex, int spacing);

  extern void go_fltk_Flex_fixed(GFlex *flex, Fl_Widget *w, int size);

  extern void go_fltk_Flex_end(GFlex *flex);

  extern const unsigned char go_FL_FLEX_COLUMN;
  extern const unsigned char go_FL_FLEX_ROW;

#ifdef __cplusplus
}
#endif
