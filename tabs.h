#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Tabs Fl_Tabs;

  extern Fl_Tabs *go_fltk_new_Tabs(int x, int y, int w, int h, const char *label);

#ifdef __cplusplus
}
#endif
