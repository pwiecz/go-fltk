#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Spinner Fl_Spinner;
  typedef struct GSpinner GSpinner;

  extern GSpinner* go_fltk_new_Spinner(int x, int y, int w, int h, const char* label);
  extern void go_fltk_Spinner_set_type(Fl_Spinner* s, unsigned char type);
  extern void go_fltk_Spinner_set_maximum(Fl_Spinner* s, double max);
  extern void go_fltk_Spinner_set_minimum(Fl_Spinner* s, double min);
  extern void go_fltk_Spinner_set_step(Fl_Spinner* s, double step);
  extern void go_fltk_Spinner_set_value(Fl_Spinner* s, double value);
  extern double go_fltk_Spinner_value(Fl_Spinner* s);

#ifdef __cplusplus
}
#endif
