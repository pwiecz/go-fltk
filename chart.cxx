#include "chart.h"

#include <FL/Fl_Chart.H>

#include "event_handler.h"


class GChart : public EventHandler<Fl_Chart> {
public:
  GChart(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Chart>(x, y, w, h, label) {}
};

GChart* go_fltk_new_Chart(int x, int y, int w, int h, const char *label) {
    return new GChart(x, y, w, h, label);
}

void go_fltk_Chart_clear(GChart* c) {
    c->clear();
}

void go_fltk_Chart_add(GChart* c, double val, const char *str = 0, unsigned col = 0) {
    c->add(val, str, col);
}

void go_fltk_Chart_insert(GChart* c, int ind, double val, const char *str = 0, unsigned col = 0) {
    c->insert(ind, val, str, col);
}

void go_fltk_Chart_replace(GChart* c, int ind, double val, const char *str = 0, unsigned col = 0){
    c->replace(ind, val, str, col);
}

void go_fltk_Chart_bounds(GChart* c, double *a,double *b) {
    c->bounds(a,b);
}

void go_fltk_Chart_set_bounds(GChart* c, double a,double b){
    c->bounds(a,b);
}

int go_fltk_Chart_size(GChart* c) {
    return c->size();
}

void go_fltk_Chart_set_size(GChart* c, int W, int H) {
    c->size(W, H);
}

int go_fltk_Chart_maxsize(GChart* c){
    return c->maxsize();
}

void go_fltk_Chart_set_maxsize(GChart* c, int m){
    c->maxsize(m);
}

int go_fltk_Chart_textfont(GChart* c) {
    return c->textfont();
}

void go_fltk_Chart_set_textfont(GChart* c, int font_s){
    c->textfont(font_s);
}

int go_fltk_Chart_textsize(GChart* c){
    return c->textsize();
}

void go_fltk_Chart_set_textsize(GChart* c, int s){
    c->textsize(s);
}

unsigned int go_fltk_Chart_textcolor(GChart* c){
    return c->textcolor();
}

void go_fltk_Chart_set_textcolor(GChart* c, unsigned int color_n){
    c->textcolor(color_n);
}

uchar go_fltk_Chart_autosize(GChart* c){
    return c->autosize();
}

void go_fltk_Chart_set_autosize(GChart* c, uchar n){
    c->autosize(n);
}
