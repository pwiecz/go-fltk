#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Text_Editor Fl_Text_Editor;

  extern Fl_Text_Editor *go_fltk_new_TextEditor(int x, int y, int w, int h, const char *text);

#ifdef __cplusplus
}
#endif
