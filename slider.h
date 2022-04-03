#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GSlider GSlider;
  typedef struct GValue_Slider GValue_Slider;

  extern GSlider* go_fltk_new_Slider(int x, int y, int w, int h, const char* label);

  extern GValue_Slider* go_fltk_new_Value_Slider(int x, int y, int w, int h, const char* label);

  extern void go_fltk_Value_Slider_set_textfont(GValue_Slider* slider, int font);

  extern void go_fltk_Value_Slider_set_textsize(GValue_Slider* slider, int size);

#ifdef __cplusplus
}
#endif

