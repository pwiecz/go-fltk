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
//  middleline()
//  topline()
//  clear()
//  column_char()
//  format_char()
//  displayed()
//  icon()
//  hide()
//  size()
//  value()
//  column_widths()
//  data()

// TODO:
//  insert()
//  lineposition()
//  load()



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

int go_fltk_Browser_topline(GBrowser *b) {
	return b->topline();
}

void go_fltk_Browser_set_bottomline(GBrowser *b, int i) {
	b->bottomline(i);
}

void go_fltk_Browser_set_middleline(GBrowser *b, int i) {
	b->middleline(i);
}

void go_fltk_Browser_set_topline(GBrowser *b, int i) {
	b->topline(i);
}

void go_fltk_Browser_clear(GBrowser *b) {
	b->clear();
}

uintptr_t go_fltk_Browser_data(GBrowser *b, int line) {
	return (uintptr_t)b->data(line);
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


void go_fltk_Browser_remove(GBrowser *b, int i) {
	b->remove(i);
}

char go_fltk_Browser_column_char(GBrowser *b) {
	return b->column_char();
}

void go_fltk_Browser_set_column_char(GBrowser *b, char c) {
	b->column_char(c);
}

void go_fltk_Browser_hide_line(GBrowser *b, int line) {
	b->hide(line);
}

int go_fltk_Browser_size(GBrowser *b) {
	return b->size();
}

Fl_Image* go_fltk_Browser_icon(GBrowser *b, int line) {
	return b->icon(line);
}

void go_fltk_Browser_set_icon(GBrowser *b, int line, Fl_Image *i) {
	b->icon(line, i);
}

char go_fltk_Browser_format_char(GBrowser *b) {
	return b->format_char();
}

void go_fltk_Browser_set_format_char(GBrowser *b, char c) {
	b->format_char(c);
}

int go_fltk_Browser_displayed(GBrowser *b, int line) {
	return b->displayed(line);
}

int go_fltk_Browser_value(GBrowser *b) {
	return b->value();
}

void go_fltk_Browser_set_value(GBrowser *b, int line) {
	b->value(line);
}

const char* go_fltk_Browser_text(GBrowser *b, int line) {
	return b->text(line);
}

void go_fltk_Browser_set_column_widths(GBrowser *b, const int *arr) {
	b->column_widths(arr);
}
