#include "box.h"

#include <FL/Fl_Box.H>

#include "event_handler.h"


class GBox : public EventHandler<Fl_Box> {
public:
  GBox(int boxType, int x, int y, int w, int h, const char* label)
    : EventHandler<Fl_Box>((Fl_Boxtype)boxType, x, y, w, h, label) {}
};

GBox *go_fltk_new_Box(int boxType, int x, int y, int w, int h, const char* label) {
  return new GBox((Fl_Boxtype)boxType, x, y, w, h, label);
}

void go_fltk_Box_set_event_handler(GBox* b, int handlerId) {
  b->set_event_handler(handlerId);
}
