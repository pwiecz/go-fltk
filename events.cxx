#include "events.h"

#include <FL/Fl.H>


const int go_FL_LEFT_MOUSE = FL_LEFT_MOUSE;
const int go_FL_MIDDLE_MOUSE = FL_MIDDLE_MOUSE;
const int go_FL_RIGHT_MOUSE = FL_RIGHT_MOUSE;

const int go_FL_SHIFT = FL_SHIFT;
const int go_FL_CAPS_LOCK = FL_CAPS_LOCK;
const int go_FL_CTRL = FL_CTRL;
const int go_FL_ALT = FL_ALT;
const int go_FL_NUM_LOCK = FL_NUM_LOCK;
const int go_FL_META = FL_META;
const int go_FL_SCROLL_LOCK = FL_SCROLL_LOCK;
const int go_FL_BUTTON1 = FL_BUTTON1;
const int go_FL_BUTTON2 = FL_BUTTON2;
const int go_FL_BUTTON3 = FL_BUTTON3;

int go_fltk_event() { return Fl::event(); }
int go_fltk_event_button() { return Fl::event_button(); }
int go_fltk_event_button1() { return Fl::event_button1(); }
int go_fltk_event_x() { return Fl::event_x(); }
int go_fltk_event_y() { return Fl::event_y(); }
int go_fltk_event_x_root() { return Fl::event_x_root(); }
int go_fltk_event_y_root() { return Fl::event_y_root(); }
int go_fltk_event_dx() { return Fl::event_dx(); }
int go_fltk_event_dy() { return Fl::event_dy(); }
int go_fltk_event_key() { return Fl::event_key(); }
int go_fltk_event_is_click() { return Fl::event_is_click(); }
int go_fltk_event_clicks() { return Fl::event_clicks(); }
void go_fltk_event_set_clicks(int i) { Fl::event_clicks(i); }
int go_fltk_event_state() { return Fl::event_state(); }
const char* go_fltk_event_text() { return Fl::event_text(); }
