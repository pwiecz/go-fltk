#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GFlex GFlex;
  typedef struct Fl_Flex Fl_Flex;
  typedef struct Fl_Widget Fl_Widget;

  extern GFlex *go_fltk_new_Flex(int x, int y, int w, int h, const char *text);

  extern void go_fltk_Flex_set_gap(Fl_Flex* flex, int spacing);

  extern void go_fltk_Flex_fixed(Fl_Flex *flex, Fl_Widget *w, int size);

  extern void go_fltk_Flex_end(Fl_Flex *flex);

  extern int go_fltk_Flex_spacing(Fl_Flex* flex);

  extern void go_fltk_Flex_set_spacing(Fl_Flex* flex, int spacing);

  extern int go_fltk_Flex_margin(Fl_Flex* flex);

  extern void go_fltk_Flex_set_margin(Fl_Flex* flex, int margin, int gap);

  extern void go_fltk_Flex_layout(Fl_Flex *flex);
 
  extern const unsigned char go_FL_FLEX_COLUMN;
  extern const unsigned char go_FL_FLEX_ROW;

#ifdef __cplusplus
}
#endif
