#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Box Fl_Box;

  extern Fl_Box *go_fltk_new_Box(int boxType, int x, int y, int w, int h, const char *label);

#ifdef __cplusplus
}
#endif
