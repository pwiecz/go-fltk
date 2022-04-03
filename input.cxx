#include "input.h"

#include <FL/Fl_Input.H>

#include "event_handler.h"


const unsigned char go_FL_FLOAT_INPUT = 1;
const unsigned char go_FL_INT_INPUT = 2;

class GInput : public EventHandler<Fl_Input> {
public:
  GInput(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Input>(x, y, w, h, label) {}
};

GInput *go_fltk_new_Input(int x, int y, int w, int h, const char *text) {
  return new GInput(x, y, w, h, text);
}

const char *go_fltk_Input_value(GInput *in) {
  return in->value();
}
int go_fltk_Input_set_value(GInput *in, const char *t) {
  return in->value(t);
}
void go_fltk_Input_resize(GInput *in, int x, int y, int w, int h) {
  in->resize(x, y, w, h);
}
