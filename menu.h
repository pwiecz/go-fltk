#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Menu_ Fl_Menu_;

  extern int go_fltk_Menu_add(Fl_Menu_* m, char* label, int shortcut, int callback, int flags);
  extern void go_fltk_Menu_set_value(Fl_Menu_* m, int value);
  extern int go_fltk_Menu_value(Fl_Menu_* m);

#ifdef __cplusplus
}
#endif
