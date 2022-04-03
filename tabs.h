#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GTabs GTabs;

  extern GTabs *go_fltk_new_Tabs(int x, int y, int w, int h, const char *label);


  extern int go_fltk_Tabs_value(GTabs* tabs);

  extern void go_fltk_Tabs_set_value(GTabs* tabs, int value);

#ifdef __cplusplus
}
#endif
