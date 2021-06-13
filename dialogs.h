#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  extern void go_fltk_message_box(const char *title, const char* message);

  extern int go_fltk_choice_dialog(const char* message, const char* option0, const char* option1, const char* option2);

#ifdef __cplusplus
}
#endif
