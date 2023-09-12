#pragma once

#include <stdint.h>


#ifdef __cplusplus
extern "C" {
#endif
  typedef struct Fl_Text_Editor Fl_Text_Editor;
  typedef struct GText_Editor GText_Editor;
  typedef struct Fl_Text_Display Fl_Text_Display;
  typedef struct GText_Display GText_Display;
  typedef struct Fl_Text_Buffer Fl_Text_Buffer;

  extern GText_Editor *go_fltk_new_TextEditor(int x, int y, int w, int h, const char *text);
  extern void go_fltk_TextEditor_copy(Fl_Text_Editor *e);
  extern void go_fltk_TextEditor_insert(Fl_Text_Editor *e);
  extern void go_fltk_TextEditor_cut(Fl_Text_Editor *e);
  extern void go_fltk_TextEditor_delete(Fl_Text_Editor *e);
  extern void go_fltk_TextEditor_paste(Fl_Text_Editor *e);
  extern void go_fltk_TextEditor_select_all(Fl_Text_Editor *e);
  extern int go_fltk_TextEditor_undo(Fl_Text_Editor *e);
  extern int go_fltk_TextEditor_redo(Fl_Text_Editor *e);

  extern GText_Display *go_fltk_new_TextDisplay(int x, int y, int w, int h, const char *text);
  extern void go_fltk_TextDisplay_set_buffer(Fl_Text_Display *d, Fl_Text_Buffer *buf);
  extern void go_fltk_TextDisplay_set_wrap_mode(Fl_Text_Display *d, int wrap, int wrapMargin);
  extern int go_fltk_TextDisplay_xy_to_position(Fl_Text_Display *d, int x, int y);
  extern int go_fltk_TextDisplay_position_to_xy(Fl_Text_Display *d, int pos, int *x, int *y);
  extern int go_fltk_TextDisplay_move_right(Fl_Text_Display *d);
  extern int go_fltk_TextDisplay_move_left(Fl_Text_Display *d);
  extern int go_fltk_TextDisplay_move_up(Fl_Text_Display *d);
  extern int go_fltk_TextDisplay_move_down(Fl_Text_Display *d);
  extern void go_fltk_TextDisplay_show_insert_position(Fl_Text_Display *d);
  extern void go_fltk_TextDisplay_hide_cursor(Fl_Text_Display *d);
  extern unsigned int go_fltk_TextDisplay_text_color(Fl_Text_Display *d);
  extern void go_fltk_TextDisplay_set_text_color(Fl_Text_Display* d, unsigned int color);
  extern int go_fltk_TextDisplay_text_font(Fl_Text_Display *d);
  extern void go_fltk_TextDisplay_set_text_font(Fl_Text_Display* d, int font);
  extern int go_fltk_TextDisplay_text_size(Fl_Text_Display *d);
  extern void go_fltk_TextDisplay_set_text_size(Fl_Text_Display *d, int size);
  extern Fl_Text_Buffer *go_fltk_TextDisplay_buffer(Fl_Text_Display *d);
  extern void go_fltk_TextDisplay_insert_position(Fl_Text_Display *d, int newPos);
  extern int go_fltk_TextDisplay_get_insert_position(Fl_Text_Display *d);
  extern void go_fltk_TextDisplay_insert_text(Fl_Text_Display *d, const char *text);
  extern void go_fltk_TextDisplay_overstrike(Fl_Text_Display *d, const char* text);
  extern void go_fltk_TextDisplay_set_linenumber_width(Fl_Text_Display *d, int width);
  extern void go_fltk_TextDisplay_set_linenumber_font(Fl_Text_Display *d, int val);
  extern void go_fltk_TextDisplay_set_linenumber_size(Fl_Text_Display *d, int val);
  extern void go_fltk_TextDisplay_set_linenumber_fgcolor(Fl_Text_Display *d, unsigned int val);
  extern void go_fltk_TextDisplay_set_linenumber_bgcolor(Fl_Text_Display *d, unsigned int val);
  extern void go_fltk_TextDisplay_set_linenumber_align(Fl_Text_Display *d, int val);

  extern Fl_Text_Buffer *go_fltk_new_TextBuffer(void);
  extern void go_fltk_TextBuffer_add_modify_callback(Fl_Text_Buffer *b, uintptr_t handlerId);
  extern void go_fltk_TextBuffer_delete(Fl_Text_Buffer* b);
  extern void go_fltk_TextBuffer_set_text(Fl_Text_Buffer *b, const char *txt);
  extern void go_fltk_TextBuffer_append(Fl_Text_Buffer *b, const char *txt);
  extern void go_fltk_TextBuffer_insert(Fl_Text_Buffer *b, int pos, const char *txt);
  extern void go_fltk_TextBuffer_remove(Fl_Text_Buffer *b, int start, int end);  
  extern unsigned int go_fltk_TextBuffer_char_at(Fl_Text_Buffer *b, int pos);
  extern int go_fltk_TextBuffer_next_char(Fl_Text_Buffer *b, int ix);
  extern int go_fltk_TextBuffer_prev_char(Fl_Text_Buffer *b, int ix);
  extern int go_fltk_TextBuffer_line_start(Fl_Text_Buffer *b, int ix);
  extern int go_fltk_TextBuffer_line_end(Fl_Text_Buffer *b, int ix);
  extern const char *go_fltk_TextBuffer_line_text(Fl_Text_Buffer *b, int ix);
  extern int go_fltk_TextBuffer_count_lines(Fl_Text_Buffer *b, int start, int end);
  extern int go_fltk_TextBuffer_skip_lines(Fl_Text_Buffer *b, int start, int nlines);
  extern int go_fltk_TextBuffer_rewind_lines(Fl_Text_Buffer *b, int start, int nlines);  
  extern int go_fltk_TextBuffer_length(Fl_Text_Buffer *b);
  extern const char *go_fltk_TextBuffer_text(Fl_Text_Buffer *b);
  extern const char *go_fltk_TextBuffer_text_range(Fl_Text_Buffer *b, int start, int end);
  extern void go_fltk_TextBuffer_highlight(Fl_Text_Buffer *b, int start, int end);
  extern void go_fltk_TextBuffer_unhighlight(Fl_Text_Buffer *b);
  extern void go_fltk_TextBuffer_replace(Fl_Text_Buffer *b, int start, int end, const char *text);
  extern void go_fltk_TextBuffer_replace_selection(Fl_Text_Buffer *b, const char *text);
  extern int go_fltk_TextBuffer_search_forward(Fl_Text_Buffer *b, int start, const char *searchString, int *foundPos, int matchCase);
  extern int go_fltk_TextBuffer_search_backward(Fl_Text_Buffer *b, int start, const char *searchString, int *foundPos, int matchCase);
  extern void go_fltk_TextBuffer_select(Fl_Text_Buffer *b, int start, int end);
  extern int go_fltk_TextBuffer_selected(Fl_Text_Buffer *b);
  extern int go_fltk_TextBuffer_selection_position(Fl_Text_Buffer *b, int *start, int *end);
  extern char* go_fltk_TextBuffer_selection_text(Fl_Text_Buffer *b);
  extern void go_fltk_TextBuffer_unselect(Fl_Text_Buffer *b);
  extern int go_fltk_TextBuffer_tab_distance(Fl_Text_Buffer *b);
  extern void go_fltk_TextBuffer_set_tab_distance(Fl_Text_Buffer *b, int tabDist);
  extern void go_fltk_TextDisplay_set_highlight_data(
      Fl_Text_Display *d, Fl_Text_Buffer *sbuff, unsigned int *color,
      int *font, int *fontsz, unsigned *attr, unsigned int *bgcolor, int sz);
  
#ifdef __cplusplus
}
#endif
