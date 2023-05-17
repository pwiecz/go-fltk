#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GOffscreen GOffscreen;
  
  extern void go_fltk_color(unsigned int color);
  extern void go_fltk_set_draw_font(int font, int size);
  extern int go_fltk_draw_font();
  extern int go_fltk_draw_font_size();   
  extern void go_fltk_draw(const char* text, int x, int y, int w, int h, unsigned int align);
  extern void go_fltk_draw_box(int boxType, int x, int y, int w, int h, unsigned int color);
  extern void go_fltk_push_clip(int x, int y, int w, int h);
  extern void go_fltk_push_no_clip(void);
  extern void go_fltk_pop_clip(void);
  extern int go_fltk_not_clipped(int x, int y, int w, int h);
  extern int go_fltk_clip_box(int x, int y, int w, int h, int *X, int *Y, int *W, int *H);
  extern void go_fltk_restore_clip(void);
  extern void go_fltk_set_clip_region(void *r);
  extern void *go_fltk_clip_region(void);
  extern void go_fltk_point(int x, int y);
  extern void go_fltk_line_style(int style, int width, char *dashes);
  extern void go_fltk_rect(int x, int y, int w, int h);
  extern void go_fltk_focus_rect(int x, int y, int w, int h);
  extern void go_fltk_rect_with_color(int x, int y, int w, int h, unsigned int c);
  extern void go_fltk_rectf(int x, int y, int w, int h);
  extern void go_fltk_rectf_with_color(int x, int y, int w, int h, unsigned int c);
  extern void go_fltk_rectf_with_rgb(int x, int y, int w, int h, unsigned char r, unsigned char g,
                          unsigned char b);
  extern void go_fltk_draw_arrow(int x, int y, int w, int h, int arr, int orient, unsigned int color);
  extern void go_fltk_line(int x, int y, int x1, int y1);
  extern void go_fltk_line2(int x, int y, int x1, int y1, int x2, int y2);
  extern void go_fltk_loop(int x, int y, int x1, int y1, int x2, int y2);
  extern void go_fltk_loop2(int x, int y, int x1, int y1, int x2, int y2, int x3, int y3);
  extern void go_fltk_polygon(int x, int y, int x1, int y1, int x2, int y2);
  extern void go_fltk_polygon2(int x, int y, int x1, int y1, int x2, int y2, int x3, int y3);
  extern void go_fltk_xyline(int x, int y, int x1);
  extern void go_fltk_xyline2(int x, int y, int x1, int y2);
  extern void go_fltk_xyline3(int x, int y, int x1, int y2, int x3);
  extern void go_fltk_yxline(int x, int y, int y1);
  extern void go_fltk_yxline2(int x, int y, int y1, int x2);
  extern void go_fltk_yxline3(int x, int y, int y1, int x2, int y3);
  extern void go_fltk_arc(int x, int y, int w, int h, double a1, double a2);
  extern void go_fltk_pie(int x, int y, int w, int h, double a1, double a2);
  extern void go_fltk_push_matrix(void);
  extern void go_fltk_pop_matrix(void);
  extern void go_fltk_scale(double x, double y);
  extern void go_fltk_scale2(double x);
  extern void go_fltk_translate(double x, double y);
  extern void go_fltk_rotate(double d);
  extern void go_fltk_mult_matrix(double a, double b, double c, double d, double x, double y);
  extern void go_fltk_begin_points(void);
  extern void go_fltk_begin_line(void);
  extern void go_fltk_begin_loop(void);
  extern void go_fltk_begin_polygon(void);
  extern void go_fltk_vertex(double x, double y);
  extern void go_fltk_curve(double X0, double Y0, double X1, double Y1, double X2, double Y2, double X3,
                  double Y3);
  extern void go_fltk_arc2(double x, double y, double r, double start, double end);
  extern void go_fltk_circle(double x, double y, double r);
  extern void go_fltk_end_points(void);
  extern void go_fltk_end_line(void);
  extern void go_fltk_end_loop(void);
  extern void go_fltk_end_polygon(void);
  extern void go_fltk_begin_complex_polygon(void);
  extern void go_fltk_gap(void);
  extern void go_fltk_end_complex_polygon(void);
  extern double go_fltk_transform_x(double x, double y);
  extern double go_fltk_transform_y(double x, double y);
  extern double go_fltk_transform_dx(double x, double y);
  extern double go_fltk_transform_dy(double x, double y);
  extern void go_fltk_transformed_vertex(double xf, double yf);
  extern void go_fltk_end_offscreen(void);
  extern int go_fltk_font(void);
  extern int go_fltk_size(void);
  extern int go_fltk_height(void);
  extern int go_fltk_set_height(int font, int size);
  extern int go_fltk_descent(void);
  extern double go_fltk_width(const char *txt);
  extern double go_fltk_width2(const char *txt, int n);
  extern double go_fltk_width3(unsigned int c);
  extern void go_fltk_text_extents(const char *, int *dx, int *dy, int *w, int *h);
  extern void go_fltk_text_extents2(const char *t, int n, int *dx, int *dy, int *w, int *h);
  extern const char *go_fltk_latin1_to_local(const char *t, int n);
  extern const char *go_fltk_local_to_latin1(const char *t, int n);
  extern const char *go_fltk_mac_roman_to_local(const char *t, int n);
  extern const char *go_fltk_local_to_mac_roman(const char *t, int n);
  extern void go_fltk_draw2(int angle, const char *str, int x, int y);
  extern void go_fltk_draw3(const char *str, int n, int x, int y);
  extern void go_fltk_draw4(int angle, const char *str, int n, int x, int y);
  extern void go_fltk_rtl_draw(const char *str, int n, int x, int y);
  extern void go_fltk_measure(const char *str, int *x, int *y, int draw_symbols);
  extern void go_fltk_draw5(const char *str, int x, int y, int w, int h, int align, void **img, int draw_symbols);
  extern void go_fltk_frame(const char *s, int x, int y, int w, int h);
  extern void go_fltk_frame2(const char *s, int x, int y, int w, int h);
  extern void go_fltk_draw_image(const unsigned char *buf, int X, int Y, int W, int H, int D, int L);
  extern void go_fltk_draw_image_mono(const unsigned char *buf, int X, int Y, int W, int H, int D, int L);
  extern char go_fltk_can_do_alpha_blending(void);
  extern unsigned char *go_fltk_read_image(unsigned char *p, int X, int Y, int W, int H, int alpha);
  extern unsigned char *go_fltk_capture_window_part(void *win, int x, int y, int w, int h);
  extern int go_fltk_draw_pixmap(const char *const *data, int x, int y, int bg);
  extern int go_fltk_draw_pixmap2(/*const*/ char *const *data, int x, int y, int bg);
  extern int go_fltk_measure_pixmap(/*const*/ char *const *data, int *w, int *h);
  extern int go_fltk_measure_pixmap2(const char *const *cdata, int *w, int *h);
  extern const char *go_fltk_shortcut_label(unsigned int shortcut);
  extern const char *go_fltk_shortcut_label2(unsigned int shortcut, const char **eom);
  extern unsigned int go_fltk_old_shortcut(const char *s);
  extern void go_fltk_overlay_rect(int x, int y, int w, int h);
  extern void go_fltk_overlay_clear(void);
  extern void go_fltk_set_cursor(int cursor);
  extern void go_fltk_set_cursor2(int cursor, int fg, int bg);
  extern const char *go_fltk_expand_text(const char *from, char *buf, int maxbuf, double maxw, int *n,
                              double *width, int wrap, int draw_symbols);
  extern void go_fltk_set_status(int X, int Y, int W, int H);
  extern void go_fltk_set_spot(int font, int size, int X, int Y, int W, int H, void *win);
  extern void go_fltk_reset_spot(void);
  extern void go_fltk_copy_offscreen(int x, int y, int w, int h, GOffscreen *pixmap, int srcx, int srcy);
  extern GOffscreen *go_fltk_create_offscreen(int w, int h);
  extern void go_fltk_begin_offscreen(GOffscreen *b);
  extern void go_fltk_end_offscreen(void);
  extern void go_fltk_delete_offscreen(GOffscreen *bitmap);
  extern void go_fltk_rescale_offscreen(GOffscreen **ctx);
  extern void go_fltk_draw_text2(const char *str, int x, int y, int w, int h, int align);
  extern void go_fltk_draw_check(int x, int y, int w, int h, unsigned int col);

#ifdef __cplusplus
}
#endif
