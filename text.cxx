#include "text.h"

#include <FL/Fl_Text_Display.H>

#include <FL/Fl_Text_Editor.H>

#include <FL/Fl_Text_Buffer.H>

#include "event_handler.h"


class GText_Display : public EventHandler<Fl_Text_Display> {
public:
  GText_Display(int x, int y, int w, int h, const char* label)
    : EventHandler<Fl_Text_Display>(x, y, w, h, label) {}
};

GText_Display *go_fltk_new_TextDisplay(int x, int y, int w, int h, const char *text) {
  return new GText_Display(x, y, w, h, text);
}

void go_fltk_TextDisplay_set_buffer(Fl_Text_Display *d, Fl_Text_Buffer *buf) {
  d->buffer(buf);
}

void go_fltk_TextDisplay_set_wrap_mode(Fl_Text_Display *d, int wrap, int wrapMargin) {
  d->wrap_mode(wrap, wrapMargin);
}

int go_fltk_TextDisplay_move_right(Fl_Text_Display *d) {
  return d->move_right();
}

int go_fltk_TextDisplay_move_left(Fl_Text_Display *d) {
  return d->move_left();
}

int go_fltk_TextDisplay_move_up(Fl_Text_Display *d) {
  return d->move_up();
}

int go_fltk_TextDisplay_move_down(Fl_Text_Display *d) {
  return d->move_down();
}

void go_fltk_TextDisplay_show_insert_position(Fl_Text_Display *d) {
  d->show_insert_position();
}

unsigned int go_fltk_TextDisplay_text_color(Fl_Text_Display *d) {
  return d->textcolor();
}

void go_fltk_TextDisplay_set_text_color(Fl_Text_Display *d, unsigned int color) {
  d->textcolor(color);
}
  
int go_fltk_TextDisplay_text_size(Fl_Text_Display *d) {
  return d->textsize();
}

void go_fltk_TextDisplay_set_text_size(Fl_Text_Display *d, int size) {
  d->textsize(size);
}

Fl_Text_Buffer *go_fltk_TextDisplay_buffer(Fl_Text_Display *d) {
  return d->buffer();
}

void go_fltk_TextDisplay_insert_position(Fl_Text_Display *d, int newPos) {
  d->insert_position(newPos);
}

int go_fltk_TextDisplay_get_insert_position(Fl_Text_Display *d) {
  return d->insert_position();
}

void go_fltk_TextDisplay_insert_text(Fl_Text_Display *d, const char *text) {
  d->insert(text);
}

void go_fltk_TextDisplay_overstrike(Fl_Text_Display *d, const char* text) {
  d->overstrike(text);
}

class GText_Editor : public EventHandler<Fl_Text_Editor> {
public:
  GText_Editor(int x, int y, int w, int h, const char* label)
    : EventHandler<Fl_Text_Editor>(x, y, w, h, label) {}
};

GText_Editor *go_fltk_new_TextEditor(int x, int y, int w, int h, const char *text) {
  return new GText_Editor(x, y, w, h, text);
}

void go_fltk_TextEditor_copy(Fl_Text_Editor *e) {
  Fl_Text_Editor::kf_copy(0, e);
}

void go_fltk_TextEditor_insert(Fl_Text_Editor *e) {
  Fl_Text_Editor::kf_insert(0, e);
}

void go_fltk_TextEditor_cut(Fl_Text_Editor *e) {
  Fl_Text_Editor::kf_cut(0, e);
}

void go_fltk_TextEditor_delete(Fl_Text_Editor *e) {
  Fl_Text_Editor::kf_delete(0, e);
}

void go_fltk_TextEditor_paste(Fl_Text_Editor *e) {
  Fl_Text_Editor::kf_paste(0, e);
}

void go_fltk_TextEditor_select_all(Fl_Text_Editor *e) {
  Fl_Text_Editor::kf_select_all(0, e);
}

Fl_Text_Buffer *go_fltk_new_TextBuffer(void) {
  return new Fl_Text_Buffer;
}

void go_fltk_TextBuffer_delete(Fl_Text_Buffer* b) {
  delete b;
}

void go_fltk_TextBuffer_set_text(Fl_Text_Buffer *b, const char *txt) {
  b->text(txt);
}

void go_fltk_TextBuffer_append(Fl_Text_Buffer *b, const char *txt) {
  b->append(txt);
}

const char *go_fltk_TextBuffer_text(Fl_Text_Buffer *b) {
  return b->text();
}

const char *go_fltk_TextBuffer_text_range(Fl_Text_Buffer *b, int start, int end) {
  return b->text_range(start, end);
}

void go_fltk_TextBuffer_highlight(Fl_Text_Buffer *b, int start, int end) {
  b->highlight(start, end);
}

void go_fltk_TextBuffer_unhighlight(Fl_Text_Buffer *b) {
  b->unhighlight();
}

void go_fltk_TextBuffer_replace(Fl_Text_Buffer *b, int start, int end, const char *text) {
  b->replace(start, end, text);
}

void go_fltk_TextBuffer_replace_selection(Fl_Text_Buffer *b, const char *text) {
  b->replace_selection(text);
}

int go_fltk_TextBuffer_search_forward(Fl_Text_Buffer *b, int start, const char *searchString, int *foundPos, int matchCase) {
  return b->search_forward(start, searchString, foundPos, matchCase);
}

int go_fltk_TextBuffer_search_backward(Fl_Text_Buffer *b, int start, const char *searchString, int *foundPos, int matchCase) {
  return b->search_backward(start, searchString, foundPos, matchCase);
}

void go_fltk_TextBuffer_select(Fl_Text_Buffer *b, int start, int end) {
  b->select(start, end);
}

int go_fltk_TextBuffer_selected(Fl_Text_Buffer *b) {
  return b->selected();
}

int go_fltk_TextBuffer_selection_position(Fl_Text_Buffer *b, int *start, int *end) {
  return b->selection_position(start, end);
}

char* go_fltk_TextBuffer_selection_text(Fl_Text_Buffer *b) {
  return b->selection_text();
}

void go_fltk_TextBuffer_unselect(Fl_Text_Buffer *b) {
  b->unselect();
}

int go_fltk_TextBuffer_tab_distance(Fl_Text_Buffer *b) {
  return b->tab_distance();
}

void go_fltk_TextBuffer_set_tab_distance(Fl_Text_Buffer *b, int tabDist) {
  b->tab_distance(tabDist);
}
