#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GText_Editor GText_Editor;
  typedef struct GText_Display GText_Display;
  typedef struct Fl_Text_Buffer Fl_Text_Buffer;

  extern GText_Editor *go_fltk_new_TextEditor(int x, int y, int w, int h, const char *text);
  extern GText_Display *go_fltk_new_TextDisplay(int x, int y, int w, int h, const char *text);
  extern void go_fltk_TextDisplay_set_buffer(GText_Display *d, Fl_Text_Buffer *buf);
  extern void go_fltk_TextDisplay_set_wrap_mode(GText_Display *b, int wrap, int wrapMargin);
  extern Fl_Text_Buffer *go_fltk_TextDisplay_buffer(GText_Display *d);
  extern Fl_Text_Buffer *go_fltk_new_TextBuffer(void);
  extern void go_fltk_TextBuffer_set_text(Fl_Text_Buffer *b, const char *txt);
  extern void go_fltk_TextBuffer_append(Fl_Text_Buffer *b, const char *txt);
  extern const char *go_fltk_TextBuffer_text(Fl_Text_Buffer *b);

#ifdef __cplusplus
}
#endif
