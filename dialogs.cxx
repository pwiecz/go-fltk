#include "dialogs.h"

#include <FL/fl_ask.H>
#include <FL/Fl_Color_Chooser.H>

#include <stddef.h>
#include <cstdlib>
#include <cstring>

void go_fltk_message_box(const char *title, const char* message) {
  fl_message_title(title);
  fl_message("%s", message);
}

int go_fltk_choice_dialog(const char* message, const char* option0, const char* option1, const char* option2) {
  return fl_choice("%s", option0, option1, option2, message);
}

void go_fltk_alert_dialog(const char* message) {
  fl_alert(message, 0);
}

int go_fltk_ask_dialog(const char* message) {
  return fl_ask(message, 0);
}

char* go_fltk_input_dialog(int maxchar, const char* message) {
  Fl_String outputFlString = fl_input_str(maxchar, message, 0);
  const char* outputStr = outputFlString.value();
  char* output = (char*) malloc(maxchar);
  if (outputStr != NULL){
    // Not sure why fltk Fl_String is limited which requires us to copy values. What a flaw!
    strncpy(output, outputStr, maxchar);
  } else {
    strcpy(output, "");
  }
  return output;
}

char* go_fltk_password_dialog(int maxchar, const char* message, const char* defaultPassword) {
  Fl_String outputFlString = fl_password_str(maxchar, message, defaultPassword, 0);
  const char* outputStr = outputFlString.value();
  char* output = (char*) malloc(maxchar);
  if (outputStr != NULL){
    strncpy(output, outputStr, maxchar);
  } else {
    strcpy(output, "");
  }
  return output;
}

void go_fltk_set_title_dialog(const char* title) {
  fl_message_title(title);
}

int go_fltk_color_chooser_dialog(const char* title, double *r, double *g, double *b, int cmode) {
  double x, y, z;
  int result = fl_color_chooser(title, &x, &y, &z, cmode);
  *r = x; *g = y; *b = z;
  return result;
}