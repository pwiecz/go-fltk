#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GInputChoice GInputChoice;
  typedef struct Fl_Input_Choice Fl_Input_Choice;
  typedef struct Fl_Input Fl_Input;
  typedef struct Fl_Menu_Button Fl_Menu_Button;

  extern GInputChoice *go_fltk_new_Input_Choice(int x, int y, int w, int h, const char *label);

  extern void go_fltk_Input_Choice_clear(Fl_Input_Choice *c);

  extern const char *go_fltk_Input_Choice_value(Fl_Input_Choice *c);

  extern void go_fltk_Input_Choice_set_value(Fl_Input_Choice *c,
                                             const char *label);

  extern void go_fltk_Input_Choice_set_value_index(Fl_Input_Choice *c,
                                                   int index);

  extern int go_fltk_Input_Choice_update_menubutton(Fl_Input_Choice* c);

  extern Fl_Input *go_fltk_Input_Choice_input(Fl_Input_Choice *c);

  extern Fl_Menu_Button *go_fltk_Input_Choice_menubutton(Fl_Input_Choice *c);
  
#ifdef __cplusplus
}
#endif
