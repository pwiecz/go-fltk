#include "check_button.h"

#include "FL/Fl_Check_Button.H"


Fl_Check_Button *go_fltk_new_Check_Button(int x, int y, int w, int h, const char *label) {
  return new Fl_Check_Button(x, y, w, h, label);
}
