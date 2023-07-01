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

f_array currentBoxTypeCb[57] = {&_go_drawBox0,  &_go_drawBox1,  &_go_drawBox2,  &_go_drawBox3,
                          &_go_drawBox4,  &_go_drawBox5,  &_go_drawBox6,  &_go_drawBox7,
                          &_go_drawBox8,  &_go_drawBox9,  &_go_drawBox10, &_go_drawBox11,
                          &_go_drawBox12, &_go_drawBox13, &_go_drawBox14, &_go_drawBox15,
                          &_go_drawBox16, &_go_drawBox17, &_go_drawBox18, &_go_drawBox19,
                          &_go_drawBox20, &_go_drawBox21, &_go_drawBox22, &_go_drawBox23,
                          &_go_drawBox24, &_go_drawBox25, &_go_drawBox26, &_go_drawBox27,
                          &_go_drawBox28, &_go_drawBox29, &_go_drawBox30, &_go_drawBox31,
                          &_go_drawBox32, &_go_drawBox33, &_go_drawBox34, &_go_drawBox35,
                          &_go_drawBox36, &_go_drawBox37, &_go_drawBox38, &_go_drawBox39,
                          &_go_drawBox40, &_go_drawBox41, &_go_drawBox42, &_go_drawBox43,
                          &_go_drawBox44, &_go_drawBox45, &_go_drawBox46, &_go_drawBox47,
                          &_go_drawBox48, &_go_drawBox49, &_go_drawBox50, &_go_drawBox51,
                          &_go_drawBox52, &_go_drawBox53, &_go_drawBox54, &_go_drawBox55,
                          &_go_drawBox56};

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

void go_fltk_get_color(unsigned int col, unsigned char *r, unsigned char *g, unsigned char *b) {
  Fl::get_color((Fl_Color)col, *r, *g, *b);
}

const char *go_fltk_get_font(int font) {
  return Fl::get_font(font);
}

const char *go_fltk_get_font_name(int font, int *attributes) {
  return Fl::get_font_name(font, attributes);
}    

void go_fltk_set_font(Fl_Font font, const char* family) {
  Fl::set_font(font, family);
}

void go_fltk_set_font2(Fl_Font font, Fl_Font font2) {
  Fl::set_font(font, font2);
}

int go_fltk_set_fonts(const char *xstarname) {
  return (int)Fl::set_fonts(xstarname);
}  

unsigned go_fltk_get_colorindex(unsigned int col) {
  return Fl::get_color((Fl_Color)col);
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
int go_fltk_wait() {
  return Fl::wait();
}
int go_fltk_wait_timed(double t) {
  return Fl::wait(t);
}
int go_fltk_check() {
  return Fl::check();
}
void timeout_handler(void *data) {
  _go_timeoutHandler(uintptr_t(data));
}

void go_fltk_add_timeout(double t, uintptr_t id) {
  Fl::add_timeout(t, timeout_handler, (void*)id);
}

void go_fltk_repeat_timeout(double t, uintptr_t id) {
  Fl::repeat_timeout(t, timeout_handler, (void*)id);
}

void go_fltk_copy(const char* data, int len, int destination) {
  Fl::copy(data, len, destination);
}
void go_fltk_dnd() {
  Fl::dnd();
}

int go_fltk_screen_num(int x, int y) {
  return Fl::screen_num(x, y);
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

int go_fltk_scrollbar_size() {
  return Fl::scrollbar_size();
}
void go_fltk_set_scrollbar_size(int size) {
  Fl::scrollbar_size(size);
}

void go_fltk_set_menu_linespacing(int size) {
	Fl::menu_linespacing(size);
}
int go_fltk_menu_linespacing() {
	return Fl::menu_linespacing();
}

int go_fltk_test_shortcut(int shortcut) {
	return Fl::test_shortcut((unsigned int)shortcut);
}
