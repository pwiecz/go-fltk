#include "gl_window.h"

#include <Fl/Fl_Gl_Window.H>

#include "_cgo_export.h"


class GGlWindow : public Fl_Gl_Window {
public:
  GGlWindow(int x, int y, int w, int h, void* drawFunId)
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
    if (m_resizeHandlerId != nullptr) {
      _go_callbackHandler(m_resizeHandlerId);
    }
  }

  void set_event_handler(int handlerId) {
    m_eventHandlerId = handlerId;
  }

  void set_resize_handler(void* resizeHandlerId) {
    m_resizeHandlerId = resizeHandlerId;
  }
private:
  void* const m_drawFunId;
  int m_eventHandlerId = -1;
  void* m_resizeHandlerId = nullptr;
};


Fl_Gl_Window *go_fltk_new_GlWindow(int x, int y, int w, int h, void* drawFunId) {
  return new GGlWindow(x, y, w, h, drawFunId);
}
char go_fltk_Gl_Window_context_valid(Fl_Gl_Window* w) { 
  return w->context_valid(); 
}
char go_fltk_Gl_Window_valid(Fl_Gl_Window* w) {
  return w->valid(); 
}
void go_fltk_Gl_Window_set_event_handler(Fl_Gl_Window* w, int handlerId) {
  ((GGlWindow*)w)->set_event_handler(handlerId);
}
void go_fltk_Gl_Window_set_resize_handler(Fl_Gl_Window* w, void* handlerId) {
  ((GGlWindow*)w)->set_resize_handler(handlerId);
}
