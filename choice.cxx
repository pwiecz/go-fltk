#include "choice.h"

#include <FL/Fl_Choice.H>

#include "event_handler.h"

class GChoice : public EventHandler<Fl_Choice> {
public:
  GChoice(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Choice>(x, y, w, h, label) {}
};

GChoice *go_fltk_new_Choice(int x, int y, int w, int h, const char *label) {
  return new GChoice(x, y, w, h, label);
}

