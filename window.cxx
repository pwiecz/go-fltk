#include "window.h"

#include "Fl/Fl_Window.H"


Fl_Window *go_fltk_new_Window(int w, int h) {
  return new Fl_Window(w, h);
}

void go_fltk_Window_show(Fl_Window *w) {
  w->show();
}

void go_fltk_Window_set_label(Fl_Window *w, const char *label) {
  w->label(label);
}

