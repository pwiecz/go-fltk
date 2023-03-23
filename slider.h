#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GSlider GSlider;
  typedef struct Fl_Value_Slider Fl_Value_Slider;
  typedef struct GValue_Slider GValue_Slider;

  extern GSlider* go_fltk_new_Slider(int x, int y, int w, int h, const char* label);

  extern const unsigned char go_FL_VERT_SLIDER;
  extern const unsigned char go_FL_HOR_SLIDER;
  extern const unsigned char go_FL_VERT_FILL_SLIDER;
  extern const unsigned char go_FL_HOR_FILL_SLIDER;
  extern const unsigned char go_FL_VERT_NICE_SLIDER;
  extern const unsigned char go_FL_HOR_NICE_SLIDER;

  extern GValue_Slider* go_fltk_new_Value_Slider(int x, int y, int w, int h, const char* label);

  extern void go_fltk_Value_Slider_set_textfont(Fl_Value_Slider* slider, int font);

  extern void go_fltk_Value_Slider_set_textsize(Fl_Value_Slider* slider, int size);
  
#ifdef __cplusplus
}
#endif

