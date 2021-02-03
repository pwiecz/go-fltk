#include "menu.h"

#include <cstdint>

#include <Fl/Fl_Menu_.H>
#include <Fl/Fl_Menu_Button.H>

#include "enumerations.h"

#include "_cgo_export.h"


int go_fltk_Menu_add(Fl_Menu_ *m, char *label, int shortcut, int callback, int flags) {
  return m->add(label, shortcut, callback_handler, (void*)(uintptr_t)callback, flags);
}

void go_fltk_Menu_set_value(Fl_Menu_ *m, int value) { 
  m->value(value);
}

int go_fltk_Menu_value(Fl_Menu_ *m) {
  return m->value();
}

class GMenuButton : public Fl_Menu_Button {
public:
  GMenuButton(int x, int y, int w, int h, const char* text, void* destroyCallbackId)
    : Fl_Menu_Button(x, y, w, h, text)
    , m_destroyCallbackId(destroyCallbackId) {}

  ~GMenuButton() {
    _go_callbackHandler(m_destroyCallbackId);
  }

private:
  void * const m_destroyCallbackId;
};

GMenuButton* go_fltk_new_MenuButton(int x, int y, int w, int h, const char* text, void* destroyCallback) {
  return new GMenuButton(x, y, w, h, text, destroyCallback);
}
void go_fltk_MenuButton_set_type(GMenuButton* m, int menuType) {
  m->type(menuType);
}
void go_fltk_MenuButton_popup(GMenuButton* m) {
  m->popup();
}
void go_fltk_MenuButton_destroy(GMenuButton* m) {
  delete m;
}

const int go_FL_POPUP1 = Fl_Menu_Button::POPUP1;
const int go_FL_POPUP2 = Fl_Menu_Button::POPUP2;
const int go_FL_POPUP12 = Fl_Menu_Button::POPUP12;
const int go_FL_POPUP3 = Fl_Menu_Button::POPUP3;
const int go_FL_POPUP13 = Fl_Menu_Button::POPUP13;
const int go_FL_POPUP23 = Fl_Menu_Button::POPUP23;
const int go_FL_POPUP123 = Fl_Menu_Button::POPUP123;
