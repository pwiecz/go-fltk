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

void go_fltk_Flex_set_gap(GFlex* flex, int spacing) {
    flex->gap(spacing);
}

void go_fltk_Flex_fixed(GFlex* flex, void* w, int size) {
    flex->fixed((Fl_Widget *)w, size);
}

extern void go_fltk_Flex_end(GFlex *flex) {
    flex->end();
}

const unsigned char go_FL_FLEX_COLUMN = Fl_Flex::COLUMN;
const unsigned char go_FL_FLEX_ROW = Fl_Flex::ROW;
