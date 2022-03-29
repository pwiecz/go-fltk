#include "helpview.h"
#include <FL/Fl_Help_View.H>

// Implemented:
//  Create and destroy HelpView
//  directory()
//  filename()
//  find()
//  load()

// TODO:
//  handle()
//  leftline()
//  link()
//  resize()
//  scrollbar_size()
//  select_all()
//  size()
//  textcolor()
//  textfont()
//  textsize()
//  title()
//  topline()
//  value()

Fl_Help_View *go_fltk_new_HelpView(int x, int y, int w, int h, const char *text) {
	return new Fl_Help_View(x, y, w, h, text);
}

void go_fltk_HelpView_delete(Fl_Help_View *h) {
	delete h;
}

void go_fltk_HelpView_load(Fl_Help_View *h, const char *f) {
	h->load(f);
}

const char *go_fltk_HelpView_directory(Fl_Help_View *h) {
	return h->directory();
}

const char *go_fltk_HelpView_filename(Fl_Help_View *h) {
	return h->filename();
}

int go_fltk_HelpView_find(Fl_Help_View *h, const char *s, int p) {
	return h->find(s, p);
}
