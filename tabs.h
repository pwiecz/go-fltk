#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Tabs Fl_Tabs;
  typedef struct GTabs GTabs;

  extern GTabs *go_fltk_new_Tabs(int x, int y, int w, int h, const char *label);


  extern int go_fltk_Tabs_value(Fl_Tabs* tabs);

  extern void go_fltk_Tabs_set_value(Fl_Tabs* tabs, int value);

  extern const int go_FL_OVERFLOW_COMPRESS;
  extern const int go_FL_OVERFLOW_CLIP;
  extern const int go_FL_OVERFLOW_PULLDOWN;
  extern const int go_FL_OVERFLOW_DRAG;

  extern void go_fltk_Tabs_handle_overflow(Fl_Tabs* tabs, int overflow);

#ifdef __cplusplus
}
#endif
