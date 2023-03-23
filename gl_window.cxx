#include "gl_window.h"

#include <array>
#include <vector>

#ifdef __linux__
#include <GL/glx.h>
#endif

#include <FL/Fl_Gl_Window.H>

#include "event_handler.h"

#include "_cgo_export.h"


class GGlWindow : public EventHandler<Fl_Gl_Window> {
public:
  GGlWindow(int x, int y, int w, int h, uintptr_t drawFunId)
    : EventHandler<Fl_Gl_Window>(x, y, w, h)
    , m_drawFunId(drawFunId) {
    mode(FL_OPENGL3 | FL_RGB | FL_DEPTH | FL_DOUBLE);
#ifdef __linux__
    // For some reason static storage is required here.    
    static std::array<int, 3> attributes{GLX_RGBA, GLX_DOUBLEBUFFER, 0};
    mode(attributes.data());
#endif    
  }
  void draw() final {
    _go_callbackHandler(m_drawFunId);
  }

private:
  uintptr_t const m_drawFunId;
  uintptr_t m_resizeHandlerId = 0;
};


GGlWindow *go_fltk_new_GlWindow(int x, int y, int w, int h, uintptr_t drawFunId) {
  return new GGlWindow(x, y, w, h, drawFunId);
}
void go_fltk_Gl_Window_make_current(Fl_Gl_Window* w) {
  w->make_current();
}
char go_fltk_Gl_Window_context_valid(Fl_Gl_Window* w) { 
  return w->context_valid(); 
}
char go_fltk_Gl_Window_valid(Fl_Gl_Window* w) {
  return w->valid(); 
}
int go_fltk_Gl_Window_can_do(Fl_Gl_Window* w) {
  return w->can_do();
}
void go_fltk_Gl_Window_set_mode(Fl_Gl_Window* w, int m) {
  w->mode(m);
#ifdef __linux__
  // For some reason static storage is required here.    
  static std::vector<int> attributes;
  attributes.clear();  
  if (m & FL_DOUBLE) {
    attributes.push_back(GLX_DOUBLEBUFFER);
  }
  if (m & FL_STEREO) {
    attributes.push_back(GLX_STEREO);
  }
  if (!(m & FL_INDEX)) {
    attributes.push_back(GLX_RGBA);
  }
  attributes.push_back(0);
  w->mode(attributes.data());
#endif    
}
