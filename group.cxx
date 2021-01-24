#include "group.h"

#include "Fl/Fl_Group.H"


Fl_Group *go_fltk_new_Group(int x, int y, int w, int h, const char *label) {
  return new Fl_Group(x, y, w, h, label);
}

void go_fltk_Group_begin(Fl_Group *g) {
  g->begin();
}
void go_fltk_Group_end(Fl_Group *g) {
  g->end();
}
void go_fltk_Group_add(Fl_Group *g, Fl_Widget *w) {
  g->add(w);
}
void go_fltk_Group_resizable(Fl_Group *g, Fl_Widget *w) {
  g->resizable(w);
}
