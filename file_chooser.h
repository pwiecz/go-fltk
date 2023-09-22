#pragma once

#include <stdint.h>


#ifdef __cplusplus
extern "C" {
#endif

  extern const int go_FL_FileChooser_SINGLE;
  extern const int go_FL_FileChooser_MULTI;
  extern const int go_FL_FileChooser_CREATE;
  extern const int go_FL_FileChooser_DIRECTORY;

  typedef struct Fl_File_Chooser Fl_File_Chooser;
  typedef struct Fl_Native_File_Chooser Fl_Native_File_Chooser;

  extern Fl_File_Chooser* go_fltk_new_FileChooser(const char* pathname, const char* pattern, int type, const char* title);

  extern void go_fltk_FileChooser_destroy(Fl_File_Chooser* fileChooser);
  extern void go_fltk_FileChooser_set_callback(Fl_File_Chooser* fileChooser, uintptr_t id);
  extern void go_fltk_FileChooser_popup(Fl_File_Chooser* fileChooser);
  extern void go_fltk_FileChooser_show(Fl_File_Chooser* fileChooser);
  extern int go_fltk_FileChooser_shown(Fl_File_Chooser* fileChooser);
  extern void go_fltk_FileChooser_preview(Fl_File_Chooser* fileChooser, int enable);
  extern int go_fltk_FileChooser_count(Fl_File_Chooser* fileChooser);
  extern const char* go_fltk_FileChooser_value(Fl_File_Chooser* fileChooser, int position);

  extern char* go_fltk_file_chooser(const char* message, const char* pattern, const char* initialFilename, int relative);

  extern const int go_FL_NativeFileChooser_NO_OPTIONS;
  extern const int go_FL_NativeFileChooser_SAVEAS_CONFIRM;
  extern const int go_FL_NativeFileChooser_NEW_FOLDER;
  extern const int go_FL_NativeFileChooser_PREVIEW;
  extern const int go_FL_NativeFileChooser_USE_FILTER_EXT;

  extern const int go_FL_NativeFileChooser_BROWSE_FILE;
  extern const int go_FL_NativeFileChooser_BROWSE_DIRECTORY;
  extern const int go_FL_NativeFileChooser_BROWSE_MULTI_FILE;
  extern const int go_FL_NativeFileChooser_BROWSE_MULTI_DIRECTORY;
  extern const int go_FL_NativeFileChooser_BROWSE_SAVE_FILE;
  extern const int go_FL_NativeFileChooser_BROWSE_SAVE_DIRECTORY;

  extern Fl_Native_File_Chooser* go_fltk_new_NativeFileChooser();
  extern void go_fltk_NativeFileChooser_destroy(Fl_Native_File_Chooser* fileChooser);
  extern int go_fltk_NativeFileChooser_count(const Fl_Native_File_Chooser* fileChooser);
  extern const char* go_fltk_NativeFileChooser_directory(const Fl_Native_File_Chooser* fileChooser);
  extern void go_fltk_NativeFileChooser_set_directory(Fl_Native_File_Chooser* fileChooser, const char* directory);
  extern const char* go_fltk_NativeFileChooser_filename(const Fl_Native_File_Chooser* fileChooser);
  extern const char* go_fltk_NativeFileChooser_nth_filename(const Fl_Native_File_Chooser* fileChooser, int n);
  extern const char* go_fltk_NativeFileChooser_filter(const Fl_Native_File_Chooser* fileChooser);
  extern void go_fltk_NativeFileChooser_set_filter(Fl_Native_File_Chooser* fileChooser, const char* filter);
  extern int go_fltk_NativeFileChooser_filter_value(const Fl_Native_File_Chooser* fileChooser);
  extern void go_fltk_NativeFileChooser_set_filter_value(Fl_Native_File_Chooser* fileChooser, int i);
  extern int go_fltk_NativeFileChooser_filters(const Fl_Native_File_Chooser* fileChooser);
  extern int go_fltk_NativeFileChooser_options(const Fl_Native_File_Chooser* fileChooser);
  extern void go_fltk_NativeFileChooser_set_options(Fl_Native_File_Chooser* fileChooser, int options);
  extern const char* go_fltk_NativeFileChooser_preset_file(const Fl_Native_File_Chooser* fileChooser);
  extern void go_fltk_NativeFileChooser_set_preset_file(Fl_Native_File_Chooser* fileChooser, const char* presetFile);
  extern int go_fltk_NativeFileChooser_show(Fl_Native_File_Chooser* fileChooser);
  extern const char* go_fltk_NativeFileChooser_title(const Fl_Native_File_Chooser* fileChooser);
  extern void go_fltk_NativeFileChooser_set_title(Fl_Native_File_Chooser* fileChooser, const char* title);
  extern int go_fltk_NativeFileChooser_type(const Fl_Native_File_Chooser* fileChooser);
  extern void go_fltk_NativeFileChooser_set_type(Fl_Native_File_Chooser* fileChooser, int type);
  
#ifdef __cplusplus
}
#endif
