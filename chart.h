#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Chart Fl_Chart;
  typedef struct GChart GChart;

  extern GChart* go_fltk_new_Chart(int x, int y, int w, int h, const char *label);
  extern void go_fltk_Chart_clear(Fl_Chart* c);
  extern void go_fltk_Chart_add(Fl_Chart* c, double val, const char *str, unsigned col);
  extern void go_fltk_Chart_insert(Fl_Chart* c, int ind, double val, const char *str, unsigned col);
  extern void go_fltk_Chart_replace(Fl_Chart* c, int ind, double val, const char *str, unsigned col);
  extern void go_fltk_Chart_bounds(Fl_Chart* c, double *a,double *b);
  extern void go_fltk_Chart_set_bounds(Fl_Chart* c, double a,double b);
  extern int go_fltk_Chart_size(Fl_Chart* c);
  extern void go_fltk_Chart_set_size(Fl_Chart* c, int W, int H);
  extern int go_fltk_Chart_maxsize(Fl_Chart* c);
  extern void go_fltk_Chart_set_maxsize(Fl_Chart* c, int m);
  extern int go_fltk_Chart_textfont(Fl_Chart* c);
  extern void go_fltk_Chart_set_textfont(Fl_Chart* c, int font_s);
  extern int go_fltk_Chart_textsize(Fl_Chart* c);
  extern void go_fltk_Chart_set_textsize(Fl_Chart* c, int s);
  extern unsigned int go_fltk_Chart_textcolor(Fl_Chart* c);
  extern void go_fltk_Chart_set_textcolor(Fl_Chart* c, unsigned int color_n);
  extern unsigned char go_fltk_Chart_autosize(Fl_Chart* c);
  extern void go_fltk_Chart_set_autosize(Fl_Chart* c, unsigned char n);

#ifdef __cplusplus
}
#endif
