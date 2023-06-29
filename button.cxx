#include "button.h"

#include <FL/Fl_Button.H>
#include <FL/Fl_Check_Button.H>
#include <FL/Fl_Light_Button.H>
#include <FL/Fl_Radio_Button.H>
#include <FL/Fl_Radio_Round_Button.H>
#include <FL/Fl_Return_Button.H>
#include <FL/Fl_Toggle_Button.H>

#include "event_handler.h"


class GButton : public EventHandler<Fl_Button> {
public:
  GButton(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Button>(x, y, w, h, label) {}
};

GButton *go_fltk_new_Button(int x, int y, int w, int h, const char *label) {
  return new GButton(x, y, w, h, label);
}

char go_fltk_Button_value(Fl_Button *b) {
  return b->value();
}
void go_fltk_Button_set_value(Fl_Button *b, int val) {
  b->value(val);
}
void go_fltk_Button_set_down_box(Fl_Button *b, int val) {
  b->down_box((Fl_Boxtype)val);
}
void go_fltk_Button_set_shortcut(Fl_Button *b, int shortcut) {
	b->shortcut(shortcut);
}
int go_fltk_Button_shortcut(Fl_Button *b) {
	return b->shortcut();
}

class GCheck_Button : public EventHandler<Fl_Check_Button> {
public:
  GCheck_Button(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Check_Button>(x, y, w, h, label) {}
};

GCheck_Button *go_fltk_new_Check_Button(int x, int y, int w, int h, const char *label) {
  return new GCheck_Button(x, y, w, h, label);
}

class GToggle_Button : public EventHandler<Fl_Toggle_Button> {
public:
  GToggle_Button(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Toggle_Button>(x, y, w, h, label) {}
};

GToggle_Button *go_fltk_new_Toggle_Button(int x, int y, int w, int h, const char *label) {
  return new GToggle_Button(x, y, w, h, label);
}

class GLight_Button : public EventHandler<Fl_Light_Button> {
public:
  GLight_Button(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Light_Button>(x, y, w, h, label) {}
};

GLight_Button *go_fltk_new_Light_Button(int x, int y, int w, int h, const char *label) {
  return new GLight_Button(x, y, w, h, label);
}

class GRadio_Button : public EventHandler<Fl_Radio_Button> {
public:
  GRadio_Button(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Radio_Button>(x, y, w, h, label) {}
};

GRadio_Button *go_fltk_new_Radio_Button(int x, int y, int w, int h, const char *label) {
  return new GRadio_Button(x, y, w, h, label);
}

class GReturn_Button : public EventHandler<Fl_Return_Button> {
public:
  GReturn_Button(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Return_Button>(x, y, w, h, label) {}
};

GReturn_Button *go_fltk_new_Return_Button(int x, int y, int w, int h, const char *label) {
  return new GReturn_Button(x, y, w, h, label);
}

class GRadio_Round_Button : public EventHandler<Fl_Radio_Round_Button> {
public:
  GRadio_Round_Button(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Radio_Round_Button>(x, y, w, h, label) {}
};

GRadio_Round_Button *go_fltk_new_Radio_Round_Button(int x, int y, int w, int h, const char *label) {
  return new GRadio_Round_Button(x, y, w, h, label);
}
