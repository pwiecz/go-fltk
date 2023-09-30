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

void go_fltk_Group_begin(Fl_Group *g) {
  g->begin();
}
void go_fltk_Group_end(Fl_Group *g) {
  g->end();
}
void go_fltk_Group_add(Fl_Group *g, Fl_Widget *w) {
  g->add(w);
}
void go_fltk_Group_remove(Fl_Group *g, Fl_Widget *w) {
  g->remove(w);
}
void go_fltk_Group_resizable(Fl_Group *g, Fl_Widget *w) {
  g->resizable(w);
}
void go_fltk_Group_draw_children(Fl_Group *g) {
  GGroup *const gg = dynamic_cast<GGroup *>(g);
  if (gg != nullptr) {
    gg->draw_children();
  }    
}
Fl_Widget* go_fltk_Group_child(Fl_Group *g, int index) {
  if (index < 0 || index >= g->children()) {
    return nullptr;
  }
  return g->child(index);
}

int go_fltk_Group_child_count(Fl_Group *g) {
  return g->children();
}
