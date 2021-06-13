#pragma once

#include <stdint.h>


#ifdef __cplusplus
extern "C" {
#endif

  extern int go_fltk_run();
  extern int go_fltk_lock();

  extern int go_fltk_awake(uintptr_t id);

  extern void go_fltk_copy(const char* data, int len, int destination);

  extern void go_fltk_screen_work_area(int *x, int *y, int *w, int *h, int n);
  extern void go_fltk_screen_dpi(float *w, float *h, int n);
  extern int go_fltk_screen_count();
  extern float go_fltk_screen_scale(int screenNum);
  extern void go_fltk_set_screen_scale(int screenNum, float scale);

#ifdef __cplusplus
}
#endif
