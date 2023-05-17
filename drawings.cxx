#include "drawings.h"

#include <FL/fl_draw.H>
#include <FL/platform.H>
#include <FL/Enumerations.H>


void go_fltk_color(unsigned int color) {
  fl_color((Fl_Color)color);
}
void go_fltk_set_draw_font(int font, int size) {
  fl_font((Fl_Font)font, size);
}
int go_fltk_draw_font() {
	return fl_font();
}
int go_fltk_draw_font_size() {
	return fl_size();
}
void go_fltk_draw(const char* text, int x, int y, int w, int h, unsigned int align) {
  fl_draw(text, x, y, w, h, (Fl_Align)align);
}
void go_fltk_draw_box(int boxType, int x, int y, int w, int h, unsigned int color) {
  fl_draw_box(Fl_Boxtype(boxType), x, y, w, h, (Fl_Color)color);
}

void go_fltk_push_clip(int x, int y, int w, int h) {
    fl_push_clip(x, y, w, h);
}

void go_fltk_push_no_clip(void) {
    fl_push_no_clip();
}

void go_fltk_pop_clip(void) {
    fl_pop_clip();
}

int go_fltk_not_clipped(int x, int y, int w, int h) {
    return fl_not_clipped(x, y, w, h);
}

int go_fltk_clip_box(int x, int y, int w, int h, int *X, int *Y, int *W, int *H) {
    return fl_clip_box(x, y, w, h, *X, *Y, *W, *H);
}

void go_fltk_restore_clip(void) {
    fl_restore_clip();
}

void go_fltk_set_clip_region(void *r) {
    fl_clip_region((Fl_Region)r);
}

void *go_fltk_clip_region(void) {
    return (void *)fl_clip_region();
}

void go_fltk_point(int x, int y) {
    fl_point(x, y);
}

void go_fltk_line_style(int style, int width, char *dashes) {
    fl_line_style(style, width, dashes);
}

void go_fltk_rect(int x, int y, int w, int h) {
    fl_rect(x, y, w, h);
}

void go_fltk_focus_rect(int x, int y, int w, int h) {
    fl_focus_rect(x, y, w, h);
}

void go_fltk_draw_arrow(int x, int y, int w, int h, int arr, int orient, unsigned int color) {
	Fl_Rect r{x, y, w, h};
	fl_draw_arrow(r, (Fl_Arrow_Type)arr, (Fl_Orientation)orient, color);
}

void go_fltk_rect_with_color(int x, int y, int w, int h, unsigned int c) {
    fl_rect(x, y, w, h, c);
}

void go_fltk_rectf(int x, int y, int w, int h) {
    fl_rectf(x, y, w, h);
}

void go_fltk_rectf_with_color(int x, int y, int w, int h, unsigned int c) {
    fl_rectf(x, y, w, h, c);
}

void go_fltk_rectf_with_rgb(int x, int y, int w, int h, unsigned char r, unsigned char g,
                       unsigned char b) {
    fl_rectf(x, y, w, h, r, g, b);
}

void go_fltk_line(int x, int y, int x1, int y1) {
    fl_line(x, y, x1, y1);
}

void go_fltk_line2(int x, int y, int x1, int y1, int x2, int y2) {
    fl_line(x, y, x1, y1, x2, y2);
}

void go_fltk_loop(int x, int y, int x1, int y1, int x2, int y2) {
    fl_loop(x, y, x1, y1, x2, y2);
}

void go_fltk_loop2(int x, int y, int x1, int y1, int x2, int y2, int x3, int y3) {
    fl_loop(x, y, x1, y1, x2, y2, x3, y3);
}

void go_fltk_polygon(int x, int y, int x1, int y1, int x2, int y2) {
    fl_polygon(x, y, x1, y1, x2, y2);
}

void go_fltk_polygon2(int x, int y, int x1, int y1, int x2, int y2, int x3, int y3) {
    fl_polygon(x, y, x1, y1, x2, y2, x3, y3);
}

void go_fltk_xyline(int x, int y, int x1) {
    fl_xyline(x, y, x1);
}

void go_fltk_xyline2(int x, int y, int x1, int y2) {
    fl_xyline(x, y, x1, y2);
}

void go_fltk_xyline3(int x, int y, int x1, int y2, int x3) {
    fl_xyline(x, y, x1, y2, x3);
}

void go_fltk_yxline(int x, int y, int y1) {
    fl_yxline(x, y, y1);
}

void go_fltk_yxline2(int x, int y, int y1, int x2) {
    fl_yxline(x, y, y1, x2);
}

void go_fltk_yxline3(int x, int y, int y1, int x2, int y3) {
    fl_yxline(x, y, y1, x2, y3);
}

void go_fltk_arc(int x, int y, int w, int h, double a1, double a2) {
    fl_arc(x, y, w, h, a1, a2);
}

void go_fltk_pie(int x, int y, int w, int h, double a1, double a2) {
    fl_pie(x, y, w, h, a1, a2);
}

void go_fltk_push_matrix(void) {
    fl_push_matrix();
}

void go_fltk_pop_matrix(void) {
    fl_pop_matrix();
}

void go_fltk_scale(double x, double y) {
    fl_scale(x, y);
}

void go_fltk_scale2(double x) {
    fl_scale(x);
}

void go_fltk_translate(double x, double y) {
    fl_translate(x, y);
}

void go_fltk_rotate(double d) {
    fl_rotate(d);
}

void go_fltk_mult_matrix(double a, double b, double c, double d, double x, double y) {
    fl_mult_matrix(a, b, c, d, x, y);
}

void go_fltk_begin_points(void) {
    fl_begin_points();
}

void go_fltk_begin_line(void) {
    fl_begin_line();
}

void go_fltk_begin_loop(void) {
    fl_begin_loop();
}

void go_fltk_begin_polygon(void) {
    fl_begin_polygon();
}

void go_fltk_vertex(double x, double y) {
    fl_vertex(x, y);
}

void go_fltk_curve(double X0, double Y0, double X1, double Y1, double X2, double Y2, double X3,
              double Y3) {
    fl_curve(X0, Y0, X1, Y1, X2, Y2, X3, Y3);
}

void go_fltk_arc2(double x, double y, double r, double start, double end) {
    fl_arc(x, y, r, start, end);
}

void go_fltk_circle(double x, double y, double r) {
    fl_circle(x, y, r);
}

void go_fltk_end_points(void) {
    fl_end_points();
}

void go_fltk_end_line(void) {
    fl_end_line();
}

void go_fltk_end_loop(void) {
    fl_end_loop();
}

void go_fltk_end_polygon(void) {
    fl_end_polygon();
}

void go_fltk_begin_complex_polygon(void) {
    fl_begin_complex_polygon();
}

void go_fltk_gap(void) {
    fl_gap();
}

void go_fltk_end_complex_polygon(void) {
    fl_end_complex_polygon();
}

double go_fltk_transform_x(double x, double y) {
    return fl_transform_x(x, y);
}

double go_fltk_transform_y(double x, double y) {
    return fl_transform_y(x, y);
}

double go_fltk_transform_dx(double x, double y) {
    return fl_transform_dx(x, y);
}

double go_fltk_transform_dy(double x, double y) {
    return fl_transform_dy(x, y);
}

void go_fltk_transformed_vertex(double xf, double yf) {
    fl_transformed_vertex(xf, yf);
}

int go_fltk_font(void) {
    return fl_font();
}

int go_fltk_size(void) {
    return fl_size();
}

int go_fltk_height(void) {
    return fl_height();
}

int go_fltk_set_height(int font, int size) {
    return fl_height(font, size);
}

int go_fltk_descent(void) {
    return fl_descent();
}

double go_fltk_width(const char *txt) {
    return fl_width(txt);
}

double go_fltk_width2(const char *txt, int n) {
    return fl_width(txt, n);
}

double go_fltk_width3(unsigned int c) {
    return fl_width(c);
}

void go_fltk_text_extents(const char *txt, int *dx, int *dy, int *w, int *h) {
    return fl_text_extents(txt, *dx, *dy, *w, *h);
}

void go_fltk_text_extents2(const char *t, int n, int *dx, int *dy, int *w, int *h) {
    return fl_text_extents(t, n, *dx, *dy, *w, *h);
}

const char *go_fltk_latin1_to_local(const char *t, int n) {
    return fl_latin1_to_local(t, n);
}

const char *go_fltk_local_to_latin1(const char *t, int n) {
    return fl_local_to_latin1(t, n);
}

const char *go_fltk_mac_roman_to_local(const char *t, int n) {
    return fl_mac_roman_to_local(t, n);
}

const char *go_fltk_local_to_mac_roman(const char *t, int n) {
    return fl_local_to_mac_roman(t, n);
}

void go_fltk_draw2(int angle, const char *str, int x, int y) {
    fl_draw(angle, str, x, y);
}

void go_fltk_draw3(const char *str, int n, int x, int y) {
    fl_draw(str, n, x, y);
}

void go_fltk_draw4(int angle, const char *str, int n, int x, int y) {
    fl_draw(angle, str, n, x, y);
}

void go_fltk_rtl_draw(const char *str, int n, int x, int y) {
    fl_rtl_draw(str, n, x, y);
}

void go_fltk_measure(const char *str, int *x, int *y, int draw_symbols) {
    fl_measure(str, *x, *y, draw_symbols);
}

void go_fltk_draw5(const char *str, int x, int y, int w, int h, int align, void **img,
              int draw_symbols) {
    fl_draw(str, x, y, w, h, align, (Fl_Image *)*img, draw_symbols);
}

void go_fltk_frame(const char *s, int x, int y, int w, int h) {
    fl_frame(s, x, y, w, h);
}

void go_fltk_frame2(const char *s, int x, int y, int w, int h) {
    fl_frame2(s, x, y, w, h);
}

void go_fltk_draw_image(const unsigned char *buf, int X, int Y, int W, int H, int D, int L) {
    fl_draw_image(buf, X, Y, W, H, D, L);
}

void go_fltk_draw_image_mono(const unsigned char *buf, int X, int Y, int W, int H, int D, int L) {
    fl_draw_image_mono(buf, X, Y, W, H, D, L);
}

char go_fltk_can_do_alpha_blending(void) {
    return fl_can_do_alpha_blending();
}

unsigned char *go_fltk_read_image(unsigned char *p, int X, int Y, int W, int H, int alpha) {
    return fl_read_image(p, X, Y, W, H, alpha);
}

unsigned char *go_fltk_capture_window_part(void *win, int x, int y, int w, int h) {
    Fl_RGB_Image *tmp = fl_capture_window((Fl_Window *)win, x, y, w, h);
    return (unsigned char *)tmp->data();
}

int go_fltk_draw_pixmap(const char *const *data, int x, int y, int bg) {
    return fl_draw_pixmap(data, x, y, bg);
}

int go_fltk_draw_pixmap2(/*const*/ char *const *data, int x, int y, int bg) {
    return fl_draw_pixmap(data, x, y, bg);
}

int go_fltk_measure_pixmap(/*const*/ char *const *data, int *w, int *h) {
    return fl_measure_pixmap(data, *w, *h);
}

int go_fltk_measure_pixmap2(const char *const *cdata, int *w, int *h) {
    return fl_measure_pixmap(cdata, *w, *h);
}

const char *go_fltk_shortcut_label(unsigned int shortcut) {
    return fl_shortcut_label(shortcut);
}

const char *go_fltk_shortcut_label2(unsigned int shortcut, const char **eom) {
    return fl_shortcut_label(shortcut, eom);
}

unsigned int go_fltk_old_shortcut(const char *s) {
    return fl_old_shortcut(s);
}

void go_fltk_overlay_rect(int x, int y, int w, int h) {
    return fl_overlay_rect(x, y, w, h);
}

void go_fltk_overlay_clear(void) {
    return fl_overlay_clear();
}

void go_fltk_set_cursor(int cursor) {
    return fl_cursor((Fl_Cursor)cursor);
}

void go_fltk_set_cursor2(int cursor, int fg, int bg) {
    return fl_cursor((Fl_Cursor)cursor, fg, bg);
}

const char *go_fltk_expand_text(const char *from, char *buf, int maxbuf, double maxw, int *n,
                           double *width, int wrap, int draw_symbols) {
    return fl_expand_text(from, buf, maxbuf, maxw, *n, *width, wrap, draw_symbols);
}

void go_fltk_set_status(int X, int Y, int W, int H) {
    fl_set_status(X, Y, W, H);
}

void go_fltk_set_spot(int font, int size, int X, int Y, int W, int H, void *win) {
    fl_set_spot(font, size, X, Y, W, H, (Fl_Window *)win);
}

void go_fltk_reset_spot(void) {
    fl_reset_spot();
}

void go_fltk_copy_offscreen(int x, int y, int w, int h, GOffscreen *pixmap, int srcx, int srcy) {
    fl_copy_offscreen(x, y, w, h, (Fl_Offscreen)pixmap, srcx, srcy);
}

GOffscreen *go_fltk_create_offscreen(int w, int h) {
    fl_open_display();
    return (GOffscreen *)fl_create_offscreen(w, h);
}

void go_fltk_begin_offscreen(GOffscreen *b) {
    fl_begin_offscreen((Fl_Offscreen)b);
}

void go_fltk_end_offscreen(void) {
    fl_end_offscreen();
}

void go_fltk_delete_offscreen(GOffscreen *bitmap) {
    fl_delete_offscreen((Fl_Offscreen)bitmap);
}

void go_fltk_rescale_offscreen(GOffscreen **ctx) {
    fl_rescale_offscreen(*(Fl_Offscreen *)ctx);
}

void go_fltk_draw_text2(const char *str, int x, int y, int w, int h, int align) {
    fl_draw(str, x, y, w, h, (Fl_Align)align, 0, 1);
}

void go_fltk_draw_check(int x, int y, int w, int h, unsigned int col) {
    fl_draw_check(Fl_Rect(x, y, w, h), (Fl_Color)col);
}
