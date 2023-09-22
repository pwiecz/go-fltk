#include "file_chooser.h"

#include <FL/Fl_File_Chooser.H>
#include <FL/Fl_Native_File_Chooser.H>

#include "_cgo_export.h"


const int go_FL_FileChooser_SINGLE = Fl_File_Chooser::SINGLE;
const int go_FL_FileChooser_MULTI = Fl_File_Chooser::MULTI;
const int go_FL_FileChooser_CREATE = Fl_File_Chooser::CREATE;
const int go_FL_FileChooser_DIRECTORY = Fl_File_Chooser::DIRECTORY;


Fl_File_Chooser *go_fltk_new_FileChooser(const char* message, const char* pattern, int type, const char* title) {
  return new Fl_File_Chooser(message, pattern, type, title);
}

void go_fltk_FileChooser_destroy(Fl_File_Chooser* fileChooser) {
  delete fileChooser;
}

static void filechooser_callback_handler(Fl_File_Chooser* fc, void* data) {
  _go_callbackHandler((uintptr_t)data);
}

void go_fltk_FileChooser_set_callback(Fl_File_Chooser* fileChooser, uintptr_t id) {
  fileChooser->callback(filechooser_callback_handler, (void*)id);
}

void go_fltk_FileChooser_popup(Fl_File_Chooser* fileChooser) {
  // Copy of popup() functions from fl_file_dir.cxx from Fltk source code.
  fileChooser->show();

  // deactivate Fl::grab(), because it is incompatible with modal windows
  Fl_Window* g = Fl::grab();
  if (g) Fl::grab(0);

  while (fileChooser->shown())
    Fl::wait();

  if (g) // regrab the previous popup menu, if there was one
    Fl::grab(g);
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

const int go_FL_NativeFileChooser_NO_OPTIONS = Fl_Native_File_Chooser::NO_OPTIONS;
const int go_FL_NativeFileChooser_SAVEAS_CONFIRM = Fl_Native_File_Chooser::SAVEAS_CONFIRM;
const int go_FL_NativeFileChooser_NEW_FOLDER = Fl_Native_File_Chooser::NEW_FOLDER;
const int go_FL_NativeFileChooser_PREVIEW = Fl_Native_File_Chooser::PREVIEW;
const int go_FL_NativeFileChooser_USE_FILTER_EXT = Fl_Native_File_Chooser::USE_FILTER_EXT;

const int go_FL_NativeFileChooser_BROWSE_FILE = Fl_Native_File_Chooser::BROWSE_FILE;
const int go_FL_NativeFileChooser_BROWSE_DIRECTORY = Fl_Native_File_Chooser::BROWSE_DIRECTORY;
const int go_FL_NativeFileChooser_BROWSE_MULTI_FILE = Fl_Native_File_Chooser::BROWSE_MULTI_FILE;
const int go_FL_NativeFileChooser_BROWSE_MULTI_DIRECTORY = Fl_Native_File_Chooser::BROWSE_MULTI_DIRECTORY;
const int go_FL_NativeFileChooser_BROWSE_SAVE_FILE = Fl_Native_File_Chooser::BROWSE_SAVE_FILE;
const int go_FL_NativeFileChooser_BROWSE_SAVE_DIRECTORY = Fl_Native_File_Chooser::BROWSE_SAVE_DIRECTORY;

Fl_Native_File_Chooser* go_fltk_new_NativeFileChooser() {
  return new Fl_Native_File_Chooser();
}

void go_fltk_NativeFileChooser_destroy(Fl_Native_File_Chooser* fileChooser) {
  delete fileChooser;
}
int go_fltk_NativeFileChooser_count(const Fl_Native_File_Chooser* fileChooser) {
  return fileChooser->count();
}
const char* go_fltk_NativeFileChooser_directory(const Fl_Native_File_Chooser* fileChooser) {
  return fileChooser->directory();
}
void go_fltk_NativeFileChooser_set_directory(Fl_Native_File_Chooser* fileChooser, const char* directory) {
  fileChooser->directory(directory);
}
const char* go_fltk_NativeFileChooser_filename(const Fl_Native_File_Chooser* fileChooser) {
  return fileChooser->filename();
}
const char* go_fltk_NativeFileChooser_nth_filename(const Fl_Native_File_Chooser* fileChooser, int n) {
  return fileChooser->filename(n);
}
const char* go_fltk_NativeFileChooser_filter(const Fl_Native_File_Chooser* fileChooser) {
  return fileChooser->filter();
}
void go_fltk_NativeFileChooser_set_filter(Fl_Native_File_Chooser* fileChooser, const char* filter) {
  return fileChooser->filter(filter);
}
int go_fltk_NativeFileChooser_filter_value(const Fl_Native_File_Chooser* fileChooser) {
  return fileChooser->filter_value();
}
void go_fltk_NativeFileChooser_set_filter_value(Fl_Native_File_Chooser* fileChooser, int i) {
  fileChooser->filter_value(i);
}
int go_fltk_NativeFileChooser_filters(const Fl_Native_File_Chooser* fileChooser) {
  return fileChooser->filters();
}
int go_fltk_NativeFileChooser_options(const Fl_Native_File_Chooser* fileChooser) {
  return fileChooser->options();
}
void go_fltk_NativeFileChooser_set_options(Fl_Native_File_Chooser* fileChooser, int options) {
  fileChooser->options(options);
}
const char* go_fltk_NativeFileChooser_preset_file(const Fl_Native_File_Chooser* fileChooser) {
  return fileChooser->preset_file();
}
void go_fltk_NativeFileChooser_set_preset_file(Fl_Native_File_Chooser* fileChooser, const char* presetFile) {
  fileChooser->preset_file(presetFile);
}
int go_fltk_NativeFileChooser_show(Fl_Native_File_Chooser* fileChooser) {
  return fileChooser->show();
}
const char* go_fltk_NativeFileChooser_title(const Fl_Native_File_Chooser* fileChooser) {
  return fileChooser->title();
}
void go_fltk_NativeFileChooser_set_title(Fl_Native_File_Chooser* fileChooser, const char* title) {
  fileChooser->title(title);
}
int go_fltk_NativeFileChooser_type(const Fl_Native_File_Chooser* fileChooser) {
  return fileChooser->type();
}
void go_fltk_NativeFileChooser_set_type(Fl_Native_File_Chooser* fileChooser, int type) {
  fileChooser->type(type);
}
