#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  extern void go_fltk_enable_tooltips();

  extern void go_fltk_disable_tooltips();

  extern int go_fltk_tooltips_enabled();

  extern float go_fltk_tooltip_delay();

  extern void go_fltk_set_tooltip_delay(float delay);

#ifdef __cplusplus
}
#endif
