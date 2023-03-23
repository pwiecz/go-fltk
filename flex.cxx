#include "flex.h"

#include <FL/Fl_Flex.H>

#include "event_handler.h"

class GFlex : public EventHandler<Fl_Flex> {
public:
    GFlex(int x, int y, int w, int h, const char* label)
        : EventHandler<Fl_Flex>(x, y, w, h, label) {}
};

GFlex* go_fltk_new_Flex(int x, int y, int w, int h, const char* label) {
    return new GFlex(x, y, w, h, label);
}

void go_fltk_Flex_set_gap(Fl_Flex* flex, int spacing) {
    flex->gap(spacing);
}

void go_fltk_Flex_fixed(Fl_Flex* flex, Fl_Widget* w, int size) {
    flex->fixed(w, size);
}

extern void go_fltk_Flex_end(Fl_Flex *flex) {
    flex->end();
}

const unsigned char go_FL_FLEX_COLUMN = Fl_Flex::COLUMN;
const unsigned char go_FL_FLEX_ROW = Fl_Flex::ROW;
