#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GProgress GProgress;

  extern GProgress* go_fltk_new_Progress(int x, int y, int w, int h, const char* label);

  extern void go_fltk_Progress_set_maximum(GProgress* p, double maximum);
  extern double go_fltk_Progress_maximum(GProgress* p);
  extern void go_fltk_Progress_set_minimum(GProgress* p, double minimum);
  extern double go_fltk_Progress_minimum(GProgress* p);
  extern void go_fltk_Progress_set_value(GProgress* p, double value);
  extern double go_fltk_Progress_value(GProgress* p);

#ifdef __cplusplus
}
#endif
