#include "text_editor.h"

#include "Fl/Fl_Text_Editor.H"


Fl_Text_Editor *go_fltk_new_TextEditor(int x, int y, int w, int h, const char *text) {
  return new Fl_Text_Editor(x, y, w, h, text);
}
