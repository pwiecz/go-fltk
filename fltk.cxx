#include "fltk.h"

#include "FL/Fl.H"

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

void go_fltk_copy(const char* data, int len, int destination) {
  Fl::copy(data, len, destination);
}

void go_fltk_screen_work_area(int *x, int *y, int *w, int *h, int n) {
  Fl::screen_work_area(*x, *y, *w, *h, n);
}
void go_fltk_screen_dpi(float *w, float *h, int n) {
  Fl::screen_dpi(*w, *h, n);
}
