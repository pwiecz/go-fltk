#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Roller Fl_Roller;

  extern Fl_Roller* go_fltk_new_Roller(int x, int y, int w, int h, const char* label);

#ifdef __cplusplus
}
#endif
