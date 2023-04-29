#include "tooltip.h"

#include <FL/Fl_Tooltip.H>


void go_fltk_enable_tooltips() {
  Fl_Tooltip::enable(1);
}

void go_fltk_disable_tooltips() {
  Fl_Tooltip::disable();
}

int go_fltk_tooltips_enabled() {
  return Fl_Tooltip::enabled();
}

float go_fltk_tooltip_delay() {
  return Fl_Tooltip::delay();  
}

void go_fltk_set_tooltip_delay(float delay) {
  Fl_Tooltip::delay(delay);  
}
