#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  extern const unsigned char go_FL_INT_INPUT;
  extern const unsigned char go_FL_FLOAT_INPUT;

  typedef struct Fl_Input Fl_Input;
  typedef struct GInput GInput;
  typedef struct GOutput GOutput;
  typedef struct GFloat_Input GFloat_Input;
  typedef struct GInt_Input GInt_Input;

  extern GInput *go_fltk_new_Input(int x, int y, int w, int h, const char *text);

  const char *go_fltk_Input_value(Fl_Input *in);
  int go_fltk_Input_set_value(Fl_Input *in, const char *t);
  void go_fltk_Input_resize(Fl_Input *in, int x, int y, int w, int h);
  void go_fltk_Input_set_insert_position(Fl_Input *in, int p, int m);
  int go_fltk_Input_insert_position(Fl_Input *in);
  int go_fltk_Input_mark(Fl_Input *in);
		
  extern GOutput *go_fltk_new_Output(int x, int y, int w, int h, const char *text);
  extern GFloat_Input *go_fltk_new_Float_Input(int x, int y, int w, int h, const char *text);
  extern GInt_Input *go_fltk_new_Int_Input(int x, int y, int w, int h, const char *text);

#ifdef __cplusplus
}
#endif
