#include "browser.h"

#include <FL/Fl_Browser.H>

#include "event_handler.h"


// Implemented:
//  Create Browser
//  add()
//  bottomline()
//  clear()

// TODO:
//  column_char()
//  column_widths()
//  data() Need to implement the data in general when creating the widget
//  display()
//  displayed()
//  format_char()
//  hide()
//  icon()
//  insert()


class GBrowser : public EventHandler<Fl_Browser> {
public:
  GBrowser(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Browser>(x, y, w, h, label) {}
};

GBrowser *go_fltk_new_Browser(int x, int y, int w, int h, const char *text) {
	return new GBrowser(x, y, w, h, text);
}

void go_fltk_Browser_add(GBrowser *b, const char *str, void *d=0) {
	b->add(str, d);
}

void go_fltk_Browser_bottomline(GBrowser *b, int i) {
	b->bottomline(i);
}

void go_fltk_Browser_clear(GBrowser *b) {
	b->clear();
}
