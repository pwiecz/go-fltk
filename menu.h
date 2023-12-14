#pragma once

#ifdef __cplusplus
extern "C" {
#endif
	
  typedef struct Fl_Menu_ Fl_Menu_;
  typedef struct GMenu_Button GMenu_Button;
  typedef struct Fl_Menu_Button Fl_Menu_Button;
  typedef struct GMenu_Bar GMenu_Bar;
  typedef struct Fl_Image Fl_Image;

  extern int go_fltk_Menu_add(Fl_Menu_* m, char* label, int shortcut, int callback, int flags);
  extern int go_fltk_Menu_add_with_icon(Fl_Menu_ *m, char *label, int shortcut, int callback, int flags, Fl_Image *img);
  extern int go_fltk_Menu_insert(Fl_Menu_* m, int index, char* label, int shortcut, int callback, int flags);
  extern void go_fltk_Menu_remove(Fl_Menu_ *m, int index);
  extern void go_fltk_Menu_clear(Fl_Menu_ *m);
  extern void go_fltk_Menu_replace(Fl_Menu_ *m, int index, const char *label);
  extern int go_fltk_Menu_find_index(Fl_Menu_ *m, const char* label);
  extern void go_fltk_Menu_set_value(Fl_Menu_* m, int value);
  extern int go_fltk_Menu_value(Fl_Menu_* m);
  extern const char* go_fltk_Menu_text(Fl_Menu_* m, int index);
  extern const char* go_fltk_Menu_selected_text(Fl_Menu_* m);
  extern int go_fltk_Menu_size(Fl_Menu_* m);

  extern GMenu_Button* go_fltk_new_MenuButton(int x, int y, int w, int h, const char* text);
  extern void go_fltk_MenuButton_set_type(Fl_Menu_Button* m, int type);
  extern void go_fltk_MenuButton_popup(Fl_Menu_Button* m);

  extern GMenu_Bar* go_fltk_new_MenuBar(int x, int y, int w, int h, const char* text);

  extern const int go_FL_POPUP1;
  extern const int go_FL_POPUP2;
  extern const int go_FL_POPUP12;
  extern const int go_FL_POPUP3;
  extern const int go_FL_POPUP13;
  extern const int go_FL_POPUP23;
  extern const int go_FL_POPUP123;

  extern const int go_FL_MENU_INACTIVE;
  extern const int go_FL_MENU_TOGGLE;
  extern const int go_FL_MENU_VALUE;
  extern const int go_FL_MENU_RADIO;
  extern const int go_FL_MENU_INVISIBLE;
  extern const int go_FL_SUBMENU;
  extern const int go_FL_MENU_DIVIDER;


#ifdef __cplusplus
}
#endif
