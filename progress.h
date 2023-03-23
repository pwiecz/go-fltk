#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Progress Fl_Progress;
  typedef struct GProgress GProgress;

  extern GProgress* go_fltk_new_Progress(int x, int y, int w, int h, const char* label);

  extern void go_fltk_Progress_set_maximum(Fl_Progress* p, double maximum);
  extern double go_fltk_Progress_maximum(Fl_Progress* p);
  extern void go_fltk_Progress_set_minimum(Fl_Progress* p, double minimum);
  extern double go_fltk_Progress_minimum(Fl_Progress* p);
  extern void go_fltk_Progress_set_value(Fl_Progress* p, double value);
  extern double go_fltk_Progress_value(Fl_Progress* p);

#ifdef __cplusplus
}
#endif
