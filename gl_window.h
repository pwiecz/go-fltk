#pragma once

#include <stdint.h>


#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GGlWindow GGlWindow;
  typedef struct Fl_Gl_Window Fl_Gl_Window;

  extern GGlWindow* go_fltk_new_GlWindow(int x, int y, int w, int h, uintptr_t drawFunId);

  extern void go_fltk_Gl_Window_make_current(Fl_Gl_Window* w);
  extern char go_fltk_Gl_Window_context_valid(Fl_Gl_Window* w);
  extern char go_fltk_Gl_Window_valid(Fl_Gl_Window* w);
  extern int go_fltk_Gl_Window_can_do(Fl_Gl_Window* w);
  extern void go_fltk_Gl_Window_set_mode(Fl_Gl_Window* w, int mode);

#ifdef __cplusplus
}
#endif
