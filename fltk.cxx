#include "fltk.h"

#include "Fl/Fl.H"

#include "_cgo_export.h"

static void lock() { Fl::lock(); }

static void unlock() {
  Fl::awake();
  Fl::unlock();
}

int go_fltk_run() { return Fl::run(); }
int go_fltk_lock() { return Fl::lock(); }
void awake_handler(void *data) {
  _go_awakeHandler(data);
}
int go_fltk_awake(void *id) {
  return Fl::awake(awake_handler, id);
}

int go_fltk_event_button1() { return Fl::event_button1(); }
int go_fltk_event_x() { return Fl::event_x(); }
int go_fltk_event_y() { return Fl::event_y(); }
int go_fltk_event_dx() { return Fl::event_dx(); }
int go_fltk_event_dy() { return Fl::event_dy(); }
int go_fltk_event_key() { return Fl::event_key(); }
int go_fltk_event_state() { return Fl::event_state(); }

