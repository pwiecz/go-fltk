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
void go_fltk_unlock() { Fl::unlock(); }

void awake_handler(void *data) {
  _go_awakeHandler(uintptr_t(data));
}
void go_fltk_awake_null_message() {
  Fl::awake();
}
int go_fltk_awake(uintptr_t id) {
  return Fl::awake(awake_handler, (void*)id);
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
int go_fltk_screen_count() {
  return Fl::screen_count();
}
float go_fltk_screen_scale(int screenNum) {
  return Fl::screen_scale(screenNum);
}
void go_fltk_set_screen_scale(int screenNum, float scale) {
  Fl::screen_scale(screenNum, scale);
}
void go_fltk_set_keyboard_screen_scaling(int value) {
  Fl::keyboard_screen_scaling(value);
}
