#include "browser.h"

#include <FL/Fl_Browser.H>
#include <FL/Fl_Select_Browser.H>
#include <FL/Fl_Hold_Browser.H>
#include <FL/Fl_Multi_Browser.H>

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

void go_fltk_Browser_add(GBrowser *b, const char *str, uintptr_t id) {
	b->add(str, (void *)id);
}

void go_fltk_Browser_bottomline(GBrowser *b, int i) {
	b->bottomline(i);
}

void go_fltk_Browser_clear(GBrowser *b) {
	b->clear();
}

uintptr_t go_fltk_Browser_data(GBrowser *b, int line) {
	return (uintptr_t)b->data(line);
}

int go_fltk_Browser_value(GBrowser *b) {
	return b->value();
}

void go_fltk_Browser_set_value(GBrowser *b, int line) {
	b->value(line);
}

class GSelectBrowser : public EventHandler<Fl_Select_Browser> {
public:
  GSelectBrowser(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Select_Browser>(x, y, w, h, label) {}
};

GSelectBrowser *go_fltk_new_Select_Browser(int x, int y, int w, int h, const char *text) {
	return new GSelectBrowser(x, y, w, h, text);
}

class GHoldBrowser : public EventHandler<Fl_Hold_Browser> {
public:
  GHoldBrowser(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Hold_Browser>(x, y, w, h, label) {}
};

GHoldBrowser *go_fltk_new_Hold_Browser(int x, int y, int w, int h, const char *text) {
	return new GHoldBrowser(x, y, w, h, text);
}

class GMultiBrowser : public EventHandler<Fl_Multi_Browser> {
public:
  GMultiBrowser(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Multi_Browser>(x, y, w, h, label) {}
};

GMultiBrowser *go_fltk_new_Multi_Browser(int x, int y, int w, int h, const char *text) {
	return new GMultiBrowser(x, y, w, h, text);
}

