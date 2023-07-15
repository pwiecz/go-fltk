#pragma once

#include <stdint.h>


#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Widget_Tracker Fl_Widget_Tracker;
  typedef struct Fl_Widget Fl_Widget;
  typedef struct Fl_Group Fl_Group;
  typedef struct Fl_Image Fl_Image;

  extern Fl_Widget_Tracker* go_fltk_new_Widget_Tracker(Fl_Widget* t);
  extern Fl_Widget* go_fltk_Widget_Tracker_widget(Fl_Widget_Tracker* t);
  extern void go_fltk_Widget_Tracker_delete(Fl_Widget_Tracker* t);
  extern int go_fltk_Widget_Tracker_exists(Fl_Widget_Tracker* t);
  extern void go_fltk_delete_widget(Fl_Widget *w);
  extern void go_fltk_Widget_set_box(Fl_Widget *w, int box);
  extern void go_fltk_Widget_set_labelfont(Fl_Widget *w, int font);
  extern void go_fltk_Widget_set_labelsize(Fl_Widget *w, int size);
  extern void go_fltk_Widget_set_labeltype(Fl_Widget *w, int type);
  extern void go_fltk_Widget_set_labelcolor(Fl_Widget *w, unsigned int tcol);
  extern void go_fltk_Widget_clear_visible_focus(Fl_Widget *w);
  extern void go_fltk_Widget_set_callback(Fl_Widget *w, uintptr_t id);
  extern int go_fltk_Widget_set_resize_handler(Fl_Widget* w, uintptr_t id);
  extern int go_fltk_Widget_set_draw_handler(Fl_Widget *w, uintptr_t id);
  extern void go_fltk_Widget_draw(Fl_Widget *w);
  // calls draw() on the widget ignoring potentially specified draw handlers.  
  extern void go_fltk_Widget_basedraw(Fl_Widget* w);
  extern int go_fltk_Widget_add_deletion_handler(Fl_Widget* w, uintptr_t id);
  extern void go_fltk_Widget_when(Fl_Widget* w, int when);
  extern int go_fltk_Widget_set_event_handler(Fl_Widget* w, int id);
  extern int go_fltk_Widget_x(Fl_Widget *w);
  extern int go_fltk_Widget_y(Fl_Widget *w);
  extern int go_fltk_Widget_w(Fl_Widget *w);
  extern int go_fltk_Widget_h(Fl_Widget *w);
  extern void go_fltk_Widget_set_align(Fl_Widget* w, unsigned int align);
  extern void go_fltk_Widget_measure_label(Fl_Widget* w, int* ww, int *hh);
  extern void go_fltk_Widget_set_position(Fl_Widget* w, int x, int y);
  extern void go_fltk_Widget_resize(Fl_Widget* w, int x, int y, int width, int height);
  extern void go_fltk_Widget_redraw(Fl_Widget* w);
  extern void go_fltk_Widget_deactivate(Fl_Widget* w);
  extern void go_fltk_Widget_activate(Fl_Widget *w);
  extern unsigned int go_fltk_Widget_active(Fl_Widget* w);  
  extern void go_fltk_Widget_set_type(Fl_Widget* w, unsigned char type);
  extern void go_fltk_Widget_show(Fl_Widget* w);
  extern void go_fltk_Widget_hide(Fl_Widget* w);
  extern int go_fltk_Widget_visible(Fl_Widget* w);
  extern unsigned int go_fltk_Widget_selection_color(Fl_Widget* w);
  extern void go_fltk_Widget_set_selection_color(Fl_Widget* w, unsigned int color);
  extern void go_fltk_Widget_set_color(Fl_Widget* w, unsigned int color);
  extern void go_fltk_Widget_set_label(Fl_Widget* w, const char* label);
  extern void go_fltk_Widget_set_image(Fl_Widget* w, Fl_Image *i);
  extern void go_fltk_Widget_set_deimage(Fl_Widget* w, Fl_Image *i);
  extern int go_fltk_Widget_box(Fl_Widget *w);
  extern const char *go_fltk_Widget_label(Fl_Widget *w);
  extern unsigned char go_fltk_Widget_type(Fl_Widget* w);
  extern unsigned int go_fltk_Widget_color(Fl_Widget *w);
  extern unsigned int go_fltk_Widget_labelcolor(Fl_Widget *w);
  extern unsigned int go_fltk_Widget_align(Fl_Widget* w);
  extern int go_fltk_Widget_labelfont(Fl_Widget *w);
  extern int go_fltk_Widget_labelsize(Fl_Widget *w);
  extern int go_fltk_Widget_labeltype(Fl_Widget *w);
  extern void go_fltk_Widget_set_tooltip(Fl_Widget* w, const char* tooltip);
  extern Fl_Group *go_fltk_Widget_parent(Fl_Widget *w);
  extern int go_fltk_Widget_take_focus(Fl_Widget *w);
  extern int go_fltk_Widget_has_focus(Fl_Widget *w);
  extern unsigned int go_fltk_Widget_changed(Fl_Widget* w);

#ifdef __cplusplus
}
#endif
