#include "roller.h"

#include <FL/Fl_Roller.H>

Fl_Roller* go_fltk_new_Roller(int x, int y, int w, int h, const char* label) {
  return new Fl_Roller(x, y, w, h, label);
}
