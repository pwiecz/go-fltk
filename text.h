#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Text_Editor Fl_Text_Editor;
  typedef struct Fl_Text_Display Fl_Text_Display;
  typedef struct Fl_Text_Buffer Fl_Text_Buffer;

  extern Fl_Text_Editor *go_fltk_new_TextEditor(int x, int y, int w, int h, const char *text);
  extern Fl_Text_Display *go_fltk_new_TextDisplay(int x, int y, int w, int h, const char *text);
  extern void go_fltk_TextDisplay_set_buffer(Fl_Text_Display *d, Fl_Text_Buffer *buf);
  extern Fl_Text_Buffer *go_fltk_TextDisplay_buffer(Fl_Text_Display *d);
  extern Fl_Text_Buffer *go_fltk_new_TextBuffer(void);
  extern void go_fltk_TextBuffer_set_text(Fl_Text_Buffer *b, const char *txt);
  extern void go_fltk_TextBuffer_append(Fl_Text_Buffer *b, const char *txt);
  extern const char *go_fltk_TextBuffer_text(Fl_Text_Buffer *b);

#ifdef __cplusplus
}
#endif
