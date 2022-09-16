#include "group.h"

#include <FL/Fl_Group.H>

#include "event_handler.h"


class GGroup : public EventHandler<Fl_Group> {
public:
  GGroup(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Group>(x, y, w, h, label) {}
  void draw_children() {
    Fl_Group::draw_children();
  }
};

GGroup *go_fltk_new_Group(int x, int y, int w, int h, const char *label) {
  return new GGroup(x, y, w, h, label);
}

void go_fltk_Group_begin(GGroup *g) {
  g->begin();
}
void go_fltk_Group_end(GGroup *g) {
  g->end();
}
void go_fltk_Group_add(GGroup *g, Fl_Widget *w) {
  g->add(w);
}
void go_fltk_Group_remove(GGroup *g, Fl_Widget *w) {
  g->remove(w);
}
void go_fltk_Group_resizable(GGroup *g, Fl_Widget *w) {
  g->resizable(w);
}
void go_fltk_Group_draw_children(GGroup *g) {
  g->draw_children();
}
