#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  extern int go_fltk_run();
  extern int go_fltk_lock();

  extern int go_fltk_awake(void* id);

  extern int go_fltk_event_button1();
  extern int go_fltk_event_x();
  extern int go_fltk_event_y();
  extern int go_fltk_event_dx();
  extern int go_fltk_event_dy();
  extern int go_fltk_event_key();
  extern int go_fltk_event_state();

  extern void go_fltk_copy(const char* data, int len, int destination);

#ifdef __cplusplus
}
#endif
