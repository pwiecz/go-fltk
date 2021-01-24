#include "spinner.h"

#include <Fl/Fl_Spinner.H>

Fl_Spinner* go_fltk_new_Spinner(int x, int y, int w, int h, const char* label) {
  return new Fl_Spinner(x, y, w, h, label);
}

void go_fltk_Spinner_set_type(Fl_Spinner* s, unsigned char type) {
  s->type(type);
}
void go_fltk_Spinner_set_maximum(Fl_Spinner* s, double max) {
  s->maximum(max);
}
void go_fltk_Spinner_set_minimum(Fl_Spinner* s, double min) {
  s->minimum(min);
}
void go_fltk_Spinner_set_step(Fl_Spinner* s, double step) {
  s->step(step);
}
void go_fltk_Spinner_set_value(Fl_Spinner* s, double value) {
  s->value(value);
}
double go_fltk_Spinner_value(Fl_Spinner*s) {
  return s->value();
}
