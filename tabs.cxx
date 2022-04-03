#include "tabs.h"

#include <FL/Fl_Tabs.H>

#include "event_handler.h"


class GTabs : public EventHandler<Fl_Tabs> {
public:
  GTabs(int x, int y, int w, int h, const char* label)
    : EventHandler<Fl_Tabs>(x, y, w, h, label) {}
};

GTabs *go_fltk_new_Tabs(int x, int y, int w, int h, const char *label) {
  return new GTabs(x, y, w, h, label);
}

int go_fltk_Tabs_value(GTabs* tabs) {
  const Fl_Widget* widget = tabs->value();
  const int childrenCount = tabs->children();
  for (int i = 0; i < childrenCount; ++i) {
    if (tabs->child(i) == widget) {
      return i;
    }
  }
  return -1;
}

void go_fltk_Tabs_set_value(GTabs* tabs, int value) {
  if (value < 0 || value >= tabs->children()) {
    return;
  }
  tabs->value(tabs->child(value));
}
