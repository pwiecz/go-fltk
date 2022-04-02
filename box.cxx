#include "box.h"

#include <FL/Fl_Box.H>

#include "_cgo_export.h"


class GBox : public Fl_Box {
public:
  GBox(int boxType, int x, int y, int w, int h, const char* label)
    : Fl_Box((Fl_Boxtype)boxType, x, y, w, h, label) {
  }

  int handle(int event) final {
    if (m_eventHandlerId >= 0) {
      const int ret = _go_eventHandler(m_eventHandlerId, event);
      if (ret != 0) {
	return ret;
      }
    }
    return Fl_Box::handle(event);
  }
  
  void set_event_handler(int handlerId) {
    m_eventHandlerId = handlerId;
  }

private:
  int m_eventHandlerId = -1;
};

GBox *go_fltk_new_Box(int boxType, int x, int y, int w, int h, const char* label) {
  return new GBox((Fl_Boxtype)boxType, x, y, w, h, label);
}

void go_fltk_Box_set_event_handler(GBox* b, int handlerId) {
  b->set_event_handler(handlerId);
}
