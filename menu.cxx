#include "menu.h"

#include <cstdint>

#include <Fl/Fl_Menu_.H>

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

