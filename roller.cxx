#include "roller.h"

#include <FL/Fl_Roller.H>

#include "event_handler.h"


class GRoller : public EventHandler<Fl_Roller> {
public:
  GRoller(int x, int y, int w, int h, const char* label)
    : EventHandler<Fl_Roller>(x, y, w, h, label) {}
};


GRoller* go_fltk_new_Roller(int x, int y, int w, int h, const char* label) {
  return new GRoller(x, y, w, h, label);
}
