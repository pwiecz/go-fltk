#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Gl_Window Fl_Gl_Window;

  extern Fl_Gl_Window* go_fltk_new_GlWindow(int x, int y, int w, int h, void* drawFunId);

  extern char go_fltk_Gl_Window_context_valid(Fl_Gl_Window* w);
  extern char go_fltk_Gl_Window_valid(Fl_Gl_Window* w);
  extern void go_fltk_Gl_Window_set_event_handler(Fl_Gl_Window* w, int handlerId);

#ifdef __cplusplus
}
#endif
