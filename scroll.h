#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GScroll GScroll;

  extern GScroll *go_fltk_new_Scroll(int x, int y, int w, int h, const char *text);

  extern const unsigned char go_FL_SCROLL_HORIZONTAL;
  extern const unsigned char go_FL_SCROLL_VERTICAL;
  extern const unsigned char go_FL_SCROLL_BOTH;
  extern const unsigned char go_FL_SCROLL_HORIZONTAL_ALWAYS;
  extern const unsigned char go_FL_SCROLL_VERTICAL_ALWAYS;
  extern const unsigned char go_FL_SCROLL_BOTH_ALWAYS;

#ifdef __cplusplus
}
#endif
