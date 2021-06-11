#include "tabs.h"

#include "Fl/Fl_Tabs.H"


Fl_Tabs *go_fltk_new_Tabs(int x, int y, int w, int h, const char *label) {
  return new Fl_Tabs(x, y, w, h, label);
}

int go_fltk_Tabs_value(Fl_Tabs* tabs) {
  const Fl_Widget* widget = tabs->value();
  const int childrenCount = tabs->children();
  for (int i = 0; i < childrenCount; ++i) {
    if (tabs->child(i) == widget) {
      return i;
    }
  }
  return -1;
}

void go_fltk_Tabs_set_value(Fl_Tabs* tabs, int value) {
  if (value < 0 || value >= tabs->children()) {
    return;
  }
  tabs->value(tabs->child(value));
}
