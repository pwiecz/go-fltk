#include "browser.h"

#include <FL/Fl_Browser.H>
#include <FL/Fl_Check_Browser.H>
#include <FL/Fl_Hold_Browser.H>
#include <FL/Fl_Multi_Browser.H>
#include <FL/Fl_Select_Browser.H>

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

void go_fltk_Browser_add(Fl_Browser* b, const char *str, uintptr_t id) {
	b->add(str, (void *)id);
}

int go_fltk_Browser_topline(Fl_Browser* b) {
	return b->topline();
}

void go_fltk_Browser_set_bottomline(Fl_Browser* b, int i) {
	b->bottomline(i);
}

void go_fltk_Browser_set_middleline(Fl_Browser* b, int i) {
	b->middleline(i);
}

void go_fltk_Browser_set_topline(Fl_Browser* b, int i) {
	b->topline(i);
}

void go_fltk_Browser_clear(Fl_Browser *b) {
        b->clear();
}  

uintptr_t go_fltk_Browser_data(Fl_Browser* b, int line) {
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

void go_fltk_Browser_remove(Fl_Browser* b, int i) {
	b->remove(i);
}

char go_fltk_Browser_column_char(Fl_Browser* b) {
	return b->column_char();
}

void go_fltk_Browser_set_column_char(Fl_Browser* b, char c) {
	b->column_char(c);
}

void go_fltk_Browser_hide_line(Fl_Browser* b, int line) {
	b->hide(line);
}

int go_fltk_Browser_size(Fl_Browser* b) {
	return b->size();
}

Fl_Image* go_fltk_Browser_icon(Fl_Browser* b, int line) {
	return b->icon(line);
}

void go_fltk_Browser_set_icon(Fl_Browser* b, int line, Fl_Image *i) {
	b->icon(line, i);
}

char go_fltk_Browser_format_char(Fl_Browser* b) {
	return b->format_char();
}

void go_fltk_Browser_set_format_char(Fl_Browser* b, char c) {
	b->format_char(c);
}

int go_fltk_Browser_displayed(Fl_Browser* b, int line) {
	return b->displayed(line);
}

int go_fltk_Browser_value(Fl_Browser* b) {
	return b->value();
}

void go_fltk_Browser_set_value(Fl_Browser* b, int line) {
	b->value(line);
}

const char* go_fltk_Browser_text(Fl_Browser* b, int line) {
	return b->text(line);
}

void go_fltk_Browser_set_column_widths(Fl_Browser* b, const int *arr) {
	b->column_widths(arr);
}

int go_fltk_Browser_select(Fl_Browser* b, int line, int val) {
	return b->select(line, val);
}

int go_fltk_Browser_selected(Fl_Browser* b, int line) {
	return b->selected(line);
}

class GCheckBrowser : public EventHandler<Fl_Check_Browser> {
public:
  GCheckBrowser(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Check_Browser>(x, y, w, h, label) {}
};

GCheckBrowser *go_fltk_new_Check_Browser(int x, int y, int w, int h, const char *text) {
  return new GCheckBrowser(x, y, w, h, text);
}
void go_fltk_Check_Browser_add(Fl_Check_Browser *b, const char *s, int checked) {
  b->add(s, checked);
}
void go_fltk_Check_Browser_set_checked(Fl_Check_Browser *b, int item,
                                       int checked) {
  b->checked(item, checked);
}
int go_fltk_Check_Browser_is_checked(Fl_Check_Browser *b, int item) {
  return b->checked(item);
}
int go_fltk_Check_Browser_nchecked(Fl_Check_Browser *b) {
  return b->nchecked();
}
void go_fltk_Check_Browser_remove(Fl_Check_Browser *b, int item) {
  b->remove(item);
}
void go_fltk_Check_Browser_clear(Fl_Check_Browser *b) {
  b->clear();  
}
int go_fltk_Check_Browser_nitems(Fl_Check_Browser *b) {
  return b->nitems();
}  
const char* go_fltk_Check_Browser_text(Fl_Check_Browser *b, int item) {
  return b->text(item);
}  
