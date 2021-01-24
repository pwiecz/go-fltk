#include "progress.h"

#include <Fl/Fl_Progress.H>

Fl_Progress* go_fltk_new_Progress(int x, int y, int w, int h, const char *label) {
  return new Fl_Progress(x, y, w, h, label);
}

void go_fltk_Progress_set_maximum(Fl_Progress* p, double maximum) {
  p->maximum((float)maximum);
}
void go_fltk_Progress_set_minimum(Fl_Progress* p, double minimum) {
  p->minimum((float)minimum);
}
void go_fltk_Progress_set_value(Fl_Progress* p, double value) {
  p->value((float)value);
}

