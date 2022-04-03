#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GChoice GChoice;

  extern GChoice *go_fltk_new_Choice(int x, int y, int w, int h, const char *label);

#ifdef __cplusplus
}
#endif
