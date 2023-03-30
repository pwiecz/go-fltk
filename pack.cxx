#include "pack.h"

#include <FL/Fl_Pack.H>

#include "event_handler.h"


class GPack : public EventHandler<Fl_Pack> {
public:
  GPack(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Pack>(x, y, w, h, label) {}
};

GPack *go_fltk_new_Pack(int x, int y, int w, int h, const char *label) {
  return new GPack(x, y, w, h, label);
}

int go_fltk_Pack_spacing(Fl_Pack* pack) {
  return pack->spacing();
}

void go_fltk_Pack_set_spacing(Fl_Pack* pack, int spacing) {
  pack->spacing(spacing);
}

const unsigned char go_FL_PACK_VERTICAL = Fl_Pack::VERTICAL;
const unsigned char go_FL_PACK_HORIZONTAL = Fl_Pack::HORIZONTAL;
