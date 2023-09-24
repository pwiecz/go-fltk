#include "wizard.h"

#include <FL/Fl_Wizard.H>

#include "event_handler.h"


class GWizard : public EventHandler<Fl_Wizard> {
public:
    GWizard(int x, int y, int w, int h, const char* label)
    : EventHandler<Fl_Wizard>(x, y, w, h, label) {}
};

GWizard *go_fltk_new_Wizard(int x, int y, int w, int h, const char *label) {
  return new GWizard(x, y, w, h, label);
}

void go_fltk_Wizard_next(Fl_Wizard* wizard) {
  wizard->next();
}
void go_fltk_Wizard_prev(Fl_Wizard* wizard) {
  wizard->prev();
}
Fl_Widget* go_fltk_Wizard_value(Fl_Wizard* wizard) {
  return wizard->value();
}
void go_fltk_Wizard_set_value(Fl_Wizard* wizard, Fl_Widget* value) {
  wizard->value(value);
}
