#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GText_Editor GText_Editor;
  typedef struct GText_Display GText_Display;
  typedef struct Fl_Text_Buffer Fl_Text_Buffer;

  extern GText_Editor *go_fltk_new_TextEditor(int x, int y, int w, int h, const char *text);
  extern void go_fltk_TextEditor_copy(GText_Editor *e);

  extern GText_Display *go_fltk_new_TextDisplay(int x, int y, int w, int h, const char *text);
  extern void go_fltk_TextDisplay_set_buffer(GText_Display *d, Fl_Text_Buffer *buf);
  extern void go_fltk_TextDisplay_set_wrap_mode(GText_Display *b, int wrap, int wrapMargin);
  extern int go_fltk_TextDisplay_move_right(GText_Display *b);
  extern int go_fltk_TextDisplay_move_left(GText_Display *b);
  extern int go_fltk_TextDisplay_move_up(GText_Display *b);
  extern int go_fltk_TextDisplay_move_down(GText_Display *b);
  extern void go_fltk_TextDisplay_show_insert_position(GText_Display *b);
  extern int go_fltk_TextDisplay_text_size(GText_Display *d);
  extern void go_fltk_TextDisplay_set_text_size(GText_Display *d, int size);
  extern Fl_Text_Buffer *go_fltk_TextDisplay_buffer(GText_Display *d);

  extern Fl_Text_Buffer *go_fltk_new_TextBuffer(void);
  extern void go_fltk_TextBuffer_delete(Fl_Text_Buffer* b);
  extern void go_fltk_TextBuffer_set_text(Fl_Text_Buffer *b, const char *txt);
  extern void go_fltk_TextBuffer_append(Fl_Text_Buffer *b, const char *txt);
  extern const char *go_fltk_TextBuffer_text(Fl_Text_Buffer *b);

#ifdef __cplusplus
}
#endif
