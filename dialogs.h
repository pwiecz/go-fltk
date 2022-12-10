#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  extern void go_fltk_message_box(const char *title, const char* message);

  extern int go_fltk_choice_dialog(const char* message, const char* option0, const char* option1, const char* option2);

  extern void go_fltk_alert_dialog(const char* message);

  extern int go_fltk_ask_dialog(const char* message);

  extern char* go_fltk_input_dialog(int maxchar, const char* message);

  extern char* go_fltk_password_dialog(int maxchar, const char* message, const char* defaultPassword);

  extern void go_fltk_set_title_dialog(const char* title);

#ifdef __cplusplus
}
#endif
