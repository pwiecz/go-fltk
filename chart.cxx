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

void go_fltk_Chart_clear(Fl_Chart* c) {
    c->clear();
}

void go_fltk_Chart_add(Fl_Chart* c, double val, const char *str = 0, unsigned col = 0) {
    c->add(val, str, col);
}

void go_fltk_Chart_insert(Fl_Chart* c, int ind, double val, const char *str = 0, unsigned col = 0) {
    c->insert(ind, val, str, col);
}

void go_fltk_Chart_replace(Fl_Chart* c, int ind, double val, const char *str = 0, unsigned col = 0){
    c->replace(ind, val, str, col);
}

void go_fltk_Chart_bounds(Fl_Chart* c, double *a,double *b) {
    c->bounds(a,b);
}

void go_fltk_Chart_set_bounds(Fl_Chart* c, double a,double b){
    c->bounds(a,b);
}

int go_fltk_Chart_size(Fl_Chart* c) {
    return c->size();
}

void go_fltk_Chart_set_size(Fl_Chart* c, int W, int H) {
    c->size(W, H);
}

int go_fltk_Chart_maxsize(Fl_Chart* c){
    return c->maxsize();
}

void go_fltk_Chart_set_maxsize(Fl_Chart* c, int m){
    c->maxsize(m);
}

int go_fltk_Chart_textfont(Fl_Chart* c) {
    return c->textfont();
}

void go_fltk_Chart_set_textfont(Fl_Chart* c, int font_s){
    c->textfont(font_s);
}

int go_fltk_Chart_textsize(Fl_Chart* c){
    return c->textsize();
}

void go_fltk_Chart_set_textsize(Fl_Chart* c, int s){
    c->textsize(s);
}

unsigned int go_fltk_Chart_textcolor(Fl_Chart* c){
    return c->textcolor();
}

void go_fltk_Chart_set_textcolor(Fl_Chart* c, unsigned int color_n){
    c->textcolor(color_n);
}

uchar go_fltk_Chart_autosize(Fl_Chart* c){
    return c->autosize();
}

void go_fltk_Chart_set_autosize(Fl_Chart* c, uchar n){
    c->autosize(n);
}
