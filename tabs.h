#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Tabs Fl_Tabs;

  extern Fl_Tabs *go_fltk_new_Tabs(int x, int y, int w, int h, const char *label);


  extern int go_fltk_Tabs_value(Fl_Tabs* tabs);

  extern void go_fltk_Tabs_set_value(Fl_Tabs* tabs, int value);

#ifdef __cplusplus
}
#endif
