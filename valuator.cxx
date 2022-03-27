#include "valuator.h"

#include <FL/Fl_Valuator.H>


void go_fltk_Valuator_set_minimum(Fl_Valuator* v, double value) {
  v->minimum(value);
}

void go_fltk_Valuator_set_maximum(Fl_Valuator* v, double value) {
  v->maximum(value);
}

void go_fltk_Valuator_set_step(Fl_Valuator* v, double value) {
  v->step(value);
}

double go_fltk_Valuator_value(Fl_Valuator* v) {
  return v->value();
}

void go_fltk_Valuator_set_value(Fl_Valuator* v, double value) {
  v->value(value);
}
