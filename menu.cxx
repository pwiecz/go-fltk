#include "menu.h"

#include <cstdint>

#include <FL/Fl_Menu_.H>
#include <FL/Fl_Menu_Button.H>
#include <FL/Fl_Menu_Bar.H>
#include <FL/Fl_Multi_Label.H>

#include "callbacks.h"
#include "event_handler.h"


int go_fltk_Menu_add(Fl_Menu_ *m, char *label, int shortcut, int callback, int flags) {
  return m->add(label, shortcut, callback_handler, (void*)(uintptr_t)callback, flags);
}

int go_fltk_Menu_add_with_icon(Fl_Menu_ *m, char *label, int shortcut, int callback, int flags, Fl_Image *img) {
	int idx = m->add(label, shortcut, callback_handler, (void*)(uintptr_t)callback, flags);
	Fl_Menu_Item *item = (Fl_Menu_Item*)&(m->menu()[idx]);
	Fl_Multi_Label *ml = new Fl_Multi_Label;
	ml->typea = FL_IMAGE_LABEL;
	ml->labela = (const char*)img;
	
	ml->typeb = FL_NORMAL_LABEL;
	ml->labelb = item->label();
	ml->label(item);
	return idx;
}

int go_fltk_Menu_insert(Fl_Menu_ *m, int index, char *label, int shortcut, int callback, int flags) {
  return m->insert(index, label, shortcut, callback_handler, (void*)(uintptr_t)callback, flags);
}

void go_fltk_Menu_remove(Fl_Menu_ *m, int index) { m->remove(index); }

void go_fltk_Menu_clear(Fl_Menu_ *m) { m->clear(); }

void go_fltk_Menu_replace(Fl_Menu_ *m, int index, const char *label) {
  m->replace(index, label);
}
int go_fltk_Menu_find_index(Fl_Menu_ *m, const char *label) {
  return m->find_index(label);  
}  

void go_fltk_Menu_set_value(Fl_Menu_ *m, int value) { 
  m->value(value);
}

int go_fltk_Menu_value(Fl_Menu_ *m) { return m->value(); }

const char* go_fltk_Menu_selected_text(Fl_Menu_* m) { return m->text(); }

const char* go_fltk_Menu_text(Fl_Menu_* m, int index) { return m->text(index); }

int go_fltk_Menu_size(Fl_Menu_ *m) { return m->size(); }


class GMenu_Button : public EventHandler<Fl_Menu_Button> {
public:
  GMenu_Button(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Menu_Button>(x, y, w, h, label) {}
};

GMenu_Button* go_fltk_new_MenuButton(int x, int y, int w, int h, const char* text) {
  return new GMenu_Button(x, y, w, h, text);
}
void go_fltk_MenuButton_set_type(Fl_Menu_Button* m, int menuType) {
  m->type(menuType);
}
void go_fltk_MenuButton_popup(Fl_Menu_Button* m) {
  m->popup();
}

class GMenu_Bar : public EventHandler<Fl_Menu_Bar> {
public:
  GMenu_Bar(int x, int y, int w, int h, const char *label)
    : EventHandler<Fl_Menu_Bar>(x, y, w, h, label) {}
};

GMenu_Bar* go_fltk_new_MenuBar(int x, int y, int w, int h, const char* text) {
  return new GMenu_Bar(x, y, w, h, text);
}

const int go_FL_POPUP1 = Fl_Menu_Button::POPUP1;
const int go_FL_POPUP2 = Fl_Menu_Button::POPUP2;
const int go_FL_POPUP12 = Fl_Menu_Button::POPUP12;
const int go_FL_POPUP3 = Fl_Menu_Button::POPUP3;
const int go_FL_POPUP13 = Fl_Menu_Button::POPUP13;
const int go_FL_POPUP23 = Fl_Menu_Button::POPUP23;
const int go_FL_POPUP123 = Fl_Menu_Button::POPUP123;


const int go_FL_MENU_INACTIVE = FL_MENU_INACTIVE;
const int go_FL_MENU_TOGGLE = FL_MENU_TOGGLE;
const int go_FL_MENU_VALUE = FL_MENU_VALUE;
const int go_FL_MENU_RADIO = FL_MENU_RADIO;
const int go_FL_MENU_INVISIBLE = FL_MENU_INVISIBLE;
const int go_FL_SUBMENU = FL_SUBMENU;
const int go_FL_MENU_DIVIDER = FL_MENU_DIVIDER;
