#include "tabs.h"

#include "Fl/Fl_Tabs.H"


Fl_Tabs *go_fltk_new_Tabs(int x, int y, int w, int h, const char *label) {
  return new Fl_Tabs(x, y, w, h, label);
}

