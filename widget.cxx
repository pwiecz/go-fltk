#include "widget.h"

#include <FL/Fl.H>
#include <FL/Fl_Widget.H>

#include "callbacks.h"
#include "enumerations.h"
#include "event_handler.h"


Fl_Widget_Tracker* go_fltk_new_Widget_Tracker(Fl_Widget* w) {
  return new Fl_Widget_Tracker(w);
}
Fl_Widget* go_fltk_Widget_Tracker_widget(Fl_Widget_Tracker* t) {
  return t->widget();
}
int go_fltk_Widget_Tracker_exists(Fl_Widget_Tracker* t) {
  return t->exists();
}
void go_fltk_Widget_Tracker_delete(Fl_Widget_Tracker* t) {
  delete t;
}

void go_fltk_delete_widget(Fl_Widget *w) {
  Fl::delete_widget(w);
}
void go_fltk_Widget_set_box(Fl_Widget *w, int box) {
  w->box((Fl_Boxtype)box);
}
void go_fltk_Widget_set_callback(Fl_Widget *w, uintptr_t id) {
  w->callback(callback_handler, (void*)id);
}
int go_fltk_Widget_add_deletion_handler(Fl_Widget* w, uintptr_t id) {
  WidgetWithDeletionHandler* wh = dynamic_cast<WidgetWithDeletionHandler*>(w);
  if (wh == nullptr) {
    return 0;
  }
  wh->add_deletion_handler(id);
  return 1;
}
void go_fltk_Widget_when(Fl_Widget* w, int when) {
  w->when(when);
}
int go_fltk_Widget_set_event_handler(Fl_Widget* w, int id) {
  WidgetWithEventHandler* wh = dynamic_cast<WidgetWithEventHandler*>(w);
  if (wh == nullptr) {
    return 0;
  }
  wh->set_event_handler(id);
  return 1;
}
int go_fltk_Widget_set_resize_handler(Fl_Widget* w, uintptr_t id) {
  WidgetWithResizeHandler* wh = dynamic_cast<WidgetWithResizeHandler*>(w);
  if (wh == nullptr) {
    return 0;
  }
  wh->set_resize_handler(id);
  return 1;
}
int go_fltk_Widget_set_draw_handler(Fl_Widget* w, uintptr_t id) {
  WidgetWithDrawHandler* wh = dynamic_cast<WidgetWithDrawHandler*>(w);
  if (wh == nullptr) {
    return 0;
  }
  wh->set_draw_handler(id);
  return 1;
}
void go_fltk_Widget_draw(Fl_Widget *w) { w->draw(); }
void go_fltk_Widget_basedraw(Fl_Widget *w) {
  WidgetWithDrawHandler* wh = dynamic_cast<WidgetWithDrawHandler*>(w);
  if (wh == nullptr) {
    w->draw();
  } else {
    wh->basedraw();    
  }    
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
void go_fltk_Widget_set_labelcolor(Fl_Widget *w, unsigned int tcol) {
  w->labelcolor(tcol);
}
void go_fltk_Widget_clear_visible_focus(Fl_Widget *w) {
  w->clear_visible_focus();
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
unsigned int go_fltk_Widget_active(Fl_Widget* w) {
  return w->active();
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
int go_fltk_Widget_visible(Fl_Widget* w) {
	return w->visible();
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
void go_fltk_Widget_set_image(Fl_Widget* w, Fl_Image *i) {
  w->image(i);
}
void go_fltk_Widget_set_deimage(Fl_Widget* w, Fl_Image *i) {
  w->deimage(i);
}
int go_fltk_Widget_box(Fl_Widget *w) { 
    return w->box();
}
unsigned int go_fltk_Widget_labelcolor(Fl_Widget *w) { 
    return w->labelcolor();
}
unsigned int go_fltk_Widget_color(Fl_Widget *w) { 
    return w->color();
}
const char *go_fltk_Widget_label(Fl_Widget *w) { 
    return w->label();
}
unsigned int go_fltk_Widget_align(Fl_Widget* w) { 
    return w->align();
}
unsigned char go_fltk_Widget_type(Fl_Widget* w) { 
    return w->type();
}
int go_fltk_Widget_labelfont(Fl_Widget *w) {
    return w->labelfont();
}
int go_fltk_Widget_labelsize(Fl_Widget *w) {
    return w->labelsize();
}
int go_fltk_Widget_labeltype(Fl_Widget *w) {
    return w->labeltype();
}
void go_fltk_Widget_set_tooltip(Fl_Widget *w, const char *tooltip) {
    w->copy_tooltip(tooltip);
}  
Fl_Group *go_fltk_Widget_parent(Fl_Widget *w) {
    return w->parent();
}
int go_fltk_Widget_take_focus(Fl_Widget *w) {
    return w->take_focus();
}
int go_fltk_Widget_has_focus(Fl_Widget *w) {
	return Fl::focus() == w;
}
unsigned int go_fltk_Widget_changed(Fl_Widget *w) {
  return w->changed();
}
