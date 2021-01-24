#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Progress Fl_Progress;

  extern Fl_Progress* go_fltk_new_Progress(int x, int y, int w, int h, const char* label);

  extern void go_fltk_Progress_set_maximum(Fl_Progress* p, double maximum);
  extern void go_fltk_Progress_set_minimum(Fl_Progress* p, double minimum);
  extern void go_fltk_Progress_set_value(Fl_Progress* p, double value);

#ifdef __cplusplus
}
#endif
