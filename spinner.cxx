#include "spinner.h"

#include <FL/Fl_Spinner.H>

#include "event_handler.h"


class GSpinner : public EventHandler<Fl_Spinner> {
public:
  GSpinner(int x, int y, int w, int h, const char* label)
    : EventHandler<Fl_Spinner>(x, y, w, h, label) {}
};

GSpinner* go_fltk_new_Spinner(int x, int y, int w, int h, const char* label) {
  return new GSpinner(x, y, w, h, label);
}

void go_fltk_Spinner_set_type(GSpinner* s, unsigned char type) {
  s->type(type);
}
void go_fltk_Spinner_set_maximum(GSpinner* s, double max) {
  s->maximum(max);
}
void go_fltk_Spinner_set_minimum(GSpinner* s, double min) {
  s->minimum(min);
}
void go_fltk_Spinner_set_step(GSpinner* s, double step) {
  s->step(step);
}
void go_fltk_Spinner_set_value(GSpinner* s, double value) {
  s->value(value);
}
double go_fltk_Spinner_value(GSpinner*s) {
  return s->value();
}
