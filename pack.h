#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GPack GPack;

  extern GPack *go_fltk_new_Pack(int x, int y, int w, int h, const char *text);

  extern void go_fltk_Pack_set_spacing(GPack* pack, int spacing);

  extern const unsigned char go_FL_PACK_VERTICAL;
  extern const unsigned char go_FL_PACK_HORIZONTAL;

#ifdef __cplusplus
}
#endif
