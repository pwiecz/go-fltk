#include "slider.h"

#include <FL/Fl_Slider.H>
#include <FL/Fl_Value_Slider.H>


Fl_Slider* go_fltk_new_Slider(int x, int y, int w, int h, const char* label) {
  return new Fl_Slider(x, y, w, h, label);
}

Fl_Value_Slider* go_fltk_new_Value_Slider(int x, int y, int w, int h, const char* label) {
  return new Fl_Value_Slider(x, y, w, h, label);
}

void go_fltk_Value_Slider_set_textfont(Fl_Value_Slider* slider, int font) {
  slider->textfont(font);
}

void go_fltk_Value_Slider_set_textsize(Fl_Value_Slider* slider, int size) {
  slider->textsize(size);
}
