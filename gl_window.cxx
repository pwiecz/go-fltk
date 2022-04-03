#include "gl_window.h"

#include <FL/Fl_Gl_Window.H>

#include "event_handler.h"

#include "_cgo_export.h"


class GGlWindow : public EventHandler<Fl_Gl_Window> {
public:
  GGlWindow(int x, int y, int w, int h, uintptr_t drawFunId)
    : EventHandler<Fl_Gl_Window>(x, y, w, h)
    , m_drawFunId(drawFunId) {
    mode(FL_OPENGL3 | FL_MULTISAMPLE);
  }
  void draw() final {
    _go_callbackHandler(m_drawFunId);
  }

  void resize(int x, int y, int w, int h) final {
    Fl_Gl_Window::resize(x, y, w, h);
    if (m_resizeHandlerId != 0) {
      _go_callbackHandler(m_resizeHandlerId);
    }
  }

  void set_resize_handler(uintptr_t resizeHandlerId) {
    m_resizeHandlerId = resizeHandlerId;
  }

private:
  uintptr_t const m_drawFunId;
  uintptr_t m_resizeHandlerId = 0;
};


GGlWindow *go_fltk_new_GlWindow(int x, int y, int w, int h, uintptr_t drawFunId) {
  return new GGlWindow(x, y, w, h, drawFunId);
}
void go_fltk_Gl_Window_make_current(GGlWindow* w) {
  w->make_current();
}
char go_fltk_Gl_Window_context_valid(GGlWindow* w) { 
  return w->context_valid(); 
}
char go_fltk_Gl_Window_valid(GGlWindow* w) {
  return w->valid(); 
}
int go_fltk_Gl_Window_can_do(GGlWindow* w) {
  return w->can_do();
}
void go_fltk_Gl_Window_set_event_handler(GGlWindow* w, int handlerId) {
  ((GGlWindow*)w)->set_event_handler(handlerId);
}
void go_fltk_Gl_Window_set_resize_handler(GGlWindow* w, uintptr_t handlerId) {
  w->set_resize_handler(handlerId);
}
void go_fltk_Gl_Window_set_mode(GGlWindow* w, int m) {
  w->mode(m);
}
