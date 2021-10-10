#pragma once

#include <stdint.h>


#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GGlWindow GGlWindow;

  extern GGlWindow* go_fltk_new_GlWindow(int x, int y, int w, int h, uintptr_t drawFunId);

  extern void go_fltk_Gl_Window_make_current(GGlWindow* w);
  extern char go_fltk_Gl_Window_context_valid(GGlWindow* w);
  extern char go_fltk_Gl_Window_valid(GGlWindow* w);
  extern int go_fltk_Gl_Window_can_do(GGlWindow* w);
  extern void go_fltk_Gl_Window_set_event_handler(GGlWindow* w, int handlerId);
  extern void go_fltk_Gl_Window_set_resize_handler(GGlWindow* w, uintptr_t handlerId);
  extern void go_fltk_Gl_Window_set_mode(GGlWindow* w, int mode);

#ifdef __cplusplus
}
#endif
