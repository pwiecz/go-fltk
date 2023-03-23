#include "slider.h"

#include <FL/Fl_Slider.H>
#include <FL/Fl_Value_Slider.H>

#include "event_handler.h"


class GSlider : public EventHandler<Fl_Slider> {
public:
  GSlider(int x, int y, int w, int h, const char* label)
    : EventHandler<Fl_Slider>(x, y, w, h, label) {}
};

GSlider* go_fltk_new_Slider(int x, int y, int w, int h, const char* label) {
  return new GSlider(x, y, w, h, label);
}

class GValue_Slider : public EventHandler<Fl_Value_Slider> {
public:
  GValue_Slider(int x, int y, int w, int h, const char* label)
    : EventHandler<Fl_Value_Slider>(x, y, w, h, label) {}
};

GValue_Slider* go_fltk_new_Value_Slider(int x, int y, int w, int h, const char* label) {
  return new GValue_Slider(x, y, w, h, label);
}

void go_fltk_Value_Slider_set_textfont(Fl_Value_Slider* slider, int font) {
  slider->textfont(font);
}

void go_fltk_Value_Slider_set_textsize(Fl_Value_Slider* slider, int size) {
  slider->textsize(size);
}

const unsigned char go_FL_VERT_SLIDER = FL_VERT_SLIDER;
const unsigned char go_FL_HOR_SLIDER = FL_HOR_SLIDER;
const unsigned char go_FL_VERT_FILL_SLIDER = FL_VERT_FILL_SLIDER;
const unsigned char go_FL_HOR_FILL_SLIDER = FL_HOR_FILL_SLIDER;
const unsigned char go_FL_VERT_NICE_SLIDER = FL_VERT_NICE_SLIDER;
const unsigned char go_FL_HOR_NICE_SLIDER = FL_HOR_NICE_SLIDER;
