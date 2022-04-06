#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GChart GChart;

  extern GChart* go_fltk_new_Chart(int x, int y, int w, int h, const char *label);
  extern void go_fltk_Chart_clear(GChart* c);
  extern void go_fltk_Chart_add(GChart* c, double val, const char *str, unsigned col);
  extern void go_fltk_Chart_insert(GChart* c, int ind, double val, const char *str, unsigned col);
  extern void go_fltk_Chart_replace(GChart* c, int ind, double val, const char *str, unsigned col);
  extern void go_fltk_Chart_bounds(GChart* c, double *a,double *b);
  extern void go_fltk_Chart_set_bounds(GChart* c, double a,double b);
  extern int go_fltk_Chart_size(GChart* c);
  extern void go_fltk_Chart_set_size(GChart* c, int W, int H);
  extern int go_fltk_Chart_maxsize(GChart* c);
  extern void go_fltk_Chart_set_maxsize(GChart* c, int m);
  extern int go_fltk_Chart_textfont(GChart* c);
  extern void go_fltk_Chart_set_textfont(GChart* c, int font_s);
  extern int go_fltk_Chart_textsize(GChart* c);
  extern void go_fltk_Chart_set_textsize(GChart* c, int s);
  extern unsigned int go_fltk_Chart_textcolor(GChart* c);
  extern void go_fltk_Chart_set_textcolor(GChart* c, unsigned int color_n);
  extern unsigned char go_fltk_Chart_autosize(GChart* c);
  extern void go_fltk_Chart_set_autosize(GChart* c, unsigned char n);

#ifdef __cplusplus
}
#endif
