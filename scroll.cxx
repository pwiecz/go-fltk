#include "scroll.h"

#include <FL/Fl_Scroll.H>

#include "event_handler.h"


class GScroll : public EventHandler<Fl_Scroll> {
public:
  GScroll(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Scroll>(x, y, w, h, label) {}
};

GScroll *go_fltk_new_Scroll(int x, int y, int w, int h, const char *label) {
  return new GScroll(x, y, w, h, label);
}

void go_fltk_Scroll_scroll_to(Fl_Scroll* scroll, int x, int y) {
  scroll->scroll_to(x, y);
}
int go_fltk_Scroll_x_position(Fl_Scroll *scroll) {
  return scroll->xposition();  
}
int go_fltk_Scroll_y_position(Fl_Scroll *scroll) {
  return scroll->yposition();
}  

const unsigned char go_FL_SCROLL_HORIZONTAL = Fl_Scroll::HORIZONTAL;
const unsigned char go_FL_SCROLL_VERTICAL = Fl_Scroll::VERTICAL;
const unsigned char go_FL_SCROLL_BOTH = Fl_Scroll::BOTH;
const unsigned char go_FL_SCROLL_HORIZONTAL_ALWAYS = Fl_Scroll::HORIZONTAL_ALWAYS;
const unsigned char go_FL_SCROLL_VERTICAL_ALWAYS = Fl_Scroll::VERTICAL_ALWAYS;
const unsigned char go_FL_SCROLL_BOTH_ALWAYS = Fl_Scroll::BOTH_ALWAYS;
