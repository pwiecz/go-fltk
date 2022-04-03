#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GSpinner GSpinner;

  extern GSpinner* go_fltk_new_Spinner(int x, int y, int w, int h, const char* label);
  extern void go_fltk_Spinner_set_type(GSpinner* s, unsigned char type);
  extern void go_fltk_Spinner_set_maximum(GSpinner* s, double max);
  extern void go_fltk_Spinner_set_minimum(GSpinner* s, double min);
  extern void go_fltk_Spinner_set_step(GSpinner* s, double step);
  extern void go_fltk_Spinner_set_value(GSpinner* s, double value);
  extern double go_fltk_Spinner_value(GSpinner* s);

#ifdef __cplusplus
}
#endif
