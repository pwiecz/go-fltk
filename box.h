#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GBox GBox;

  extern GBox *go_fltk_new_Box(int boxType, int x, int y, int w, int h, const char *label);
  extern void go_fltk_Box_set_event_handler(GBox* w, int handlerId);

#ifdef __cplusplus
}
#endif
