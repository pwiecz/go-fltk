#include "pack.h"

#include <Fl/Fl_Pack.H>


Fl_Pack *go_fltk_new_Pack(int x, int y, int w, int h, const char *label) {
  return new Fl_Pack(x, y, w, h, label);
}

void go_fltk_Pack_set_spacing(Fl_Pack* pack, int spacing) {
  pack->spacing(spacing);
}

const unsigned char go_FL_PACK_VERTICAL = Fl_Pack::VERTICAL;
const unsigned char go_FL_PACK_HORIZONTAL = Fl_Pack::HORIZONTAL;
