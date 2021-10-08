#include "gl_window.h"

#include <FL/Fl_Gl_Window.H>

#include "_cgo_export.h"


class GGlWindow : public Fl_Gl_Window {
public:
  GGlWindow(int x, int y, int w, int h, uintptr_t drawFunId)
    : Fl_Gl_Window(x, y, w, h)
    , m_drawFunId(drawFunId) {
    mode(FL_OPENGL3 | FL_MULTISAMPLE);
  }
  void draw() final {
    _go_callbackHandler(m_drawFunId);
  }

  int handle(int event) final {
    if (m_eventHandlerId >= 0) {
      const int ret = _go_eventHandler(m_eventHandlerId, event);
      if (ret != 0) {
	return ret;
      }
    }
    return Fl_Gl_Window::handle(event);
  }
  
  void resize(int x, int y, int w, int h) final {
    Fl_Gl_Window::resize(x, y, w, h);
    if (m_resizeHandlerId != 0) {
      _go_callbackHandler(m_resizeHandlerId);
    }
  }

  void set_event_handler(int handlerId) {
    m_eventHandlerId = handlerId;
  }

  void set_resize_handler(uintptr_t resizeHandlerId) {
    m_resizeHandlerId = resizeHandlerId;
  }
private:
  uintptr_t const m_drawFunId;
  int m_eventHandlerId = -1;
  uintptr_t m_resizeHandlerId = 0;
};


GGlWindow *go_fltk_new_GlWindow(int x, int y, int w, int h, uintptr_t drawFunId) {
  return new GGlWindow(x, y, w, h, drawFunId);
}
char go_fltk_Gl_Window_context_valid(GGlWindow* w) { 
  return w->context_valid(); 
}
char go_fltk_Gl_Window_valid(GGlWindow* w) {
  return w->valid(); 
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
