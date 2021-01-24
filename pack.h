#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Pack Fl_Pack;

  extern Fl_Pack *go_fltk_new_Pack(int x, int y, int w, int h, const char *text);

#ifdef __cplusplus
}
#endif
