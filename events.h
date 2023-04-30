#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  extern const int go_FL_LEFT_MOUSE;
  extern const int go_FL_MIDDLE_MOUSE;
  extern const int go_FL_RIGHT_MOUSE;

  extern const int go_FL_SHIFT;
  extern const int go_FL_CAPS_LOCK;
  extern const int go_FL_CTRL;
  extern const int go_FL_ALT;
  extern const int go_FL_NUM_LOCK;
  extern const int go_FL_META;
  extern const int go_FL_SCROLL_LOCK;
  extern const int go_FL_BUTTON1;
  extern const int go_FL_BUTTON2;
  extern const int go_FL_BUTTON3;

  extern int go_fltk_event();
  extern int go_fltk_event_button();
  extern int go_fltk_event_button1();
  extern int go_fltk_event_x();
  extern int go_fltk_event_y();
  extern int go_fltk_event_x_root();
  extern int go_fltk_event_y_root();
  extern int go_fltk_event_dx();
  extern int go_fltk_event_dy();
  extern int go_fltk_event_key();
  extern int go_fltk_event_is_click();
  extern int go_fltk_event_clicks();
  extern void go_fltk_event_set_clicks(int i);
  extern int go_fltk_event_state();
  extern const char* go_fltk_event_text();

#ifdef __cplusplus
}
#endif
