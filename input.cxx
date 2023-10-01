#include "input.h"

#include <FL/Fl_Input.H>
#include <FL/Fl_Output.H>
#include <FL/Fl_Float_Input.H>
#include <FL/Fl_Int_Input.H>

#include "event_handler.h"


const unsigned char go_FL_FLOAT_INPUT = 1;
const unsigned char go_FL_INT_INPUT = 2;

class GInput : public EventHandler<Fl_Input> {
public:
  GInput(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Input>(x, y, w, h, label) {}
};

GInput *go_fltk_new_Input(int x, int y, int w, int h, const char *text) {
  return new GInput(x, y, w, h, text);
}

const char *go_fltk_Input_value(Fl_Input *in) {
  return in->value();
}
int go_fltk_Input_set_value(Fl_Input *in, const char *t) {
  return in->value(t);
}
void go_fltk_Input_resize(Fl_Input *in, int x, int y, int w, int h) {
  in->resize(x, y, w, h);
}
void go_fltk_Input_set_insert_position(Fl_Input *in, int p, int m) {
	in->insert_position(p, m);
}
int go_fltk_Input_insert_position(Fl_Input *in) {
	return in->insert_position();
}
int go_fltk_Input_mark(Fl_Input *in) {
	return in->mark();
}

class GOutput : public EventHandler<Fl_Output> {
public:
  GOutput(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Output>(x, y, w, h, label) {}
};

GOutput *go_fltk_new_Output(int x, int y, int w, int h, const char *text) {
  return new GOutput(x, y, w, h, text);
}

class GFloat_Input : public EventHandler<Fl_Float_Input> {
public:
  GFloat_Input(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Float_Input>(x, y, w, h, label) {}
};

GFloat_Input *go_fltk_new_Float_Input(int x, int y, int w, int h, const char *text) {
  return new GFloat_Input(x, y, w, h, text);
}

class GInt_Input : public EventHandler<Fl_Int_Input> {
public:
  GInt_Input(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Int_Input>(x, y, w, h, label) {}
};

GInt_Input *go_fltk_new_Int_Input(int x, int y, int w, int h, const char *text) {
  return new GInt_Input(x, y, w, h, text);
}

