#include "dialogs.h"

#include <FL/fl_ask.H>


void go_fltk_message_box(const char *title, const char* message) {
  fl_message_title(title);
  fl_message("%s", message);
}
