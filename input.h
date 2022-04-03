#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  extern const unsigned char go_FL_INT_INPUT;
  extern const unsigned char go_FL_FLOAT_INPUT;
  
  typedef struct GInput GInput;

  extern GInput *go_fltk_new_Input(int x, int y, int w, int h, const char *text);

  const char *go_fltk_Input_value(GInput *in);
  int go_fltk_Input_set_value(GInput *in, const char *t);
  void go_fltk_Input_resize(GInput *in, int x, int y, int w, int h);

#ifdef __cplusplus
}
#endif
