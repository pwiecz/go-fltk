#include "helpview.h"

#include <FL/Fl_Help_View.H>

#include "event_handler.h"


// Implemented:
//  Create HelpView
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


class GHelp_View : public EventHandler<Fl_Help_View> {
public:
  GHelp_View(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Help_View>(x, y, w, h, label) {}
};

GHelp_View *go_fltk_new_HelpView(int x, int y, int w, int h, const char *text) {
	return new GHelp_View(x, y, w, h, text);
}

void go_fltk_HelpView_load(GHelp_View *h, const char *f) {
	h->load(f);
}

const char *go_fltk_HelpView_directory(GHelp_View *h) {
	return h->directory();
}

const char *go_fltk_HelpView_filename(GHelp_View *h) {
	return h->filename();
}

int go_fltk_HelpView_find(GHelp_View *h, const char *s, int p) {
	return h->find(s, p);
}
