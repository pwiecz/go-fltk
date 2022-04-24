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

void go_fltk_Progress_set_maximum(GProgress* p, double maximum) {
  p->maximum((float)maximum);
}
double go_fltk_Progress_maximum(GProgress* p) {
  return (double)p->maximum();
}
void go_fltk_Progress_set_minimum(GProgress* p, double minimum) {
  p->minimum((float)minimum);
}
double go_fltk_Progress_minimum(GProgress* p) {
  return (double)p->minimum();
}
void go_fltk_Progress_set_value(GProgress* p, double value) {
  p->value((float)value);
}
double go_fltk_Progress_value(GProgress* p) {
  return (double)p->value();
}

