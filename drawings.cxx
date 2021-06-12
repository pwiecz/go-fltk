#include "drawings.h"

#include <FL/fl_draw.H>
#include <FL/Enumerations.H>


void go_fltk_color(unsigned int color) {
  fl_color((Fl_Color)color);
}
void go_fltk_draw(const char* text, int x, int y, int w, int h, unsigned int align) {
  fl_draw(text, x, y, w, h, (Fl_Align)align);
}
void go_fltk_draw_box(int boxType, int x, int y, int w, int h, unsigned int color) {
  fl_draw_box(Fl_Boxtype(boxType), x, y, w, h, (Fl_Color)color);
}

