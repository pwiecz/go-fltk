#include "choice.h"

#include <FL/Fl_Choice.H>


Fl_Choice *go_fltk_new_Choice(int x, int y, int w, int h, const char *label) {
  return new Fl_Choice(x, y, w, h, label);
}

