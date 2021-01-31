#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  extern int go_fltk_run();
  extern int go_fltk_lock();

  extern int go_fltk_awake(void* id);

  extern void go_fltk_copy(const char* data, int len, int destination);

#ifdef __cplusplus
}
#endif
