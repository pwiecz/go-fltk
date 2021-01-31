#include "widget.h"

#include <cstdint>

#include <Fl/Fl_Widget.H>

#include "enumerations.h"


void go_fltk_Widget_set_box(Fl_Widget *w, int box) {
  w->box((Fl_Boxtype)box);
}
void go_fltk_Widget_set_callback(Fl_Widget *w, void* id) {
  w->callback(callback_handler, id);
}
void go_fltk_Widget_when(Fl_Widget* w, int when) {
  w->when(when);
}
void go_fltk_Widget_set_labelfont(Fl_Widget *w, int font) {
  w->labelfont((Fl_Font)font);
}
void go_fltk_Widget_set_labelsize(Fl_Widget *w, int size) {
  w->labelsize(size);
}
void go_fltk_Widget_set_labeltype(Fl_Widget *w, int type) {
  w->labeltype((Fl_Labeltype)type);
}
int go_fltk_Widget_x(Fl_Widget *w) {
  return w->x();
}
int go_fltk_Widget_y(Fl_Widget *w) {
  return w->y();
}
int go_fltk_Widget_w(Fl_Widget *w) {
  return w->w();
}
int go_fltk_Widget_h(Fl_Widget *w) {
  return w->h();
}
void go_fltk_Widget_set_align(Fl_Widget *w, unsigned int align) {
  w->align((Fl_Align)align);
}
void go_fltk_Widget_measure_label(Fl_Widget *w, int *ww, int *hh) {
  w->measure_label(*ww, *hh);
}
void go_fltk_Widget_set_position(Fl_Widget *w, int x, int y) {
  w->resize(x, y, w->w(), w->h());
}
void go_fltk_Widget_resize(Fl_Widget* w, int x, int y, int width, int height) {
  w->resize(x, y, width, height);
}
void go_fltk_Widget_redraw(Fl_Widget *w) { 
  w->redraw();
}
void go_fltk_Widget_deactivate(Fl_Widget* w) {
  w->deactivate();
}
void go_fltk_Widget_activate(Fl_Widget* w) {
  w->activate();
}
void go_fltk_Widget_set_type(Fl_Widget* w, unsigned char type) {
  w->type(type);
}
void go_fltk_Widget_show(Fl_Widget* w) {
  w->show();
}
void go_fltk_Widget_hide(Fl_Widget* w) {
  w->hide();
}
unsigned int go_fltk_Widget_selection_color(Fl_Widget* w) {
  return w->selection_color();
}
void go_fltk_Widget_set_selection_color(Fl_Widget* w, unsigned int color) {
  w->selection_color(color);
}
void go_fltk_Widget_set_color(Fl_Widget* w, unsigned int color) {
  w->color(color);
}
void go_fltk_Widget_set_label(Fl_Widget* w, const char* label) {
  w->copy_label(label);
}
