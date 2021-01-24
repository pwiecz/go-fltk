#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  extern void go_fltk_color(unsigned int color);
  extern void go_fltk_draw(const char* text, int x, int y, int w, int h, unsigned int align);
  extern void go_fltk_draw_box(int boxType, int x, int y, int w, int h, unsigned int color);


#ifdef __cplusplus
}
#endif
