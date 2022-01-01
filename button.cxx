#include "button.h"

#include <FL/Fl_Button.H>
#include <FL/Fl_Check_Button.H>
#include <FL/Fl_Toggle_Button.H>
#include <FL/Fl_Radio_Button.H>
#include <FL/Fl_Return_Button.H>
#include <FL/Fl_Radio_Round_Button.H>

Fl_Button *go_fltk_new_Button(int x, int y, int w, int h, const char *label) {
  return new Fl_Button(x, y, w, h, label);
}

char go_fltk_Button_value(Fl_Button *b) {
  return b->value();
}
void go_fltk_Button_set_value(Fl_Button *b, int val) {
  b->value(val);
}
void go_fltk_Button_set_down_box(Fl_Button *b, int val) {
  b->down_box((Fl_Boxtype)val);
}

Fl_Check_Button *go_fltk_new_Check_Button(int x, int y, int w, int h, const char *label) {
  return new Fl_Check_Button(x, y, w, h, label);
}

Fl_Toggle_Button *go_fltk_new_Toggle_Button(int x, int y, int w, int h, const char *label) {
  return new Fl_Toggle_Button(x, y, w, h, label);
}

Fl_Radio_Button *go_fltk_new_Radio_Button(int x, int y, int w, int h, const char *label) {
  return new Fl_Radio_Button(x, y, w, h, label);
}

Fl_Return_Button *go_fltk_new_Return_Button(int x, int y, int w, int h, const char *label) {
  return new Fl_Return_Button(x, y, w, h, label);
}

Fl_Radio_Round_Button *go_fltk_new_Radio_Round_Button(int x, int y, int w, int h, const char *label) {
  return new Fl_Radio_Round_Button(x, y, w, h, label);
}
