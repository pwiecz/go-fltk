#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Choice Fl_Choice;

  extern Fl_Choice *go_fltk_new_Choice(int x, int y, int w, int h, const char *label);

#ifdef __cplusplus
}
#endif
