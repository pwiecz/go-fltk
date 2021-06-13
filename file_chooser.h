#pragma once

#include <stdint.h>


#ifdef __cplusplus
extern "C" {
#endif

  extern const int go_FL_SINGLE;
  extern const int go_FL_MULTI;
  extern const int go_FL_CREATE;
  extern const int go_FL_DIRECTORY;

  typedef struct Fl_File_Chooser Fl_File_Chooser;

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

#ifdef __cplusplus
}
#endif
