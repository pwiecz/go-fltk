#include "dialogs.h"

#include <FL/fl_ask.H>


void go_fltk_message_box(const char *title, const char* message) {
  fl_message_title(title);
  fl_message("%s", message);
}

int go_fltk_choice_dialog(const char* message, const char* option0, const char* option1, const char* option2) {
  return fl_choice("%s", option0, option1, option2, message);
}
