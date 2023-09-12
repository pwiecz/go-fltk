#include "text.h"

#include <FL/Fl_Text_Display.H>

#include <FL/Fl_Text_Editor.H>

#include <FL/Fl_Text_Buffer.H>

#include "event_handler.h"
#include "_cgo_export.h"


// --- Text Display ---

class GText_Display : public EventHandler<Fl_Text_Display> {
public:
  GText_Display(int x, int y, int w, int h, const char *label)
      : EventHandler<Fl_Text_Display>(x, y, w, h, label) {}

  // make xy_to_position() public
  int xy_to_position(int x, int y) {
    return EventHandler<Fl_Text_Display>::xy_to_position(x, y);
  }  
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

int go_fltk_TextDisplay_xy_to_position(Fl_Text_Display *d, int x, int y) {
  return ((GText_Display*) d)->xy_to_position(x, y);
}

int go_fltk_TextDisplay_position_to_xy(Fl_Text_Display *d, int pos, int *x, int *y) {
  return d->position_to_xy(pos, x, y);
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

void go_fltk_TextDisplay_hide_cursor(Fl_Text_Display *d) {
  d->hide_cursor();
}

unsigned int go_fltk_TextDisplay_text_color(Fl_Text_Display *d) {
  return d->textcolor();
}

void go_fltk_TextDisplay_set_text_color(Fl_Text_Display *d, unsigned int color) {
  d->textcolor(color);
}

int go_fltk_TextDisplay_text_font(Fl_Text_Display *d) {
  return d->textfont();
}

void go_fltk_TextDisplay_set_text_font(Fl_Text_Display *d, int font) {
  d->textfont(font);
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

void go_fltk_TextDisplay_set_linenumber_width(Fl_Text_Display *d, int width) {
  d->linenumber_width(width);
}

void go_fltk_TextDisplay_set_linenumber_font(Fl_Text_Display *d, int val) {
  d->linenumber_font(val);
}

void go_fltk_TextDisplay_set_linenumber_size(Fl_Text_Display *d, int val) {
  d->linenumber_size(val);
}

void go_fltk_TextDisplay_set_linenumber_fgcolor(Fl_Text_Display *d, unsigned int val) {
  d->linenumber_fgcolor(val);
}

void go_fltk_TextDisplay_set_linenumber_bgcolor(Fl_Text_Display *d, unsigned int val) {
  d->linenumber_bgcolor(val);
}

void go_fltk_TextDisplay_set_linenumber_align(Fl_Text_Display *d, int val) {
  d->linenumber_align(val);
}

// --- Text Editor ---

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

int go_fltk_TextEditor_undo(Fl_Text_Editor *e) {
  return Fl_Text_Editor::kf_undo(0, e);
}

int go_fltk_TextEditor_redo(Fl_Text_Editor *e) {
  return Fl_Text_Editor::kf_redo(0, e);
}

// --- Text Buffer ---

void modify_callback_handler(int pos, int nInserted, int nDeleted, int nRestyled, const char *deletedText, void *cbArg) {
  uintptr_t id = (uintptr_t)cbArg;
  _go_modifyCallbackHandler(id, pos, nInserted, nDeleted, nRestyled, (char*)deletedText);
}

Fl_Text_Buffer *go_fltk_new_TextBuffer(void) {
  return new Fl_Text_Buffer;
}

void go_fltk_TextBuffer_delete(Fl_Text_Buffer* b) {
  delete b;
}

void go_fltk_TextBuffer_add_modify_callback(Fl_Text_Buffer *b, uintptr_t handlerId) {
	b->add_modify_callback(modify_callback_handler, (void*)handlerId);
}

void go_fltk_TextBuffer_set_text(Fl_Text_Buffer *b, const char *txt) {
  b->text(txt);
}

void go_fltk_TextBuffer_append(Fl_Text_Buffer *b, const char *txt) {
  b->append(txt);
}

void go_fltk_TextBuffer_insert(Fl_Text_Buffer *b, int pos, const char *txt) {
  b->insert(pos, txt);
}

void go_fltk_TextBuffer_remove(Fl_Text_Buffer *b, int start, int end) {
  b->remove(start, end);
}

unsigned int go_fltk_TextBuffer_char_at(Fl_Text_Buffer *b, int pos) {
  return b->char_at(pos);
}

int go_fltk_TextBuffer_next_char(Fl_Text_Buffer *b, int ix) {
  return b->next_char(ix);
}

int go_fltk_TextBuffer_prev_char(Fl_Text_Buffer *b, int ix) {
  return b->prev_char(ix);
}

int go_fltk_TextBuffer_line_start(Fl_Text_Buffer *b, int ix) {
  return b->line_start(ix);
}

int go_fltk_TextBuffer_line_end(Fl_Text_Buffer *b, int ix) {
  return b->line_end(ix);
}

const char *go_fltk_TextBuffer_line_text(Fl_Text_Buffer *b, int ix) {
  return b->line_text(ix);
}

int go_fltk_TextBuffer_count_lines(Fl_Text_Buffer *b, int start, int end) {
  return b->count_lines(start, end);
}

int go_fltk_TextBuffer_skip_lines(Fl_Text_Buffer *b, int start, int nlines) {
  return b->skip_lines(start, nlines);
}

int go_fltk_TextBuffer_rewind_lines(Fl_Text_Buffer *b, int start, int nlines) {
  return b->rewind_lines(start, nlines);
}

int go_fltk_TextBuffer_length(Fl_Text_Buffer *b) {
  return b->length();
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

void go_fltk_TextDisplay_set_highlight_data(Fl_Text_Display *self, Fl_Text_Buffer *sbuff, unsigned int *color, int *font,    
                                     int *fontsz, unsigned *attr, unsigned int *bgcolor, int sz) { 
    Fl_Text_Display::Style_Table_Entry *stable = new Fl_Text_Display::Style_Table_Entry[sz];   
    if (!stable)                                                                               
        return;                                                                                
    for (int i = 0; i < sz; ++i) {                                                             
        stable[i] = (Fl_Text_Display::Style_Table_Entry){color[i], font[i], fontsz[i], attr[i], bgcolor[i]};                       
    }                                                                                          
    self->highlight_data(sbuff, stable, sz, 'A', 0, 0);                
}                                                                                              
