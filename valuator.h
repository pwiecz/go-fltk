#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Valuator Fl_Valuator;

  extern void go_fltk_Valuator_set_maximum(Fl_Valuator* v, double value);
  extern void go_fltk_Valuator_set_minimum(Fl_Valuator* v, double value);
  extern void go_fltk_Valuator_set_step(Fl_Valuator* v, double value);
  extern double go_fltk_Valuator_value(Fl_Valuator* v);
  extern void go_fltk_Valuator_set_value(Fl_Valuator* v, double value);

#ifdef __cplusplus
}
#endif
