#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Wizard Fl_Wizard;
  typedef struct GWizard GWizard;
  typedef struct Fl_Widget Fl_Widget;

  extern GWizard *go_fltk_new_Wizard(int x, int y, int w, int h, const char *label);

  extern void go_fltk_Wizard_next(Fl_Wizard* wizard);
  extern void go_fltk_Wizard_prev(Fl_Wizard* wizard);
  extern Fl_Widget* go_fltk_Wizard_value(Fl_Wizard* wizard);
  extern void go_fltk_Wizard_set_value(Fl_Wizard* wizard, Fl_Widget* value);


#ifdef __cplusplus
}
#endif
