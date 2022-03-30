#include "fltk.h"

#include <FL/Fl.H>

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

typedef void (*f_array)(int, int, int, int, unsigned int);

// idk if it's worth adding all the boxtypes this way, but this is enough to
// override all the default designs
f_array currentBoxTypeCb[56] = {&_go_drawBox0,  &_go_drawBox1,  &_go_drawBox2,  &_go_drawBox3,
                          &_go_drawBox4,  &_go_drawBox5,  &_go_drawBox6,  &_go_drawBox7,
                          &_go_drawBox8,  &_go_drawBox9,  &_go_drawBox10, &_go_drawBox11,
                          &_go_drawBox12, &_go_drawBox13, &_go_drawBox14, &_go_drawBox15,
                          &_go_drawBox16, &_go_drawBox17, &_go_drawBox18};

int go_fltk_set_scheme(const char *scheme) {
  return Fl::scheme(scheme);
}

void go_fltk_set_background_color(unsigned char r, unsigned char g, unsigned char b) {
    Fl::background(r, g, b);
}

void go_fltk_set_background2_color(unsigned char r, unsigned char g, unsigned char b) {
    Fl::background2(r, g, b);
}

void go_fltk_set_boxtype(int i, int x, int y, int w, int h) {
	Fl::set_boxtype((Fl_Boxtype)i, currentBoxTypeCb[i], x, y, w, h);
}

void go_fltk_set_foreground_color(unsigned char r, unsigned char g, unsigned char b) {
    Fl::foreground(r, g, b);
}

void go_fltk_set_color(unsigned int col, unsigned char r, unsigned char g, unsigned char b) {
    Fl::set_color((Fl_Color)col, r, g, b);
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
