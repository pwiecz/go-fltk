#include "input.h"

#include <FL/Fl_Input.H>


const unsigned char go_FL_FLOAT_INPUT = 1;
const unsigned char go_FL_INT_INPUT = 2;

Fl_Input *go_fltk_new_Input(int x, int y, int w, int h, const char *text) {
  return new Fl_Input(x, y, w, h, text);
}

const char *go_fltk_Input_value(Fl_Input *in) {
  return in->value();
}
int go_fltk_Input_set_value(Fl_Input *in, const char *t) {
  return in->value(t);
}
void go_fltk_Input_resize(Fl_Input *in, int x, int y, int w, int h) {
  in->resize(x, y, w, h);
}
