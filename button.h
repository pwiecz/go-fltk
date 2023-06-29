#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GButton GButton;
  typedef struct Fl_Button Fl_Button;
  typedef struct GCheck_Button GCheck_Button;
  typedef struct GRadio_Button GRadio_Button;
  typedef struct GToggle_Button GToggle_Button;
  typedef struct GLight_Button GLight_Button;
  typedef struct GReturn_Button GReturn_Button;
  typedef struct GRadio_Round_Button GRadio_Round_Button;

  extern GButton* go_fltk_new_Button(int x, int y, int w, int h, const char* label);
  extern char go_fltk_Button_value(Fl_Button* b);
  extern void go_fltk_Button_set_value(Fl_Button* b, int value);
  extern void go_fltk_Button_set_down_box(Fl_Button* b, int value);
  extern void go_fltk_Button_set_shortcut(Fl_Button *b, int shortcut);
  extern int go_fltk_Button_shortcut(Fl_Button *b);
		
  extern GCheck_Button *go_fltk_new_Check_Button(int x, int y, int w, int h, const char *label);
  extern GRadio_Button *go_fltk_new_Radio_Button(int x, int y, int w, int h, const char *label);
  extern GToggle_Button *go_fltk_new_Toggle_Button(int x, int y, int w, int h, const char *label);
  extern GLight_Button *go_fltk_new_Light_Button(int x, int y, int w, int h, const char *label);
  extern GReturn_Button *go_fltk_new_Return_Button(int x, int y, int w, int h, const char *label);
  extern GRadio_Round_Button *go_fltk_new_Radio_Round_Button(int x, int y, int w, int h, const char *label);

#ifdef __cplusplus
}
#endif
