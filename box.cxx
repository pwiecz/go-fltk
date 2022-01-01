#include "box.h"

#include <FL/Fl_Box.H>


Fl_Box *go_fltk_new_Box(int boxType, int x, int y, int w, int h, const char *label) {
  return new Fl_Box((Fl_Boxtype)boxType, x, y, w, h, label);
}
