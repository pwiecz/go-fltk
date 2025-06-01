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

#define GEN_DRAW_BOX(n) \
void _go_drawBox##n (int x, int y, int w, int h, unsigned int c) { \
    _go_drawBox(n, x, y, w, h, c); \
}

#define DRAW_BOX_LIST \
    X(0) X(1) X(2) X(3) X(4) X(5) X(6) X(7) X(8) X(9) \
    X(10) X(11) X(12) X(13) X(14) X(15) X(16) X(17) X(18) X(19) \
    X(20) X(21) X(22) X(23) X(24) X(25) X(26) X(27) X(28) X(29) \
    X(30) X(31) X(32) X(33) X(34) X(35) X(36) X(37) X(38) X(39) \
    X(40) X(41) X(42) X(43) X(44) X(45) X(46) X(47) X(48) X(49) \
    X(50) X(51) X(52) X(53) X(54) X(55) X(56) X(57) X(58) X(59) \
    X(60) X(61) X(62) X(63) X(64) X(65) X(66) X(67) X(68) X(69) \
    X(70) X(71) X(72) X(73) X(74) X(75) X(76) X(77) X(78) X(79) \
    X(80) X(81) X(82) X(83) X(84) X(85) X(86) X(87) X(88) X(89) \
    X(90) X(91) X(92) X(93) X(94) X(95) X(96) X(97) X(98) X(99) \
    X(100) X(101) X(102) X(103) X(104) X(105) X(106) X(107) X(108) X(109) \
    X(110) X(111) X(112) X(113) X(114) X(115) X(116) X(117) X(118) X(119) \
    X(120) X(121) X(122) X(123) X(124) X(125) X(126) X(127) X(128) X(129) \
    X(130) X(131) X(132) X(133) X(134) X(135) X(136) X(137) X(138) X(139) \
    X(140) X(141) X(142) X(143) X(144) X(145) X(146) X(147) X(148) X(149) \
    X(150) X(151) X(152) X(153) X(154) X(155) X(156) X(157) X(158) X(159) \
    X(160) X(161) X(162) X(163) X(164) X(165) X(166) X(167) X(168) X(169) \
    X(170) X(171) X(172) X(173) X(174) X(175) X(176) X(177) X(178) X(179) \
    X(180) X(181) X(182) X(183) X(184) X(185) X(186) X(187) X(188) X(189) \
    X(190) X(191) X(192) X(193) X(194) X(195) X(196) X(197) X(198) X(199) \
    X(200) X(201) X(202) X(203) X(204) X(205) X(206) X(207) X(208) X(209) \
    X(210) X(211) X(212) X(213) X(214) X(215) X(216) X(217) X(218) X(219) \
    X(220) X(221) X(222) X(223) X(224) X(225) X(226) X(227) X(228) X(229) \
    X(230) X(231) X(232) X(233) X(234) X(235) X(236) X(237) X(238) X(239) \
    X(240) X(241) X(242) X(243) X(244) X(245) X(246) X(247) X(248) X(249) \
    X(250) X(251) X(252) X(253) X(254) X(255)

#define X(n) GEN_DRAW_BOX(n)
DRAW_BOX_LIST
#undef X

typedef void (*f_array)(int, int, int, int, unsigned int);

// Helper macro to turn a number into the address of _go_drawBox<number>
#define F_PTR(n) &_go_drawBox##n

// Macro that expands to the list of function pointers, 10 per line.
#define DRAW_BOX_PTR_LIST \
    F_PTR(0),  F_PTR(1),  F_PTR(2),  F_PTR(3),  F_PTR(4),  F_PTR(5),  F_PTR(6),  F_PTR(7),  F_PTR(8),  F_PTR(9),  \
    F_PTR(10), F_PTR(11), F_PTR(12), F_PTR(13), F_PTR(14), F_PTR(15), F_PTR(16), F_PTR(17), F_PTR(18), F_PTR(19), \
    F_PTR(20), F_PTR(21), F_PTR(22), F_PTR(23), F_PTR(24), F_PTR(25), F_PTR(26), F_PTR(27), F_PTR(28), F_PTR(29), \
    F_PTR(30), F_PTR(31), F_PTR(32), F_PTR(33), F_PTR(34), F_PTR(35), F_PTR(36), F_PTR(37), F_PTR(38), F_PTR(39), \
    F_PTR(40), F_PTR(41), F_PTR(42), F_PTR(43), F_PTR(44), F_PTR(45), F_PTR(46), F_PTR(47), F_PTR(48), F_PTR(49), \
    F_PTR(50), F_PTR(51), F_PTR(52), F_PTR(53), F_PTR(54), F_PTR(55), F_PTR(56), F_PTR(57), F_PTR(58), F_PTR(59), \
    F_PTR(60), F_PTR(61), F_PTR(62), F_PTR(63), F_PTR(64), F_PTR(65), F_PTR(66), F_PTR(67), F_PTR(68), F_PTR(69), \
    F_PTR(70), F_PTR(71), F_PTR(72), F_PTR(73), F_PTR(74), F_PTR(75), F_PTR(76), F_PTR(77), F_PTR(78), F_PTR(79), \
    F_PTR(80), F_PTR(81), F_PTR(82), F_PTR(83), F_PTR(84), F_PTR(85), F_PTR(86), F_PTR(87), F_PTR(88), F_PTR(89), \
    F_PTR(90), F_PTR(91), F_PTR(92), F_PTR(93), F_PTR(94), F_PTR(95), F_PTR(96), F_PTR(97), F_PTR(98), F_PTR(99), \
    F_PTR(100), F_PTR(101), F_PTR(102), F_PTR(103), F_PTR(104), F_PTR(105), F_PTR(106), F_PTR(107), F_PTR(108), F_PTR(109), \
    F_PTR(110), F_PTR(111), F_PTR(112), F_PTR(113), F_PTR(114), F_PTR(115), F_PTR(116), F_PTR(117), F_PTR(118), F_PTR(119), \
    F_PTR(120), F_PTR(121), F_PTR(122), F_PTR(123), F_PTR(124), F_PTR(125), F_PTR(126), F_PTR(127), F_PTR(128), F_PTR(129), \
    F_PTR(130), F_PTR(131), F_PTR(132), F_PTR(133), F_PTR(134), F_PTR(135), F_PTR(136), F_PTR(137), F_PTR(138), F_PTR(139), \
    F_PTR(140), F_PTR(141), F_PTR(142), F_PTR(143), F_PTR(144), F_PTR(145), F_PTR(146), F_PTR(147), F_PTR(148), F_PTR(149), \
    F_PTR(150), F_PTR(151), F_PTR(152), F_PTR(153), F_PTR(154), F_PTR(155), F_PTR(156), F_PTR(157), F_PTR(158), F_PTR(159), \
    F_PTR(160), F_PTR(161), F_PTR(162), F_PTR(163), F_PTR(164), F_PTR(165), F_PTR(166), F_PTR(167), F_PTR(168), F_PTR(169), \
    F_PTR(170), F_PTR(171), F_PTR(172), F_PTR(173), F_PTR(174), F_PTR(175), F_PTR(176), F_PTR(177), F_PTR(178), F_PTR(179), \
    F_PTR(180), F_PTR(181), F_PTR(182), F_PTR(183), F_PTR(184), F_PTR(185), F_PTR(186), F_PTR(187), F_PTR(188), F_PTR(189), \
    F_PTR(190), F_PTR(191), F_PTR(192), F_PTR(193), F_PTR(194), F_PTR(195), F_PTR(196), F_PTR(197), F_PTR(198), F_PTR(199), \
    F_PTR(200), F_PTR(201), F_PTR(202), F_PTR(203), F_PTR(204), F_PTR(205), F_PTR(206), F_PTR(207), F_PTR(208), F_PTR(209), \
    F_PTR(210), F_PTR(211), F_PTR(212), F_PTR(213), F_PTR(214), F_PTR(215), F_PTR(216), F_PTR(217), F_PTR(218), F_PTR(219), \
    F_PTR(220), F_PTR(221), F_PTR(222), F_PTR(223), F_PTR(224), F_PTR(225), F_PTR(226), F_PTR(227), F_PTR(228), F_PTR(229), \
    F_PTR(230), F_PTR(231), F_PTR(232), F_PTR(233), F_PTR(234), F_PTR(235), F_PTR(236), F_PTR(237), F_PTR(238), F_PTR(239), \
    F_PTR(240), F_PTR(241), F_PTR(242), F_PTR(243), F_PTR(244), F_PTR(245), F_PTR(246), F_PTR(247), F_PTR(248), F_PTR(249), \
    F_PTR(250), F_PTR(251), F_PTR(252), F_PTR(253), F_PTR(254), F_PTR(255)

f_array currentBoxTypeCb[256] = { DRAW_BOX_PTR_LIST };

#undef F_PTR
#undef DRAW_BOX_PTR_LIST

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
