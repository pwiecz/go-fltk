#include "text.h"

#include <FL/Fl_Text_Display.H>
#include <FL/Fl_Text_Editor.H>
#include <FL/Fl_Text_Buffer.H>

Fl_Text_Display *go_fltk_new_TextDisplay(int x, int y, int w, int h, const char *text) {
  return new Fl_Text_Display(x, y, w, h, text);
}

void go_fltk_TextDisplay_set_buffer(Fl_Text_Display *d, Fl_Text_Buffer *buf) {
  d->buffer(buf);
}

Fl_Text_Buffer *go_fltk_TextDisplay_buffer(Fl_Text_Display *d) {
  return d->buffer();
}

Fl_Text_Editor *go_fltk_new_TextEditor(int x, int y, int w, int h, const char *text) {
  return new Fl_Text_Editor(x, y, w, h, text);
}

Fl_Text_Buffer *go_fltk_new_TextBuffer(void) {
  return new Fl_Text_Buffer;
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