#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GText_Editor GText_Editor;
  typedef struct GText_Display GText_Display;
  typedef struct Fl_Text_Buffer Fl_Text_Buffer;

  extern GText_Editor *go_fltk_new_TextEditor(int x, int y, int w, int h, const char *text);
  extern void go_fltk_TextEditor_copy(GText_Editor *e);
  extern void go_fltk_TextEditor_insert(GText_Editor *e);
  extern void go_fltk_TextEditor_cut(GText_Editor *e);
  extern void go_fltk_TextEditor_delete(GText_Editor *e);
  extern void go_fltk_TextEditor_paste(GText_Editor *e);
  extern void go_fltk_TextEditor_select_all(GText_Editor *e);

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
  void go_fltk_TextDisplay_insert_position(GText_Display *d, int newPos);
  int go_fltk_TextDisplay_get_insert_position(GText_Display *d);
  void go_fltk_TextDisplay_insert_text(GText_Display *d, const char *text);
  void go_fltk_TextDisplay_overstrike(GText_Display *d, const char* text);
  extern int go_fltk_TextDisplay_total_lines(GText_Display *d);
  extern int go_fltk_TextDisplay_line_to_position(GText_Display *d, int lineNum);
  extern int go_fltk_TextDisplay_move_end(GText_Display *d);

  extern Fl_Text_Buffer *go_fltk_new_TextBuffer(void);
  extern void go_fltk_TextBuffer_delete(Fl_Text_Buffer* b);
  extern void go_fltk_TextBuffer_set_text(Fl_Text_Buffer *b, const char *txt);
  extern void go_fltk_TextBuffer_append(Fl_Text_Buffer *b, const char *txt);
  extern const char *go_fltk_TextBuffer_text(Fl_Text_Buffer *b);
  const char *go_fltk_TextBuffer_text_range(Fl_Text_Buffer *b, int start, int end);
  void go_fltk_TextBuffer_highlight(Fl_Text_Buffer *b, int start, int end);
  void go_fltk_TextBuffer_unhighlight(Fl_Text_Buffer *b);
  void go_fltk_TextBuffer_replace(Fl_Text_Buffer *b, int start, int end, const char *text);
  void go_fltk_TextBuffer_replace_selection(Fl_Text_Buffer *b, const char *text);
  int go_fltk_TextBuffer_search_forward(Fl_Text_Buffer *b, int start, const char *searchString, int *foundPos, int matchCase);
  int go_fltk_TextBuffer_search_backward(Fl_Text_Buffer *b, int start, const char *searchString, int *foundPos, int matchCase);
  void go_fltk_TextBuffer_select(Fl_Text_Buffer *b, int start, int end);
  int go_fltk_TextBuffer_selected(Fl_Text_Buffer *b);
  int go_fltk_TextBuffer_selection_position(Fl_Text_Buffer *b, int *start, int *end);
  char* go_fltk_TextBuffer_selection_text(Fl_Text_Buffer *b);
  void go_fltk_TextBuffer_unselect(Fl_Text_Buffer *b);
  int go_fltk_TextBuffer_tab_distance(Fl_Text_Buffer *b);
  void go_fltk_TextBuffer_set_tab_distance(Fl_Text_Buffer *b, int tabDist);

#ifdef __cplusplus
}
#endif
