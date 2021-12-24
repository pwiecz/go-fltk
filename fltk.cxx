#include "fltk.h"

#include "FL/Fl.H"

#include "_cgo_export.h"


static void lock() { Fl::lock(); }

static void unlock() {
  Fl::awake();
  Fl::unlock();
}

void go_fltk_init_styles(void) {
    fl_define_FL_ROUND_UP_BOX();
    fl_define_FL_SHADOW_BOX();
    fl_define_FL_ROUNDED_BOX();
    fl_define_FL_RFLAT_BOX();
    fl_define_FL_RSHADOW_BOX();
    fl_define_FL_DIAMOND_BOX();
    fl_define_FL_OVAL_BOX();
    fl_define_FL_PLASTIC_UP_BOX();
    fl_define_FL_GTK_UP_BOX();
    fl_define_FL_GLEAM_UP_BOX();
    fl_define_FL_SHADOW_LABEL();
    fl_define_FL_ENGRAVED_LABEL();
    fl_define_FL_EMBOSSED_LABEL();
    fl_define_FL_MULTI_LABEL();
    fl_define_FL_ICON_LABEL();
    fl_define_FL_IMAGE_LABEL();
    Fl::use_high_res_GL(1);
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
