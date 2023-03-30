#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Pack Fl_Pack;
  typedef struct GPack GPack;

  extern GPack *go_fltk_new_Pack(int x, int y, int w, int h, const char *text);

  extern int go_fltk_Pack_spacing(Fl_Pack* pack);

  extern void go_fltk_Pack_set_spacing(Fl_Pack* pack, int spacing);

  extern const unsigned char go_FL_PACK_VERTICAL;
  extern const unsigned char go_FL_PACK_HORIZONTAL;

#ifdef __cplusplus
}
#endif
