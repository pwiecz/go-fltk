#include "input_choice.h"

#include <FL/Fl_Input_Choice.H>

#include "event_handler.h"


class GInputChoice : public EventHandler<Fl_Input_Choice> {
public:
  GInputChoice(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Input_Choice>(x, y, w, h, label) {}
};

GInputChoice *go_fltk_new_Input_Choice(int x, int y, int w, int h, const char *label) {
  return new GInputChoice(x, y, w, h, label);
}

void go_fltk_Input_Choice_clear(Fl_Input_Choice *c) {
  c->clear();
}

const char *go_fltk_Input_Choice_value(Fl_Input_Choice *c) {
  return c->value();
}

void go_fltk_Input_Choice_set_value(Fl_Input_Choice *c,
				    const char *label) {
  c->value(label);
}

void go_fltk_Input_Choice_set_value_index(Fl_Input_Choice *c,
					  int index) {
  c->value(index);
}

int go_fltk_Input_Choice_update_menubutton(Fl_Input_Choice* c) {
  return c->update_menubutton();
}


Fl_Input *go_fltk_Input_Choice_input(Fl_Input_Choice *c) {
  return c->input();
}

Fl_Menu_Button *go_fltk_Input_Choice_menubutton(Fl_Input_Choice *c) {
  return c->menubutton();
}
