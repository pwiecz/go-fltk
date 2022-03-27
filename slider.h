#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Slider Fl_Slider;
  typedef struct Fl_Value_Slider Fl_Value_Slider;

  extern Fl_Slider* go_fltk_new_Slider(int x, int y, int w, int h, const char* label);

  extern Fl_Value_Slider* go_fltk_new_Value_Slider(int x, int y, int w, int h, const char* label);

  extern void go_fltk_Value_Slider_set_textfont(Fl_Value_Slider* slider, int font);

  extern void go_fltk_Value_Slider_set_textsize(Fl_Value_Slider* slider, int size);

#ifdef __cplusplus
}
#endif

