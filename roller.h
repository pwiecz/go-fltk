#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GRoller GRoller;

  extern GRoller* go_fltk_new_Roller(int x, int y, int w, int h, const char* label);

#ifdef __cplusplus
}
#endif
