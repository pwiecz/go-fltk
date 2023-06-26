#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Tile Fl_Tile;
  typedef struct GTile GTile;

  extern GTile *go_fltk_new_Tile(int x, int y, int w, int h, const char *text);

#ifdef __cplusplus
}
#endif
