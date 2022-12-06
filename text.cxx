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

void go_fltk_TextDisplay_set_buffer(GText_Display *d, Fl_Text_Buffer *buf) {
  d->buffer(buf);
}

void go_fltk_TextDisplay_set_wrap_mode(GText_Display *b, int wrap, int wrapMargin) {
  b->wrap_mode(wrap, wrapMargin);
}

int go_fltk_TextDisplay_move_right(GText_Display *b) {
  return b->move_right();
}

int go_fltk_TextDisplay_move_left(GText_Display *b) {
  return b->move_left();
}

int go_fltk_TextDisplay_move_up(GText_Display *b) {
  return b->move_up();
}

int go_fltk_TextDisplay_move_down(GText_Display *b) {
  return b->move_down();
}

void go_fltk_TextDisplay_show_insert_position(GText_Display *b) {
  b->show_insert_position();
}

int go_fltk_TextDisplay_text_size(GText_Display *d) {
  return d->textsize();
}

void go_fltk_TextDisplay_set_text_size(GText_Display *d, int size) {
  d->textsize(size);
}

Fl_Text_Buffer *go_fltk_TextDisplay_buffer(GText_Display *d) {
  return d->buffer();
}

void go_fltk_TextDisplay_insert_position(GText_Display *d, int newPos) {
  d->insert_position(newPos);
}

int go_fltk_TextDisplay_get_insert_position(GText_Display *d) {
  return d->insert_position();
}

void go_fltk_TextDisplay_insert_text(GText_Display *d, const char *text) {
  d->insert(text);
}

void go_fltk_TextDisplay_overstrike(GText_Display *d, const char* text) {
  d->overstrike(text);
}

int go_fltk_TextDisplay_total_lines(GText_Display *d) {
  Fl_Text_Buffer* buffer = d->buffer();
  int endPos = go_fltk_TextDisplay_move_end(d);
  return buffer->count_lines(0, endPos);
}

// fltk has position_to_line but not this one
int go_fltk_TextDisplay_line_to_position(GText_Display *d, int lineNum) {
  int currentPos, currentLine, movedown, startPos;
  Fl_Text_Buffer* buff = d->buffer();
  int total_lines = go_fltk_TextDisplay_total_lines(d);
  if (lineNum > total_lines) {
    return -1;
  }
  currentPos = d->insert_position();
  currentLine = buff->count_lines(0, currentPos);

  while ( true )  {
    if ( currentLine == lineNum ) {
      break;
    }
    movedown = d->move_down();
    if (movedown == 0) {
      d->insert_position(0);
      currentLine = 0;
    } else {
      currentPos = d->insert_position();
      currentLine++;
    }
  }

  return d->line_start(currentPos);
}

// Move insert to the end of display
int go_fltk_TextDisplay_move_end(GText_Display *d) {
  int movedown;
  while ( movedown = d->move_down() == 1) {
  }
  int currentPos = d->insert_position();
  int lineEnd;
  lineEnd = d->line_end(currentPos, false);
  d->insert_position(lineEnd);
  return lineEnd;
}

class GText_Editor : public EventHandler<Fl_Text_Editor> {
public:
  GText_Editor(int x, int y, int w, int h, const char* label)
    : EventHandler<Fl_Text_Editor>(x, y, w, h, label) {}
};

GText_Editor *go_fltk_new_TextEditor(int x, int y, int w, int h, const char *text) {
  return new GText_Editor(x, y, w, h, text);
}

void go_fltk_TextEditor_copy(GText_Editor *e) {
  Fl_Text_Editor::kf_copy(0, e);
}

void go_fltk_TextEditor_insert(GText_Editor *e) {
  Fl_Text_Editor::kf_insert(0, e);
}

void go_fltk_TextEditor_cut(GText_Editor *e) {
  Fl_Text_Editor::kf_cut(0, e);
}

void go_fltk_TextEditor_delete(GText_Editor *e) {
  Fl_Text_Editor::kf_delete(0, e);
}

void go_fltk_TextEditor_paste(GText_Editor *e) {
  Fl_Text_Editor::kf_paste(0, e);
}

void go_fltk_TextEditor_select_all(GText_Editor *e) {
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
