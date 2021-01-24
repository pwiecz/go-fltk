#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  extern const unsigned char go_FL_INT_INPUT;
  extern const unsigned char go_FL_FLOAT_INPUT;
  
  typedef struct Fl_Input Fl_Input;

  extern Fl_Input *go_fltk_new_Input(int x, int y, int w, int h, const char *text);

  const char *go_fltk_Input_value(Fl_Input *in);
  int go_fltk_Input_set_value(Fl_Input *in, const char *t);
  void go_fltk_Input_resize(Fl_Input *in, int x, int y, int w, int h);

#ifdef __cplusplus
}
#endif
