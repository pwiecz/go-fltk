#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Button Fl_Button;

  extern Fl_Button* go_fltk_new_Button(int x, int y, int w, int h, const char* label);
  extern char go_fltk_Button_value(Fl_Button* b);

#ifdef __cplusplus
}
#endif
