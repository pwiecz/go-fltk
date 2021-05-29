#include "file_chooser.h"

#include <Fl/Fl_File_Chooser.H>

#include "_cgo_export.h"


const int go_FL_SINGLE = Fl_File_Chooser::SINGLE;
const int go_FL_MULTI = Fl_File_Chooser::MULTI;
const int go_FL_CREATE = Fl_File_Chooser::CREATE;
const int go_FL_DIRECTORY = Fl_File_Chooser::DIRECTORY;


Fl_File_Chooser *go_fltk_new_FileChooser(const char* message, const char* pattern, int type, const char* title) {
  return new Fl_File_Chooser(message, pattern, type, title);
}

void go_fltk_FileChooser_destroy(Fl_File_Chooser* fileChooser) {
  delete fileChooser;
}

static void filechooser_callback_handler(Fl_File_Chooser* fc, void* data) {
  _go_callbackHandler(data);
}

void go_fltk_FileChooser_set_callback(Fl_File_Chooser* fileChooser, void* id) {
  fileChooser->callback(filechooser_callback_handler, id);
}

void go_fltk_FileChooser_show(Fl_File_Chooser* fileChooser) {
  fileChooser->show();
}
int go_fltk_FileChooser_shown(Fl_File_Chooser* fileChooser) {
  return fileChooser->shown();
}

void go_fltk_FileChooser_preview(Fl_File_Chooser* fileChooser, int enable) {
  fileChooser->preview(enable);
}

int go_fltk_FileChooser_count(Fl_File_Chooser* fileChooser) {
  return fileChooser->count();
}
const char* go_fltk_FileChooser_value(Fl_File_Chooser* fileChooser, int position) {
  return fileChooser->value(position);
}

char* go_fltk_file_chooser(const char* message, const char* pattern, const char* initialFilename, int relative) {
  return fl_file_chooser(message, pattern, initialFilename, relative);
}

