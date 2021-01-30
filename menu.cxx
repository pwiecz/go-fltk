#include "menu.h"

#include <cstdint>

#include <Fl/Fl_Menu_.H>
#include <Fl/Fl_Menu_Button.H>

#include "enumerations.h"


int go_fltk_Menu_add(Fl_Menu_ *m, char *label, int shortcut, int callback, int flags) {
  return m->add(label, shortcut, callback_handler, (void*)(uintptr_t)callback, flags);
}

void go_fltk_Menu_set_value(Fl_Menu_ *m, int value) { 
  m->value(value);
}

int go_fltk_Menu_value(Fl_Menu_ *m) {
  return m->value();
}

Fl_Menu_Button* go_fltk_new_MenuButton(int x, int y, int w, int h, const char* text) {
  return new Fl_Menu_Button(x, y, w, h, text);
}
void go_fltk_MenuButton_set_type(Fl_Menu_Button* m, int menuType) {
  m->type(menuType);
}

const int go_FL_POPUP1 = Fl_Menu_Button::POPUP1;
const int go_FL_POPUP2 = Fl_Menu_Button::POPUP2;
const int go_FL_POPUP12 = Fl_Menu_Button::POPUP12;
const int go_FL_POPUP3 = Fl_Menu_Button::POPUP3;
const int go_FL_POPUP13 = Fl_Menu_Button::POPUP13;
const int go_FL_POPUP23 = Fl_Menu_Button::POPUP23;
const int go_FL_POPUP123 = Fl_Menu_Button::POPUP123;
