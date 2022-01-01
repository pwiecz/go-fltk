#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Button Fl_Button;
  typedef struct Fl_Check_Button Fl_Check_Button;
  typedef struct Fl_Radio_Button Fl_Radio_Button;
  typedef struct Fl_Toggle_Button Fl_Toggle_Button;
  typedef struct Fl_Return_Button Fl_Return_Button;
  typedef struct Fl_Radio_Round_Button Fl_Radio_Round_Button;

  extern Fl_Button* go_fltk_new_Button(int x, int y, int w, int h, const char* label);
  extern char go_fltk_Button_value(Fl_Button* b);
  extern void go_fltk_Button_set_value(Fl_Button* b, int value);
  extern void go_fltk_Button_set_down_box(Fl_Button* b, int value);
  extern Fl_Check_Button *go_fltk_new_Check_Button(int x, int y, int w, int h, const char *label);
  extern Fl_Radio_Button *go_fltk_new_Radio_Button(int x, int y, int w, int h, const char *label);
  extern Fl_Toggle_Button *go_fltk_new_Toggle_Button(int x, int y, int w, int h, const char *label);
  extern Fl_Return_Button *go_fltk_new_Return_Button(int x, int y, int w, int h, const char *label);
  extern Fl_Radio_Round_Button *go_fltk_new_Radio_Round_Button(int x, int y, int w, int h, const char *label);

#ifdef __cplusplus
}
#endif
