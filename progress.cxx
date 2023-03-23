#include "progress.h"

#include <FL/Fl_Progress.H>

#include "event_handler.h"


class GProgress : public EventHandler<Fl_Progress> {
public:
  GProgress(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Progress>(x, y, w, h, label) {}
};

GProgress* go_fltk_new_Progress(int x, int y, int w, int h, const char *label) {
  return new GProgress(x, y, w, h, label);
}

void go_fltk_Progress_set_maximum(Fl_Progress* p, double maximum) {
  p->maximum((float)maximum);
}
double go_fltk_Progress_maximum(Fl_Progress* p) {
  return (double)p->maximum();
}
void go_fltk_Progress_set_minimum(Fl_Progress* p, double minimum) {
  p->minimum((float)minimum);
}
double go_fltk_Progress_minimum(Fl_Progress* p) {
  return (double)p->minimum();
}
void go_fltk_Progress_set_value(Fl_Progress* p, double value) {
  p->value((float)value);
}
double go_fltk_Progress_value(Fl_Progress* p) {
  return (double)p->value();
}

