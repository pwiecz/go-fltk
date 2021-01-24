#include "button.h"

#include <Fl/Fl_Button.H>

Fl_Button *go_fltk_new_Button(int x, int y, int w, int h, const char *label) {
  return new Fl_Button(x, y, w, h, label);
}

char go_fltk_Button_value(Fl_Button *b) {
  return b->value();
}

