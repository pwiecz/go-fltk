#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GProgress GProgress;

  extern GProgress* go_fltk_new_Progress(int x, int y, int w, int h, const char* label);

  extern void go_fltk_Progress_set_maximum(GProgress* p, double maximum);
  extern void go_fltk_Progress_set_minimum(GProgress* p, double minimum);
  extern void go_fltk_Progress_set_value(GProgress* p, double value);

#ifdef __cplusplus
}
#endif
