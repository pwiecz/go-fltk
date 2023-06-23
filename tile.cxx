#include "tile.h"

#include <FL/Fl_Tile.H>

#include "event_handler.h"


class GTile : public EventHandler<Fl_Tile> {
public:
  GTile(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Tile>(x, y, w, h, label) {}
};

GTile *go_fltk_new_Tile(int x, int y, int w, int h, const char *label) {
  return new GTile(x, y, w, h, label);
}
